package msg

func init() {
	Processor.Register(&GameClassGet{})
}

type GameClassGet struct {
}

type GameClass struct {
	ClassId   int64
	ClassName string
}
