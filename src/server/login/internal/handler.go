package internal

import (
	"fmt"
	"reflect"
	"server/conf"
	db "server/database/mysqlbase"
	agent "server/manage_agent"
	"server/msg"
	"server/util/arithmetic"
	"strings"
	"time"

	"github.com/name5566/leaf/gate"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {
	handleMsg(&msg.PlayerRegist{}, handleRegist)
	handleMsg(&msg.SendIdentifyCode{}, handleSendIdentifyCode)
	handleMsg(&msg.PlayerLogin{}, handleLogin)
}

func handleSendIdentifyCode(args []interface{}) {
	var (
		err error
	)
	m := args[0].(*msg.SendIdentifyCode)

	vCode := arithmetic.GenCode(4)

	if strings.Contains(m.MobileOrEmail, "@") {
		if err = skeleton.SendEmail(conf.Server.Email_From_Name, m.MobileOrEmail, "", "", vCode); err != nil {
			goto errDeal
		}
	} else {
		if err = skeleton.SendMsg("86", m.MobileOrEmail, []string{vCode}); err != nil {
			goto errDeal
		}
	}
	skeleton.SetSession(m.MobileOrEmail, vCode)
	msg.SuccessHandler(args[1].(gate.Agent), msg.Buss_Verify_Code, "")
	return
errDeal:
	msg.FailedHandler(args[1].(gate.Agent), err)
}

func handleLogin(args []interface{}) {
	var (
		err   error
		token string
		play  *db.GamePlayer
	)
	m := args[0].(*msg.PlayerLogin)

	play = &db.GamePlayer{MobileOrEmail: m.MobileOrEmail, Password: m.Password}

	if err = db.GetPlayerBy(play, "MobileOrEmail", "Password"); err != nil {
		err = msg.Err_Login_NotExist
		goto errDeal
	}

	agentAdd(play.Id, args[1].(gate.Agent))

	if token, err = skeleton.TokenGen(fmt.Sprintf("%v", play.Id)); err != nil {
		goto errDeal
	}

	msg.SuccessHandler(args[1].(gate.Agent), msg.Buss_RegistAndLogin_Code, token)
	return
errDeal:
	msg.FailedHandler(args[1].(gate.Agent), err)
}

func handleRegist(args []interface{}) {
	var (
		err   error
		play  *db.GamePlayer
		pId   int64
		token string
	)
	m := args[0].(*msg.PlayerRegist)

	code := skeleton.GetSession(m.MobileOrEmail)

	if code == nil {
		err = msg.Err_VerificationCode_TimeOut
		goto errDeal
	} else {
		if 0 != strings.Compare(m.VerifyCode, code.(string)) {
			err = msg.Err_VerificationCode_Wrong
			goto errDeal
		}
	}

	play = &db.GamePlayer{
		MobileOrEmail: m.MobileOrEmail,
		Password:      m.Password,
		Createtime:    time.Now().Unix(),
		Nickname:      arithmetic.GenCode(8),
	}

	if pId, err = db.AddPlayer(play); err != nil {
		goto errDeal
	}

	agentAdd(play.Id, args[1].(gate.Agent))

	if token, err = skeleton.TokenGen(fmt.Sprintf("%v", pId)); err != nil {
		goto errDeal
	}
	msg.SuccessHandler(args[1].(gate.Agent), msg.Buss_RegistAndLogin_Code, token)
	return
errDeal:
	msg.FailedHandler(args[1].(gate.Agent), err)
}

func agentAdd(k int64, g gate.Agent) {
	agent.MAgent.InsertAgent(k, g)
	g.SetUserData(k)
}
