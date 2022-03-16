package md5_test

import (
	"reflect"
	"testing"

	"github.com/sendya/pkg/encode/md5"
)

func TestSum(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "test-sum-md5",
			args: args{
				v: "testpwd1",
			},
			// `testpwd1` sum
			want: []byte{170, 28, 94, 171, 83, 94, 188, 144, 24, 162, 42, 237, 37, 162, 97, 29},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := md5.Sum(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("md5.Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSums(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test-pwd-gen",
			args: args{v: "mypwd-test"},
			// echo -n "mypwd-test" |md5sum |cut -d" " -f1
			want: "7bd97c6665ec11fc45d390d4b61920d9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := md5.Sums(tt.args.v); got != tt.want {
				t.Errorf("md5.Sums() = %v, want %v", got, tt.want)
			}
		})
	}
}
