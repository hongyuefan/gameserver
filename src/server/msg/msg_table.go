package msg

func init() {
	Processor.Register(&GameTableJoinReq{})
	Processor.Register(&GameTableExitReq{})
}

type GameTableJoinReq struct {
	ClassId  int64
	RoomId   int64
	PlayerId int64
	TableId  int64
}

type GameTableJoinRsp struct {
	TableId  int64
	PlayerId int64
}

type GameTableExitReq struct {
	ClassId  int64
	RoomId   int64
	PlayerId int64
	TableId  int64
}

type GameTableExitRsp struct {
	TableId  int64
	PlayerId int64
}

type GameTablePKReq struct {
	TableId  int64
	PlayerId int64
	CardId   CardTypeId
}

type GameTablePKRsp struct {
	PlayerId int64
}
