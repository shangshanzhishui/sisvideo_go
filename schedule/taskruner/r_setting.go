package taskruner

type ControlChan chan string

type DataChan chan interface{}
type Function func(d DataChan) error

const (
	READY_DISPATCH = "d"
	READY_EXECUTE = "e"
	CLOSE = "c"

	VIDEO_DICTORY = "../../videos/" //shipinlujing
)


