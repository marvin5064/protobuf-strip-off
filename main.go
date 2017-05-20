package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/marvin5064/protobuf-strip-off/protobuf"
)

func main() {
	fmt.Println("Testing if protobuf can be strip off by define automatically")
	TestP3()
	TestP2()
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
	r.Run() // listen and serve on 0.0.0.0:8080
}
