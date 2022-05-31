package main

import (
	"fmt"
	"mypb/simple"

	"google.golang.org/protobuf/proto"
)

func main() {
	sMsg := &simple.SimplestMessage{Msg: "My Simplest message!"}
	SerializeToBytes(sMsg)
	simplePB := &simple.SimpleMessage{
		Id:        1,
		Msg:       "Have fun!",
		Sender:    "Logesh",
		Receivers: []string{"To me", "To all"},
	}
	SerializeToBytes(simplePB)
}

func SerializeToBytes(pb proto.Message) []byte {
	fmt.Println("Serializing the message to byte slice")
	fmt.Printf("Message      : %v\n", pb)
	fmt.Printf("Message type : %T\n", pb)
	msgByteSlice, err := proto.Marshal(pb)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	fmt.Println("Serialized byte slice")
	fmt.Println(msgByteSlice)
	return msgByteSlice
}
