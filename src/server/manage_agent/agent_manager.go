package manage_agent

import (
	"github.com/name5566/leaf/gate"
)

var MAgent *gate.MAgent

func init() {
	MAgent = gate.NewMAgent()
}
