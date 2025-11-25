package main

import (
    "log"
    "net"
    "net/rpc"
)

// Args holds arguments passed to RPC methods
type Args struct {
    A, B int
}

// Arith provides arithmetic methods
type Arith struct {}

// Add sums two numbers
func (a *Arith) Add(args Args, reply *int) error {
    *reply = args.A + args.B
    return nil
}

func main() {
    rpc.Register(new(Arith))
    l, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatalf("Listen error: %v", err)
    }
    log.Println("RPC server listening on :1234")
    for {
        conn, err := l.Accept()
        if err != nil {
            log.Printf("Accept error: %v", err)
            continue
        }
        go rpc.ServeConn(conn)
    }
}