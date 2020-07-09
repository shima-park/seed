package executor

import (
	"fmt"
)

type Executor interface {
	Start() error
	Stop()
	State() State
}

var (
	Idle    State = 0
	Running State = 1
	Exited  State = 3
)

type State int32

func (s State) String() string {
	switch s {
	case Idle:
		return "idle"
	case Running:
		return "running"
	case Exited:
		return "exited"
	}
	return fmt.Sprintf("unknown(%d)", s)
}
