package resp

type Resp interface {
	i()                                                                   // i 为了避免被其他包实现
	WithData(data interface{}) Resp                                       // WithData 设置成功时返回的数据
	WithPage(data interface{}, pageSize int, current int, total int) Resp // WithPageData 设置成功时返回的分页结构体
	WithMessage(id string) Resp                                           // WithMessage 设置当前请求的错误信息
}

type RespPage interface {
	GetData() interface{}
	GetTotal() int64
	GetPageSize() int
	GetCurrent() int
}

type resp struct {
	Code    int         `json:"code"`              // 业务编码
	Message string      `json:"message,omitempty"` // 错误描述
	Data    interface{} `json:"data,omitempty"`    // 成功时返回的数据
} //@name Resp

type respPage struct {
	resp
	Current  *int `json:"current,omitempty"`  // 分页::当前页 从 1 开始
	PageSize *int `json:"pageSize,omitempty"` // 分页::每页大小
	Total    *int `json:"total,omitempty"`    // 分页::总记录数
} //@name RespPage

func NewResp(code int, msg string) Resp {
	return &resp{
		Code:    code,
		Message: msg,
		Data:    nil,
	}
}

func (e *resp) i() {}

func (e *resp) WithData(data interface{}) Resp {
	e.Data = data
	return e
}

func (e *resp) WithPage(data interface{}, pageSize int, current int, total int) Resp {
	p := &respPage{}
	p.Code = e.Code
	p.Message = e.Message
	p.Data = data
	p.PageSize = &pageSize
	p.Current = &current
	p.Total = &total
	return p
}

func (e *respPage) WithPage(data interface{}, pageSize int, current int, total int) Resp {
	e.Data = data
	e.PageSize = &pageSize
	e.Current = &current
	e.Total = &total
	return e
}

func (e *resp) WithMessage(msg string) Resp {
	e.Message = msg
	return e
}
