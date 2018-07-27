package msg

func init() {
	Processor.Register(&GameRoomAdd{})
	Processor.Register(&GameRoomGetReq{})
	Processor.Register(&GameRoomJoinReq{})
	Processor.Register(&GameRoomExitReq{})
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

type GameRoomJoinReq struct {
	RoomName string
	PlayerId int64
	RoomPass string
	ClassId  int64
}

type GameRoomExitReq struct {
	RoomName string
	PlayerId int64
	ClassId  int64
}
