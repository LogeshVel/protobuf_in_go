package main

import (
	"fmt"
	"mypb/simple"

	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	simplePB := &simple.SimpleMessage{
		Id:        1,
		Msg:       "Have fun!",
		Sender:    "Logesh",
		Receivers: []string{"To me", "To all"},
	}
	fmt.Println("Message format")
	fmt.Println(simplePB)

	simpleJSON, _ := protojson.Marshal(simplePB)
	fmt.Println("\nJSON format in Bytes Slice")
	fmt.Println(simpleJSON)
	fmt.Println("\nJSON format")
	fmt.Println(string(simpleJSON))

	option := protojson.MarshalOptions{Multiline: true}
	simpleJSON, _ = option.Marshal(simplePB)
	fmt.Println("\nJSON format in Bytes Slice with option set to Multiline")
	fmt.Println(simpleJSON)
	fmt.Println("\nJSON format")
	fmt.Println(string(simpleJSON))
}
