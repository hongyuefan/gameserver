package msg

func init() {
	Processor.Register(&GameClassGetReq{})
}

type GameClassGetReq struct {
}

type GameClassGetRsp struct {
	ClassId   int64
	ClassName string
}
