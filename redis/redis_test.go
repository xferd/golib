package redis

import (
    "testing"
)

func Test_RedisConnection(t *testing.T) {
    r := NewRedisServer("127.0.0.1", 6379, "")
    if conn, err := r.getConn(); nil != err {
        t.Error(err)
    } else {
        t.Log(conn)
    }
}

func Test_Send(t *testing.T) {
    r := NewRedisServer("127.0.0.1", 6379, "")
    resp, _ := r.request("*3\r\n$3\r\nset\r\n$1\r\na\r\n$1\r\n1\r\n")
    t.Log(<- resp)
}