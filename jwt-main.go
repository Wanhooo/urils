package main

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

func main() {

	// 头部
	header := map[string]string{
		"typ": "JWT",
		"alg": "HS256",
	}
	ret, _ := json.Marshal(header)
	input := []byte(string(ret))

	headerString := base64.StdEncoding.EncodeToString(input)
	fmt.Println(headerString)


	// 载荷
	payload := map[string]string{
		"sub": "1234567890",
		"exp": "3422335555",
		"name": "John Doe",
		"admin": "1",
		"info": "232323ssdgerere3335dssss",
	}
	ret, _ = json.Marshal(payload)
	input  = []byte(string(ret))

	// base64编码
	payloadString := base64.StdEncoding.EncodeToString(input)
	fmt.Println(headerString+"."+payloadString)

	// 签证
	// HMACSHA256(base64UrlEncode(header) + "." + base64UrlEncode(payload), secret)
	secret := "secret-insecure-(_+qtd5edmhm%2rdsg+qc3wi@s_k*3cbk-+k2gpg3@qx)z6r+p"
	sign   := fmt.Sprintf("%s.%s.%s", headerString, payloadString, secret)
	h := sha256.New()
	h.Write([]byte(sign))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	fmt.Printf("%s.%s.%s", headerString, payloadString, signature)
}
