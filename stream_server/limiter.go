package stream_server

import "log"

type ConnLimiter struct {
	currentConn int
	bucket chan int
}

func NewConnLimiter(cc int) *ConnLimiter{
	return &ConnLimiter{
		currentConn:cc,
		bucket:make(chan int,cc),
	}
}

func (c *ConnLimiter) GetCon()bool{
	if len(c.bucket) >= c.currentConn{
		log.Print("Limiter buffer is finished")
		return false
	}
	c.bucket <- 1
	return true
}

func (c *ConnLimiter) ReleaseConn(){
	cc := <-c.bucket
	log.Println("new connection %s",cc)
}