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
	RoomId      int64
	CreatorName string
	IsNeedPass  bool
	Creatime    int64
}

type GameRoomJoinReq struct {
	RoomId   int64
	PlayerId int64
	RoomPass string
	ClassId  int64
}

type GameRoomJoinRsp struct {
	RoomId  int64
	Players interface{}
}
type GameRoomExitReq struct {
	RoomId   int64
	PlayerId int64
	ClassId  int64
}
