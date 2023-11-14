package solumid

import (
	"fmt"
	"leetgo/tool"
	"time"
)

func Test() {
	timer := tool.Timer{}
	timer.Start()
	time.Sleep(4 * time.Second)
	timer.Stop()
	fmt.Println("Duration", timer.Duration())
}
