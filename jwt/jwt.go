package jwt

import (
	"context"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenType int

const timeLayout = "2006-01-02 15:04:05"

const (
	AccessToken TokenType = iota + 1
	RefreshToken
)

type Claims struct {
	Typ TokenType `json:"typ"`
	UID int       `json:"id"`
	*jwt.RegisteredClaims
}

type JWT struct {
	typ    TokenType
	secret []byte
	issuer string
	expire time.Duration

	keyFunc jwt.Keyfunc
}

func New(sec []byte, issuer string, expire time.Duration) *JWT {
	return &JWT{
		secret: sec,
		issuer: issuer,
		expire: expire,
		// 创建缓存的 keyFunc
		keyFunc: func(t *jwt.Token) (interface{}, error) {
			return sec, nil
		},
	}
}

func (j *JWT) Typ(typ TokenType) *JWT {
	j.typ = typ
	return j
}

func (j *JWT) GenerateWithContext(ctx context.Context, id int) (string, error) {
	// 记录当前时间
	t := time.Now()
	// 计算授权超时
	expireTime := t.Add(j.expire)

	claims := &Claims{
		j.typ,
		id,
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			Issuer:    j.issuer,
		},
	}
	// 不要在基础代码里写 log， 统一交给调用方
	// log.WithCtx(ctx).Info("Issuer new token",
	// 	log.String("type", typName(j.typ)),
	// 	log.Int("id", id),
	// 	// log.String("username", username),
	// 	log.String("expired", expireTime.Format(timeLayout)),
	// )

	// 计算出 Claims 对象
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString(j.secret)
}

func (j *JWT) Generate(id int) (string, error) {
	return j.GenerateWithContext(context.Background(), id)
}

func (j *JWT) Parse(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, j.keyFunc)

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func TypName(typ TokenType) string {
	if typ == AccessToken {
		return "access"
	}
	return "refresh"
}
