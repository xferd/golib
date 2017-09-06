package redis

import (
    "log"
)

const (
    ErrorReply  = '-'
    StatusReply = '+'
    IntReply    = ':'
    StringReply = '$'
    ArrayReply  = '*'
)

type response struct {
    replyType rune
    content interface{}

    buf []byte
}

func (this *response)Parse(seg []byte) (finished bool){
    if nil == this.buf {
        this.buf = make([]byte, 0, 4096)
    }
    log.Println(this.replyType)
    return false
}