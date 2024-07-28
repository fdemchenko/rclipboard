package main

import (
	"io"
	"log"
	"net"
)

const Message = "CLIPBOARD_PULL_START"

func pullClipboardEvents() (chan struct{}, error) {
	events := make(chan struct{})
	listener, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				log.Printf("error: %s", err.Error())
				continue
			}

			data, err := io.ReadAll(conn)
			if err != nil {
				log.Printf("error: %s", err.Error())
				conn.Close()
				continue
			}

			if string(data) == Message {
				events <- struct{}{}
			}

			conn.Close()
		}
	}()

	return events, nil
}
