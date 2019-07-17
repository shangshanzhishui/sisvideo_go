package taskruner

import (
	"errors"
	"log"
	"os"
	"sis_video_go/schedule/sch_db"
	"sync"
)

func ClearVideoDispatcher(d DataChan)error{
	log.Println("ClearVideoDispatcher")
	res,err := sch_db.ReadDelVideoRecord(3)
	if err!= nil{
		log.Println(err)
		return err
	}
	if len(res) ==0{
		return errors.New("all videos already is deleted")
	}
	for _,id := range res{
		d <- id
	}
	log.Println("ClearVideoDispatcher")
	return nil
}

func DelVideo(vid string) error{
	err:= os.Remove(VIDEO_DICTORY+vid)
	log.Println(err)
	if err !=nil && !os.IsNotExist(err){
		log.Println(err)
		return err
	}
	return nil
}

func ClearVideoExecuter(d DataChan) error{
	log.Println("ClearVideoExecuter start")
	errMap := &sync.Map{}
	out:
		for{
			select{
			case vid := <- d:
				go func(id interface{}){
					if err := DelVideo(id.(string));err !=nil{
						errMap.Store(id.(string),err)
						return
					}
					if err := sch_db.DelVideoDelRecord(id.(string));err != nil{
						log.Println(err)
						errMap.Store(id.(string),err)
						return
					}
			}(vid)
			default:
				break out


			}
		}
	var err error
	errMap.Range(func (k,v interface{}) bool {
		err = v.(error)
		if v !=nil{
			return false
		}
		return true
	})
	log.Println("ClearVideoExecuter")
	return err
}
