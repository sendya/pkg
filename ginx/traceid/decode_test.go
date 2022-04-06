package traceid

import (
	"reflect"
	"testing"
	"time"
)

func TestDecodeTraceID(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want *TraceID
	}{
		{name: "basic_test", args: args{v: "ac1a498e1649234654302110716286"}, want: &TraceID{
			IPAddr:     "172.26.73.142",
			Time:       time.UnixMilli(1649234654302),
			TraceIndex: 1107,
			Pid:        16286,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeTraceID(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeTraceID() = %v, want %v", got, tt.want)
			}
		})
	}
}
