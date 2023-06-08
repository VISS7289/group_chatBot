package EncryptPassword

import (
	"crypto/md5"
	"encoding/hex"
)

const secret = "7289978DUAN"

// 密码md5加密
func EP(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}
