package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"order-go/models"
	"strconv"
	"time"
)

const AsciiLetters = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

/*
生成16位的随机密钥
*/
func GenerateKey() string {
	bytes := []byte(AsciiLetters)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 16; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

/*
Md5加密
*/
func Md5Encrypt(pwd, salt string) string {
	md5Ctx := md5.New()                            //md5 init
	md5Ctx.Write([]byte(pwd + "-" + salt))         //md5 updata
	cipherStr := md5Ctx.Sum(nil)                   //md5 final
	encryptedData := hex.EncodeToString(cipherStr) //hex_digest
	return encryptedData
}

/*
生成授权码
*/
func GenerateAuthCode(userInfo models.User) string {
	md5Ctx := md5.New()
	rawStr := strconv.Itoa(userInfo.Uid) + userInfo.LoginName + userInfo.LoginPwd + userInfo.LoginSalt
	md5Ctx.Write([]byte(rawStr))              //md5 updata
	cipherStr := md5Ctx.Sum(nil)              //md5 final
	authCOde := hex.EncodeToString(cipherStr) //hex_digest
	return authCOde
}
