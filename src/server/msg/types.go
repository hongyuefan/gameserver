package msg

import (
	"errors"
)

type BussTypeId uint32

const (
	Buss_Verify_Code BussTypeId = iota
	Buss_RegistAndLogin_Code
	Buss_GameClassGet_Code
	Buss_Chat_Code
	Buss_GameRoomAdd_Code
	Buss_GameRoomGet_Code
	Buss_GameRoomJoin_Code
	Buss_GameRoomExit_Code
)

var (
	Err_VerificationCode_TimeOut = errors.New("Verification Code TimeOut")
	Err_VerificationCode_Wrong   = errors.New("Verification Code Wrong")
	Err_Token_TimeOut            = errors.New("Token TimeOut")
	Err_Login_NotExist           = errors.New("User Name or Pass wrong")
)
