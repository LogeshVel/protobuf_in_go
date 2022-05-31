package main

import (
	"fmt"
	"log"
	"mypb/simple"

	"google.golang.org/protobuf/proto"
)

func main() {
	simplePB := &simple.SimpleMessage{
		Id:        1,
		Msg:       "Have fun!",
		Sender:    "Logesh",
		Receivers: []string{"To me", "To all"},
	}
	fmt.Println(simplePB)
	simpleByteSlice, err := proto.Marshal(simplePB)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(simpleByteSlice)
}
