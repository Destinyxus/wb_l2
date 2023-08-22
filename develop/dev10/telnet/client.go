package telnet

import (
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type Telnet struct {
	addr    string
	timeout time.Duration
}

func NewTelnet(addr string, timeout time.Duration) *Telnet {
	return &Telnet{
		addr:    addr,
		timeout: timeout,
	}
}

func (t *Telnet) Run() error {
	conn, err := net.DialTimeout("tcp", t.addr, t.timeout*time.Second)
	if err != nil {
		time.Sleep(t.timeout)
		fmt.Println("unreached server")
		return err
	}

	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	fmt.Printf("connected to %s\n", t.addr)

	chSignal := make(chan os.Signal)
	signal.Notify(chSignal, syscall.SIGINT, syscall.SIGTERM)
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg.Add(2)

	go func() {
		select {
		case <-chSignal:
			fmt.Println("shutdown signal..\n,press enter to finish the work")
			conn.Close()
			cancel()
			return
		}
	}()

	go func() {
		defer wg.Done()
		for {
			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				if err != io.EOF {
					return
				} else {
					fmt.Fprintf(os.Stderr, "error reading from server: %v\n", err)
					fmt.Println("press ENTER to terminate the work")
					cancel()
					return
				}
			}
			fmt.Print(string(buf[:n]))
		}
	}()

	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			default:
				buf := make([]byte, 1024)
				n, err := os.Stdin.Read(buf)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error reading from stdin: %v\n", err)
					return
				}
				_, err = conn.Write(buf[:n])
				if err != nil {
					fmt.Fprintf(os.Stderr, "error writing to server: %v\n", err)
					return
				}
			}

		}
	}(ctx)

	wg.Wait()
	return nil
}
