package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	
	go func() {
	  for {
	    c1 <- "500ms"
	    time.Sleep(time.Millisecond * 500)
	  }
	}()
	go func() {
	  for {
	    c1 <- "2s"
	    time.Sleep(time.Second * 2)
	  }
	}()
	
	for i := 1; i < 20; i++{
	  select {
	    case msg1 := <- c1:
	      fmt.Println(msg1)
	    case msg2 := <- c2:
	      fmt.Println(msg2)
	  }
	}
}
