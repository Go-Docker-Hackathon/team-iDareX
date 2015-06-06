package qiniu

type ReturnJson struct {
	Success string
	Data    interface{}
}

type ReturnSuccess struct {
	Hash string
	Key  string
}
