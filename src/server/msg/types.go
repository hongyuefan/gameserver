package msg

import (
	"errors"
	"server/util/arithmetic"
)

var (
	Flag_Game_Start int32 = 1
	Flag_Game_Over  int32 = 0
)

type BussTypeId uint32

const (
	Buss_Verify_Code BussTypeId = iota
	Buss_RegistAndLogin_Code
	Buss_GameClassGet_Code
	Buss_GameRoomAdd_Code
	Buss_GameRoomGet_Code
	Buss_GameRoomJoin_Code
	Buss_GameRoomExit_Code
	Buss_GameTableJoin_Code
	Buss_GameTableExit_Code
	Buss_Chat_Code
)

type CardTypeId uint32

const (
	Card_Type_Jan CardTypeId = iota
	Card_Type_Ken
	Card_Type_Po
)

func (m CardTypeId) GetRand() CardTypeId {
	var seed int
	seed = arithmetic.GetRand(0, 3)
	switch seed {
	case 0:
		return Card_Type_Jan
	case 1:
		return Card_Type_Ken
	case 2:
		return Card_Type_Po
	}
	return Card_Type_Po
}

var (
	Err_VerificationCode_TimeOut = errors.New("Verification Code TimeOut")
	Err_VerificationCode_Wrong   = errors.New("Verification Code Wrong")
	Err_Token_TimeOut            = errors.New("Token TimeOut")
	Err_Login_NotExist           = errors.New("User Name or Pass wrong")
)
