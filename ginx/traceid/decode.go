package traceid

import (
	"encoding/hex"
	"net"
	"strconv"
	"time"
)

type TraceID struct {
	IPAddr     string    `json:"ip_addr"`
	Time       time.Time `json:"time"`
	TraceIndex int       `json:"trace_index"`
	Pid        int       `json:"pid"`
}

func DecodeTraceID(v string) *TraceID {
	if v == "" {
		return nil
	}
	ips := v[0:8]
	ts := v[8:21]
	index, _ := strconv.Atoi(v[21:25])
	pid, _ := strconv.Atoi(v[25:])

	ipb, _ := hex.DecodeString(ips)
	tst, _ := strconv.ParseInt(ts, 10, 64)

	return &TraceID{
		IPAddr:     net.IP(ipb).String(),
		Time:       time.UnixMilli(tst),
		TraceIndex: index,
		Pid:        pid,
	}
}
