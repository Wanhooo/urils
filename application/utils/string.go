package utils

import (
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
	. "urils/application/constants"
)

func Random(n int, chars string) string {
	if n <= 0 {
		return ""
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	bytes := make([]byte, n, n)
	l := len(chars)
	for i := 0; i < n; i++ {
		bytes[i] = chars[r.Intn(l)]
	}
	return string(bytes)
}


/**
生成指定长度的字符串(字母)
*/

func RandomLetters(n int) string {
	return Random(n, LETTERS)
}


/**
生成指定长度的字符串(数字)
*/

func RandomNumeric(n int) string {
	return Random(n, NUMBERS)
}


/**
生成指定长度的字符串(数字+字母)
*/

func RandomLettersNumeric(n int) string {
	return Random(n, LETTERS_NUMBERIC)
}


/**
生成指定长度的字符串(字母+数字+特殊字符)
*/

func RandomAscii(n int) string {
	return Random(n, ASCII)
}


/*
 加密密码
*/

func MakeHashPassword(RawPassword string) (HashPasswrd string, err error) {
	pwd := []byte(RawPassword)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return
	}
	HashPasswrd = string(hash)
	return
}

/**
 验证密码
*/

func CheckPassword(HashPassword string, RawPassword string) bool {
	ByteHash := []byte(HashPassword)
	BytePwd := []byte(RawPassword)
	err := bcrypt.CompareHashAndPassword(ByteHash, BytePwd)
	if err != nil {
		return false
	}
	return true
}
