package stp

import (
	"os"
	"os/signal"
)

func Hold() {
	s := make(chan os.Signal, 10)
	signal.Notify(s, os.Interrupt)
	<-s
	signal.Stop(s)
	close(s)
}
