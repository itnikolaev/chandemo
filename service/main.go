package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/itnikolaev/chandemo/rps"
)

func main() {
	runtime.GOMAXPROCS(3)
	go func() {
		for {
			time.Sleep(time.Millisecond * time.Duration(rand.Int31n(5)))
			rps.AddMetric(time.Now().Unix())
		}
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, os.Kill, syscall.SIGTERM)
	for {
		select {
		case killSignal := <-interrupt:
			fmt.Println("Got signal:", killSignal)
			if killSignal == os.Interrupt {
				fmt.Println("process was interruped by system signal")
			}
			return
		}
	}

}
