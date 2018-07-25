package msg

func init() {
	Processor.Register(&GameRoomAdd{})
}

type GameRoomAdd struct {
	GameClassId int64
	CreatorId   int64
	RoomPass    string
}

type GameRoomGet struct {
	ClassId  int64
	RoomId   int64
	RommName string
}
