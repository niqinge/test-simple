package runtime

type BaseContext interface {
	Schema() string
	GReq() interface{}
	SResult(interface{})
	GResult() interface{}
}

type PageContext interface {
	Schema() string
	GReq() interface{}
	GSize() uint
	GPageNo() uint
	SResult(interface{})
	GResult() interface{}
}
