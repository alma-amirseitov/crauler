package main

import (
	"encoding/json"
	"log"
)

type Info struct {
	Message string
}


func main(){
	InfoMsg := &Info{
		Message: "logging",
	}
	Data,err := json.Marshal(InfoMsg)
	log.Println(Data,err)
}