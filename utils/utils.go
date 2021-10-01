package utils

import "log"

func HandlErr(err error){
	if err != nil{
		log.Panic(err)
	}
}