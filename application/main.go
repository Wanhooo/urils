package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"path/filepath"
	"urils/application/config"
	. "urils/application/database"
	. "urils/application/initialize"
	"urils/application/model"
)

func main() {
	// 1. 获取一个基于整个目录入口所在的路径
	dir, err := filepath.Abs(filepath.Dir("../"))
	if err != nil {
		panic(err.Error())
	}
	//println(dir)
	// 2. 配置初始化
	if err := config.Init(fmt.Sprintf("%s/config.json", dir)); err != nil {
		panic(err.Error())
	}

	// 设置调试模式
	gin.SetMode(config.Conf.Mode)

	// 3. 日志初始化
	if err := InitLogger(config.Conf.LogConfig); err != nil {
		fmt.Printf("init logger failed, err:%v\n", err)
		return
	}

	zap.S().Debugf("调试信息:%d", config.Conf.Port)
	// 3.创建路由
	Router := InitRouter()
	// zap 提供了一个S函数和L函数给我们开发者使用，调用S函数或L函数，可以得到一个全局的线程安全的logger对象
	zap.S().Infof("服务端启动...端口：%d", config.Conf.Port)

	// 4. 数据库初始化
	Orm := InitDB(config.Conf.DatabaseConfig)
	// 禁用复数
	Orm.SingularTable(true)
	// 数据迁移
	Orm.AutoMigrate(&model.User{})

	defer Orm.Close()

	// 5. redis数据库初始化
	RedisPool0 = InitRedisPool(&config.Conf.RedisConfigGroup.Default)
	defer RedisPool0.Close()
	RedisPool1 = InitRedisPool(&config.Conf.RedisConfigGroup.Sms)
	defer RedisPool1.Close()

	// 4.监听端口，默认在8080，因为客户端的vue已经运行占用了8080了，所以改成8000
	//Router.Run(fmt.Sprintf(":%d", config.Conf.Port))
	if err := Router.Run( fmt.Sprintf(":%d", config.Conf.Port) ); err != nil {
		zap.S().Panic("服务端启动失败：", err.Error())
	}
}
