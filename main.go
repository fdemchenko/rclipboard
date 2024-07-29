package main

import (
	"bytes"
	"flag"
	"fmt"
	"net"
	"os"

	"golang.design/x/clipboard"
)

type Config struct {
	ListenAddr string
	RemoteAddr string
}

const PullClipboardMessage = "PULL_CLIPBOARD\n"

func main() {
	var cfg Config
	flag.StringVar(&cfg.ListenAddr, "listen-addr", "", "TCP listen address")
	flag.StringVar(&cfg.RemoteAddr, "remote-addr", "", "TCP remote address of another clipboard")
	flag.Parse()

	pullRequests, err := pullClipboardEvents()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		os.Exit(1)
	}

	errors := startPullClipboardListener(cfg.ListenAddr)
	go func() {
		for err := range errors {
			fmt.Fprintf(os.Stderr, "error: %s\n", err.Error())
		}
	}()

	for {
		<-pullRequests

		conn, err := net.Dial("tcp", cfg.RemoteAddr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dial error: %s\n", err.Error())
			continue
		}

		_, err = conn.Write([]byte(PullClipboardMessage))
		if err != nil {
			fmt.Fprintf(os.Stderr, "writing error: %s\n", err.Error())
			conn.Close()
			continue
		}

		buffer := new(bytes.Buffer)
		_, err = buffer.ReadFrom(conn)
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading error: %s\n", err.Error())
			conn.Close()
			continue
		}
		clipboard.Write(clipboard.FmtText, buffer.Bytes())
		conn.Close()
	}
}
