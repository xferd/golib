package redis

import (
    "net"
    "strconv"
    "github.com/xferd/golib/types"
)

type redisServer struct {
    ip string
    port int
    auth string
}

var (
    connPool map[string]*net.TCPConn
)

func init() {
    connPool = make(map[string]*net.TCPConn)
}

func NewRedisServer(ip string, port int, auth string) *redisServer {
    return &redisServer{ip, port, auth}
}

func (this *redisServer)getConn() (*net.TCPConn, error) {
    key := types.Md5String(this.ip + strconv.Itoa(this.port))
    if conn, ok := connPool[key]; ok {
        return conn, nil
    } else {
        tcpAddr, err := net.ResolveTCPAddr("tcp4", this.ip + ":" + strconv.Itoa(this.port))
        c, err := net.DialTCP("tcp", nil, tcpAddr)
        if nil != err {
            return nil, err
        }
        connPool[key] = c
        return c, nil
    }
}

func (this *redisServer)request(cmd string) (<- chan string, error) {
    c, err := this.getConn()
    if nil != err {
        return nil, err
    }
    // response := make(chan string)
    resp = response
    go func() {
        c.Write([]byte(cmd))
        for {
            buf := make([]byte, 0, 4096)
            n, e := c.Read(buf)
            if nil != e {
                break;
            }

            response <- string(buf[:n])
        }
    }()
    return response, nil
}
