package main

import (
	"fmt"
	"github.com/niqinge/test-simple/common"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	common.InitLogger("testsimple", ".")

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		s := <-c
		fmt.Println(fmt.Sprintf("server get signal %s", s.String()))
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:

			fmt.Println("exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
