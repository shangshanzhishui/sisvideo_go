package taskruner

import "log"

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
	log.Println("22222222222222222222222222")

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
			log.Println("11111111111111111111111111111111111")
			log.Printf("..........................%v",c)
			if c==READY_DISPATCH{//开始分发
			log.Println("kaishi fenfa renwu")
				err:= r.Dispather(r.Data)
				if err!=nil{
					log.Println(err)
					r.Error <- CLOSE
				}else{
					r.Controller <- READY_EXECUTE
				}


			}

			if c == READY_EXECUTE{
				err := r.Executor(r.Data)
				if err != nil{

					r.Error <- CLOSE
				}else{
					r.Controller <- READY_DISPATCH
				}


			}
		case e:= <- r.Error:
			if e == CLOSE{
				log.Println("renwu wancheng")
				return
			}
		default:


		}
	}
}

func (r *Runner)Start(){

	r.Controller <- READY_DISPATCH
	log.Printf("kaishi fenfa %v",r.Controller)
	r.StartDispatch()
}