package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

func handlerError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err.Error())
		os.Exit(1)
	}
}

const DefaultMessage = "PING"

func main() {
	var daemonAddress string
	flag.StringVar(&daemonAddress, "address", "", "Remote address to write")
	flag.Parse()

	message := flag.Arg(0)
	if message == "" {
		message = DefaultMessage
	}

	conn, err := net.Dial("tcp", daemonAddress)
	handlerError(err)

	_, err = conn.Write([]byte(message))
	handlerError(err)
	conn.Close()
}
