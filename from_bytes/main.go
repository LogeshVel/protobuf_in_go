package main

import (
	"fmt"
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
	fmt.Printf("Type of simplePB : %T\n", simplePB)
	simpleByteSlice, _ := proto.Marshal(simplePB)
	fmt.Printf("Type of simpleByteSlice : %T\n", simpleByteSlice)
	fmt.Printf("simpleByteSlice : %v\n\n", simpleByteSlice)
	// Unmarshal (Deserialize)
	fmt.Println("Deserialize...")
	decodedSimplePB := &simple.SimpleMessage{}
	proto.Unmarshal(simpleByteSlice, decodedSimplePB)
	fmt.Println(decodedSimplePB)
	fmt.Printf("Type of decodedSimplePB : %T\n", decodedSimplePB)

}
