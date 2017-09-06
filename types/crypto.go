package types

import (
    "crypto/md5"
    "encoding/hex"
)

func Md5String(str string) string {
    md5Ctx := md5.New()
    md5Ctx.Write([]byte(str))
    cipherStr := md5Ctx.Sum(nil)
    return hex.EncodeToString(cipherStr)
}