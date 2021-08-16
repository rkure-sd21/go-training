package main

import (
	"fmt"
	"syncutil/syncutil"
)

func main() {
	c := &syncutil.Counter{
		Name: "Access",
	}
	fmt.Println(c.Increment)
	fmt.Println(c.View())
}
