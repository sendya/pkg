# go pkg

业务用扩展包

## 所有包

### encode

> MD5 计算, MD5加盐密码

### env

> 编译环境变量，运行时变量

### Ginx

> gin-gonic req/resp 的封装，middleware 封装

- ginlog  # gin-logger with zap
- resp    # resp.OK(ctx) / resp.JSON(ctx, interface{}) 
- traceid # 请求跟踪链中间件

### Json

> json 使用了 std json 和 jsoniter ，编译时可以使用 -tags=jsoniter 来使用 jsoniter

### Log

> log 使用 zap log 进行二次封装，加装了 context 的能力

### Rand

> 随机包

#### str 字符串随机

## 编译附加信息

### 基本使用方式
```
-ldflags="-w -extldflags=-static \
    -X github.com/sendya/pkg/env.Version=${BUILD_VER} \
    -X github.com/sendya/pkg/env.Githash=${BUILD_HASH} \
    -X github.com/sendya/pkg/env.OSArch=${BUILD_ARCH} \
    -X github.com/sendya/pkg/env.Built=${BUILD_TIME} \
    " 
```
