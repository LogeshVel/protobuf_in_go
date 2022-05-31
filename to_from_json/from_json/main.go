package main

import (
	"fmt"
	"mypb/simple"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {
	jsonStr := `{"id":1, "msg":"Have fun!", "sender":"Logesh", "receivers":["To me", "To all"]}`
	fromJSON(jsonStr, &simple.SimpleMessage{})

	jsonStr = `{"id":1, "msg":"Have fun!", "unknown":"Logesh", "receivers":["To me", "To all"]}`
	fromJSON(jsonStr, &simple.SimpleMessage{})

	fromJSONWithOption(jsonStr, &simple.SimpleMessage{})
	simplestMsgStr := `{"msg": "My Simplest Message is to have fun"}`
	fromJSON(simplestMsgStr, &simple.SimplestMessage{})

}

func fromJSON(jsonStr string, pb proto.Message) {
	fmt.Printf("\nGiven JSON string : %v\n", jsonStr)
	err := protojson.Unmarshal([]byte(jsonStr), pb)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Not able to Unmarshal the given bytes of slice")
		fmt.Printf("%v\n\n", pb)
		return
	}
	fmt.Println(pb)
}

func fromJSONWithOption(jsonStr string, pb proto.Message) {
	fmt.Printf("Given JSON string : %v\n", jsonStr)
	fmt.Println("option set to Discard the Unknown field")
	option := protojson.UnmarshalOptions{DiscardUnknown: true}
	err := option.Unmarshal([]byte(jsonStr), pb)
	if err != nil {
		fmt.Println("\n", err)
		fmt.Println("Not able to Unmarshal the given bytes of slice")
		fmt.Printf("%v\n\n", pb)
		return
	}
	fmt.Println(pb)
}
