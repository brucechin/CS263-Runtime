package socket

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net"
	"os"
	"sync"
	"testing"
)

const (
	aHugeFile = "../../rust/compression/data/alice29.txt"
)

func Server(port string, serverReady *Barrier) {
	fmt.Println("start server...")

	// listen on port
	ln, _ := net.Listen("tcp", port)

	// a barrier when both server is ready
	serverReady.Wait()

	// accept connection
	conn, _ := ln.Accept()

	// loop. receive message forever
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Received: ", string(message))
	}
}

func Client(b *testing.B, port string, numBytes int, serverReady *Barrier, connectReady *sync.WaitGroup, workStart *Barrier, workStop *Barrier) {

	// load a large amount of text to send
	buf, err := os.ReadFile(aHugeFile)
	if err != nil {
		b.Fatal(err)
	}

	buf1 := new(bytes.Buffer)
	w := bufio.NewWriter(buf1)
	for i := 0; i < numBytes; i += len(buf) {
		if len(buf) > numBytes-i {
			buf = buf[:numBytes-i]
		}
		io.Copy(w, bytes.NewReader(buf))
	}

	// a barrier when both sender and recver are ready
	serverReady.Wait()

	// connect to server
	conn, _ := net.Dial("tcp", "127.0.0.1"+port)

	// signal the main thread that it can start timing
	connectReady.Done()

	fmt.Println("client sending...")
	for {
		workStart.Wait()

		// send numBytes text
		reader := bufio.NewReader(bytes.NewReader(buf1.Bytes()))
		text, _ := reader.ReadString('\n')

		fmt.Fprintf(conn, text+"\n")

		workStop.Wait()
	}
	fmt.Println("client exit...")
}

func BenchmarkTcpSencRecv32B(b *testing.B) { benchTcpSendRecv(b, 32) }

func benchTcpSendRecv(b *testing.B, numBytes int) {

	serverReady := NewBarrier(2) // only between server and client

	var connectReady sync.WaitGroup
	connectReady.Add(1)

	workStart := NewBarrier(2) // only between main and client
	workStop := NewBarrier(2)  // only between main and client

	go Server(":8000", serverReady)

	go Client(b, ":8000", numBytes, serverReady, &connectReady, workStart, workStop)

	connectReady.Wait()

	// start timing
	b.StartTimer()

	for i := 0; i < b.N; i++ {

		// signal sender to begin
		workStart.Wait()

		// wait for sender done
		workStop.Wait()

	}

}

// structures to help sync workers
type Barrier struct {
	total int
	count int
	mutex *sync.Mutex
	cond  *sync.Cond
}

func NewBarrier(size int) *Barrier {
	lockToUse := &sync.Mutex{}
	condToUse := sync.NewCond(lockToUse)
	return &Barrier{size, size, lockToUse, condToUse}
}

func (b *Barrier) Wait() {
	b.mutex.Lock()
	b.count -= 1
	if b.count == 0 {
		b.count = b.total
		b.cond.Broadcast()
	} else {
		b.cond.Wait()
	}
	b.mutex.Unlock()
}
