package jwt_test

import (
	"testing"
	"time"

	"github.com/sendya/pkg/jwt"
)

func TestNew(t *testing.T) {
	type args struct {
		sec    []byte
		issuer string
		expire time.Duration
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Generate and Parse",
			args: args{
				sec:    []byte("mykey"),
				issuer: "myapp",
				expire: time.Minute * 10, // 10 分钟过期
			},
			want: 10000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			signed := jwt.New(tt.args.sec, tt.args.issuer, tt.args.expire).Typ(jwt.AccessToken)

			token, _ := signed.Generate(tt.want)
			claims, _ := signed.Parse(token)

			if tt.want != claims.UID {
				t.Errorf("sigend token = %s, got %d want %d", token, claims.UID, tt.want)
			}
		})
	}
}

func TestParse(t *testing.T) {
	// gen default config
	sec := []byte("wLDFa7wNG854YA8N")
	sec2 := []byte("xxxx")
	issuer := "myapp"
	expire := time.Minute * 10
	uid := 10000

	jwtx := jwt.New(sec, issuer, expire)
	jwtx2 := jwt.New(sec2, issuer, expire)

	token, err := jwtx.Generate(uid)
	if err != nil {
		t.Error(err)
	}

	// 必须 转换成功
	if c, err := jwtx.Parse(token); err != nil || c.UID != uid {
		t.Error("prase token error", err)
	}

	// 密钥不同，必须 转换失败
	if c, err := jwtx2.Parse(token); err == nil {
		t.Error(c)
	}
}
