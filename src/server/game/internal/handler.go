package internal

import (
	"reflect"
	db "server/database/mysqlbase"
	"server/msg"

	"github.com/name5566/leaf/gate"
	//	"github.com/name5566/leaf/log"
)

func init() {
	handler(&msg.GameClass{}, handlerGetGameClass)
	handler(&msg.RoomCreate{}, handlerCreateGameRoom)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handlerGetGameClass(args []interface{}) {

	var (
		gCs []msg.GameClass
		gC  msg.GameClass
		err error
		ml  []interface{}
	)

	a := args[1].(gate.Agent)

	query := make(map[string]string, 0)

	if ml, _, err = db.GetAllGameClass(query, []string{}, []string{"game_id"}, []string{"desc"}, 0, 100); err != nil {
		goto errDeal
	}
	for _, v := range ml {
		gC.Id = v.(db.GameClass).Id
		gC.GameName = v.(db.GameClass).GameName
		gC.GamePlayerCount = v.(db.GameClass).GamePlayerCount

		gCs = append(gCs, gC)
	}
	msg.SuccessHandler(a, msg.Buss_GetGameClass_Code, gCs)
	return
errDeal:
	msg.FailedHandler(a, err)
}

func handlerCreateGameRoom(args []interface{}) {
	m := args[0].(*msg.RoomCreate)

}
