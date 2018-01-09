package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	protohelper "github.com/alecthomas/protobuf"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/marvin5064/protobuf-strip-off/protobuf"
)

func main() {
	CheckOutProtobufProperty()
	fmt.Println("Testing if protobuf can be strip off by define automatically")
	TestP3()
	TestP2ToP3()
	TestP2()
}

func CheckOutProtobufProperty() {
	fmt.Println("CheckOutProtobufProperty ...")
	smallDefine := &protobuf.SmallerP3Define{}

	v := reflect.ValueOf(*smallDefine)
	t := v.Type()
	result := protohelper.ProtoFields(t)
	fmt.Println("result", result)
	for i, r := range result {
		fmt.Println("result - ", i, ":")
		fmt.Println("id - ", r.ID, "; name - ", r.Field.Name)
	}

	fmt.Println("smallDefine", smallDefine)
}
func TestP3() {
	fmt.Println("Testing proto 3")
	proto3Large := &protobuf.LargerP3Define{
		Provider:        1,
		CompetitionId:   2,
		EventId:         3,
		MarketId:        4,
		Outcome:         "5",
		SpecialBetValue: "6",
		Odds:            7.1,
		IsLive:          true,
		CurrencyPair:    "8",
		StakeFactor:     9.1,
	}
	fmt.Println("proto3 large origin:", proto3Large)
	p3Payload, err := proto.Marshal(proto3Large)
	if err != nil {
		fmt.Println("P3 large marshall err:", err)
		return
	}
	fmt.Println("proto3 large binary:", p3Payload)
	proto3Small := &protobuf.SmallerP3Define{}
	err = proto.Unmarshal(p3Payload, proto3Small)
	if err != nil {
		fmt.Println("P3 small unmarshall err:", err)
		return
	}
	fmt.Println("proto3 small converted:", proto3Small)

	p3PayloadSmall, err := proto.Marshal(proto3Small)
	if err != nil {
		fmt.Println("P3 small marshall err:", err)
		return
	}
	fmt.Println("P3 small payload:", p3PayloadSmall)
	proto3BackLarge := &protobuf.LargerP3Define{}
	err = proto.Unmarshal(p3PayloadSmall, proto3BackLarge)
	if err != nil {
		fmt.Println("P3 small unmarshall err:", err)
		return
	}
	fmt.Println("proto3 back large converted:", proto3BackLarge)

	fmt.Println("Let s try it with from p3 to p2")
	proto2Small := &protobuf.SmallerP2Define{}
	err = proto.Unmarshal(p3Payload, proto2Small)
	if err != nil {
		fmt.Println("p2 small unmarshall err:", err)
		return
	}
	fmt.Println("proto3 large to p2 small:", proto2Small)

}

func TestP2() {
	fmt.Println("Testing proto 2")
	proto2Large := &protobuf.LargerP2Define{
		Provider:      proto.Uint32(1),
		CompetitionId: proto.Uint32(2),
		EventId:       proto.Uint32(3),
		MarketId:      proto.Uint32(4),
	}
	fmt.Println("proto2 large origin:", proto2Large)
	p2Payload, err := proto.Marshal(proto2Large)
	if err != nil {
		fmt.Println("p2 large marshall err:", err)
		return
	}
	fmt.Println("proto2 large binary:", p2Payload)
	proto2Small := &protobuf.SmallerP2Define{}
	err = proto.Unmarshal(p2Payload, proto2Small)
	if err != nil {
		fmt.Println("p2 small unmarshall err:", err)
		return
	}
	fmt.Println("proto2 small converted:", proto2Small)

	p2PayloadSmall, err := proto.Marshal(proto2Small)
	if err != nil {
		fmt.Println("p2 small marshall err:", err)
		return
	}
	fmt.Println("p2 small payload:", p2PayloadSmall)
	proto2BackLarge := &protobuf.LargerP2Define{}
	err = proto.Unmarshal(p2PayloadSmall, proto2BackLarge)
	if err != nil {
		fmt.Println("p2 small unmarshall err:", err)
		return
	}
	fmt.Println("proto2 back large converted:", proto2BackLarge)

	r := gin.Default()
	r.GET("/p2s", func(c *gin.Context) {
		c.JSON(200, proto2Small)
	})
	r.GET("/p2bl", func(c *gin.Context) {
		c.JSON(200, proto2BackLarge)
	})

	// based on the value in p2Payload and p2PayloadSmall, the application/octet-stream return will be the same
	marshaledData1, _ := proto.Marshal(proto2Small)
	r.GET("/p2s/binary", func(c *gin.Context) {
		c.Data(200, "application/octet-stream", marshaledData1)
	})

	marshaledData2, _ := proto.Marshal(proto2BackLarge)
	r.GET("/p2bl/binary", func(c *gin.Context) {
		c.Data(200, "application/octet-stream", marshaledData2)
	})

	recPayload, err := json.Marshal(proto2Small)
	if err != nil {
		fmt.Println("proto 2 large marshal json", err)
		return
	}

	// this only way I can think of is convert the the data to json inorder to strip
	jsonSmall := &protobuf.SmallerP2Define{}
	err = json.Unmarshal(recPayload, jsonSmall)
	if err != nil {
		fmt.Println("Unmarshal(recPayload, jsonSmall)", err)
		return
	}
	fmt.Println("jsonSmall", jsonSmall)
	fmt.Println("recPayload", recPayload)
	lastPayload, err := proto.Marshal(jsonSmall)
	if err != nil {
		fmt.Println("proto marshal", err)
		return
	}
	fmt.Println("the binary payload of p2 small after json convert:", lastPayload)
	r.Run() // listen and serve on 0.0.0.0:8080
}

func TestP2ToP3() {
	fmt.Println("Testing from p2 to p3")
	proto2Large := &protobuf.LargerP2ToP3Define{
		Provider:        proto.Uint32(1),
		CompetitionId:   proto.Uint32(2),
		EventId:         proto.Uint32(3),
		MarketId:        proto.Uint32(4),
		Outcome:         proto.String("5"),
		SpecialBetValue: proto.String("6"),
		Odds:            proto.Float64(7.1),
		IsLive:          proto.Bool(true),
		CurrencyPair:    proto.String("8"),
		StakeFactor:     proto.Float64(9.1),
	}

	fmt.Println("proto2 large origin:", proto2Large)
	p2Payload, err := proto.Marshal(proto2Large)
	if err != nil {
		fmt.Println("p2 large marshall err:", err)
		return
	}
	fmt.Println("proto2 large binary:", p2Payload)
	proto3Small := &protobuf.SmallerP3Define{}
	err = proto.Unmarshal(p2Payload, proto3Small)
	if err != nil {
		fmt.Println("P3 small unmarshall err:", err)
		return
	}
	fmt.Println("proto3 small converted:", proto3Small)

	p3PayloadSmall, err := proto.Marshal(proto3Small)
	if err != nil {
		fmt.Println("P3 small marshall err:", err)
		return
	}
	fmt.Println("P3 small payload:", p3PayloadSmall)
	proto3BackLarge := &protobuf.LargerP3Define{}
	err = proto.Unmarshal(p3PayloadSmall, proto3BackLarge)
	if err != nil {
		fmt.Println("P3 small unmarshall err:", err)
		return
	}
	fmt.Println("proto3 back large converted:", proto3BackLarge)
}
