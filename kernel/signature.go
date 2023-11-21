package kernel

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

func GenerateSalt() string {
	// 获取当前时间戳作为随机数
	salt := fmt.Sprintf("%d", time.Now().Unix())
	return salt
}

// GenerateSignature 生成签名
func GenerateSignature(appID, query, salt, secret string) string {
	// 拼接字符串1
	raw := appID + query + salt + secret

	// 计算MD5
	hasher := md5.New()
	hasher.Write([]byte(raw))
	return hex.EncodeToString(hasher.Sum(nil))
}
