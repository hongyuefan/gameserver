package msg

func init() {
	Processor.Register(&PlayerRegist{})
	Processor.Register(&PlayerLogin{})
}

type PlayerRegist struct {
	MobileOrEmail string
	Password      string
	VerifyCode    string
}

type PlayerLogin struct {
	MobileOrEmail string
	Password      string
}

type PlayerGetReq struct {
	PlayerId int64
}
type PlayerGetRsp struct {
	Stars uint32
	Gold  uint64
}
