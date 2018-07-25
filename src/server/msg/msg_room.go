package msg

func init() {
	Processor.Register(&GameRoomAdd{})
}

type GameRoomAdd struct {
	GameClassId int64
	CreatorId   int64
	RoomPass    string
}

type GameRoomGetReq struct {
	ClassId int64
}

type GameRoomGetRsp struct {
	RoomName    string
	CreatorName string
	IsNeedPass  bool
	Creatime    int64
}
