package main

import (
	"fmt"
	"time"

	"github.com/tinylib/msgp/msgp"
)

type Sample struct {
	StartTime time.Time
}

var buf [64]byte

func main() {
	e := Sample{}
	b := buf[:0]
	b = msgp.AppendMapHeader(b, 1)
	b = msgp.AppendString(b, "st")
	b = msgp.AppendTime(b, e.StartTime)

	var o Sample
	sz, b, err := msgp.ReadMapHeaderBytes(b)
	if err != nil {
		panic(err)
	}
	if sz != 1 {
		panic("bad sz")
	}

	sb, b, err := msgp.ReadStringZC(b)
	switch string(sb) {
	case "st":
		var t time.Time
		t, _, err = msgp.ReadTimeBytes(b)
		if err != nil {
			panic("unexpected failure:" + err.Error())
		}
		o.StartTime = t
	default:
		panic("unexpected field:" + string(sb))
	}

	fmt.Println(e)
	fmt.Println(o)
}
