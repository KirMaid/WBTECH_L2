package telnet

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
	//"github.com/reiver/go-telnet"
)

const defaultTimeout = 10 * time.Second

type Options struct {
	Host    string
	Port    int
	Timeout time.Duration
}

func main() {
	opts := parseOptions()

	ctx, cancel := context.WithTimeout(context.Background(), opts.Timeout)
	defer cancel()

	go func() {
		select {
		case <-ctx.Done():
			if ctx.Err() == context.DeadlineExceeded {
				fmt.Println("Connection timed out.")
				os.Exit(1)
			}
		}
	}()

	address := fmt.Sprintf("%s:%d", opts.Host, opts.Port)
	conn, err := net.DialTimeout("tcp", address, opts.Timeout)
	if err != nil {
		fmt.Printf("Failed to connect to %s: %v\n", address, err)
		os.Exit(1)
	}
	defer conn.Close()

	telnetConn := telnet.NewConn(conn)
	defer telnetConn.Close()

	go handleInput(telnetConn)
	handleOutput(telnetConn)
}

func parseOptions() Options {
	host := ""
	port := 23 // Default Telnet port
	timeout := defaultTimeout

	args := os.Args[1:]
	for i, arg := range args {
		switch arg {
		case "--timeout":
			if i+1 < len(args) {
				duration, err := time.ParseDuration(args[i+1])
				if err == nil {
					timeout = duration
				}
			}
		default:
			if host == "" {
				host = arg
			} else if port == 23 {
				p, err := strconv.Atoi(arg)
				if err == nil {
					port = p
				}
			}
		}
	}

	return Options{Host: host, Port: port, Timeout: timeout}
}

func handleInput(conn *telnet.Conn) {
	reader := bufio.NewReader(os.Stdin)
	buffer := make([]byte, 1024)

	for {
		data, err := reader.ReadBytes('\x04') // Ctrl+D
		if err != nil {
			break
		}
		_, _ = conn.Write(data)
	}
}

func handleOutput(conn *telnet.Conn) {
	buffer := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			break
		}
		os.Stdout.Write(buffer[:n])
	}
}
