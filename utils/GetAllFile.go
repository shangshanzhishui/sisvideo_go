package utils

import (
	"io/ioutil"
	"log"
)

func GetAllFil(path string) error{
	r,err := ioutil.ReadDir(path)

	for _ , f := range r{
		if f.IsDir(){
			log.Println("[%s]\n",path+f.Name())
			GetAllFil(path+f.Name())
		}
		log.Println(f.Name())
	}
	return err
}
