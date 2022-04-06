package traceid

import (
	"encoding/hex"
	"fmt"
	"net"
	"strconv"
	"strings"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/sendya/pkg/log"
)

type Counter struct {
	c     int32
	begin int32
	max   int32
}

var (
	counter *Counter
	ip      string = ""
	pid     int    = 0
	pidStr  string = strconv.Itoa(pid)
)

/*
e.g. 0ad1348f1403169275002100356696

前 8 位 0ad1348f 即产生 TraceId 的机器的 IP，这是一个十六进制的数字
每两位代表 IP 中的一段，我们把这个数字，按每两位转成 10 进制即可得到常见的 IP 地址表示方式 10.209.52.143，
您也可以根据这个规律来查找到请求经过的第一个服务器。
后面的 13 位 1403169275002 是产生 TraceId 的时间。
之后的 4 位 1003 是一个自增的序列，从 1000 涨到 9000，到达 9000 后回到 1000 再开始往上涨。
最后的 5 位 56696 是当前的进程 ID，为了防止单机多进程出现 TraceId 冲突的情况，所以在 TraceId 末尾添加了当前的进程 ID。
*/
func New() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := newID()
		ctx, _ := log.WithFields(c.Request.Context(), zap.String("traceId", traceID))
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}

func init() {
	pid = syscall.Getpid()
	pidStr = strconv.Itoa(pid)
	ip = getHexIP()
	counter = &Counter{
		c:     1000,
		begin: 1000,
		max:   9999,
	}
}

func newID() string {
	var id strings.Builder

	ts := time.Now().UnixMilli()
	index := counter.Get()

	id.WriteString(ip)
	id.WriteString(fmt.Sprint(ts))
	id.WriteString(fmt.Sprint(index))
	id.WriteString(pidStr)

	return id.String()
}

func getHexIP() string {
	ifaces, _ := net.Interfaces()
	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			continue
		}
		// handle err
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}

			return hex.EncodeToString(ip.To4())
		}
	}
	return "00000000"
}

func (counter *Counter) Get() int32 {
	n := atomic.AddInt32(&counter.c, 1)
	if n >= 9000 {
		return atomic.SwapInt32(&counter.c, counter.begin)
	}
	return n
}

func (counter *Counter) Incr() int32 {
	return atomic.AddInt32(&counter.c, 1)
}
