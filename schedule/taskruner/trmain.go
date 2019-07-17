package taskruner

import (
	"log"
	"time"
)

type Work struct {
	ticker *time.Ticker
	runner *Runner
}

func NewWorker(interval time.Duration,r *Runner) *Work{
	return &Work{
		ticker:time.NewTicker(interval * time.Second),
		runner:r,
	}

}



func(w * Work) StartWork(){
	for {
		select {
			case <- w.ticker.C:
				go w.runner.Start()
				log.Println(123)

		}
	}
}

//start clear video task
func WorkStart(){
	r := NewRunner(true,3,ClearVideoExecuter,ClearVideoDispatcher)
	w := NewWorker(3,r)
	go w.StartWork()

}