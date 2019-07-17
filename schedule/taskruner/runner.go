package taskruner

type Runner struct {
	Controller ControlChan
	Error ControlChan
	Data DataChan
	Longlife bool
	Dispather Function
	Executor Function
}


func NewRunner(longlife bool,datasize int,executor Function,dispather Function) *Runner{
	return &Runner{
		Controller:make(chan string,1),
		Error:make(chan string,1),
		Data:make(chan interface{},datasize),
		Longlife:longlife,
		Dispather:dispather,
		Executor:executor,
	}
}

func (r *Runner) StartDispatch(){
	defer func() {
		if !r.Longlife{
			close(r.Controller)
			close(r.Data)
			close(r.Error)
		}
	}()


	for{
		select {
		case c:= <- r.Controller:
			if c==READY_DISPATCH{//开始分发
				if err:= r.Dispather(r.Data);err!=nil{
					r.Error <- CLOSE
				}else {
					r.Controller <- READY_EXECUTE
				}
			}

			if c == READY_EXECUTE{
				if err := r.Executor;err != nil{

					r.Error <- CLOSE
				}else{
					r.Controller <- READY_DISPATCH
				}
			}
		case e:= <- r.Error:
			if e == CLOSE{
				return
			}
		default:


		}
	}
}

func (r *Runner)Start(){
	r.Controller <- READY_DISPATCH
	r.StartDispatch()
}