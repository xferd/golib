package redis

import (
    "net"
    "strconv"
    "github.com/xferd/golib/types"
    "fmt"
    "io"
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
    response := make(chan string)
    go func() {
        c.Write([]byte(cmd))
        for {
            buf := make([]byte, 0, 4096)
            n, e := c.Read(buf)
            fmt.Println(n, e, buf[:n])
            if io.EOF == e {
                break;
            }
            // response <- string(buf[:n])
        }
    }()
    return response, nil
}

func buildCmd(cmd []types.Any) string {
    str := fmt.Sprintf("*%d\r\n", len(cmd))
    for _, c := range cmd {
        switch c.(type) {
        case string:
            str += fmt.Sprintf("$%d\r\n%s\r\n", len(c.(string)), c)
        // case int:
        //     var ilen int = 0
        //     for tmp := c.(int); ; tmp /= 10 {ilen++}
        //     str += fmt.Sprintf(":%d\r\n%d\r\n", ilen, c.(int))
        default:
            panic("unknown cmd type, ")
        }
    }
    return str
}

func (this *redisServer)Get(key string) string {
    cmd := []types.Any{"GET", key}
    cmdString := buildCmd(cmd)
    response, _ := this.request(cmdString)
    var respStr string
    for r := range response {
        fmt.Print(r)
    }
    return respStr
}
