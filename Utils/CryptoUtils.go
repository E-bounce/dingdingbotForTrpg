package Utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strconv"
	"time"
)


func GetTimestamp()int{
	return int(time.Now().UnixNano()/1e6)
}

func GetSignatureSubmit(timestamp int) string {
	signString := strconv.Itoa(timestamp)+"\n"+Userconfig["secretkey"]
	key := []byte(Userconfig["secretkey"])
	hashConstructor := hmac.New(sha256.New,key)
	hashConstructor.Write([]byte(signString))
	cryptoText:=hashConstructor.Sum(nil)
	return base64.StdEncoding.EncodeToString([]byte(cryptoText))
}

func GetSignatureCheck(timestamp int)string{
	signString := strconv.Itoa(timestamp)+"\n"+Userconfig["AppKey"]
	key := []byte(Userconfig["AppKey"])
	hashConstructor := hmac.New(sha256.New,key)
	hashConstructor.Write([]byte(signString))
	cryptoText:=hashConstructor.Sum(nil)
	return base64.StdEncoding.EncodeToString([]byte(cryptoText))
}

func checkTimestamp(get_time int) bool {
	return (GetTimestamp()-get_time) < 3600000
}

func checkSign(get_time int,get_sign string) bool {
	return GetSignatureCheck(get_time) == get_sign
}

func CheckAvailable(get_time int,get_sign string) bool {
	return checkTimestamp(get_time) && checkSign(get_time,get_sign)
}