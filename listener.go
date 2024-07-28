package main

import (
	"bytes"
	"net"

	"golang.design/x/clipboard"
)

func startPullClipboardListener(address string) chan error {
	errorsChan := make(chan error)
	listener, err := net.Listen("tcp", address)

	go func() {
		if err != nil {
			errorsChan <- err
		}
		for {
			conn, err := listener.Accept()
			if err != nil {
				errorsChan <- err
			}

			buffer := new(bytes.Buffer)
			_, err = buffer.ReadFrom(conn)
			if err != nil {
				errorsChan <- err
			}

			if buffer.String() != PullClipboardMessage {
				conn.Close()
				continue
			}

			clipboardBytes := clipboard.Read(clipboard.FmtText)
			if clipboardBytes == nil {
				conn.Close()
				continue
			}

			_, err = conn.Write(clipboardBytes)
			if err != nil {
				errorsChan <- err
			}
		}
	}()

	return errorsChan
}
