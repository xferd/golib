package crypt

import (
    "crypto/md5"
    "encoding/hex"
)

func MD5(cipher []byte) string {
    h := md5.New()
    h.Write(cipher) // 需要加密的字符串为 123456
    cipherStr := h.Sum(nil)
    return hex.EncodeToString(cipherStr) // 输出加密结果
}
