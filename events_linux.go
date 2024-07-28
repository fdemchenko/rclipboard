package main

import (
	"os"
	"os/signal"
	"syscall"
)

func pullClipboardEvents() (chan struct{}, error) {
	events := make(chan struct{})

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGUSR1)

	go func() {
		for {
			<-signals
			events <- struct{}{}
		}
	}()

	return events, nil
}
