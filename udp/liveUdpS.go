package main

import (
    "fmt"
    "os"
    "net"
)

func main() {
    args := os.Args
    if len(args) == 1 {
        fmt.Println("Please provide port.")
        return
    }

    port := args[1]
    addr := ":" + port

    laddr, err := net.ResolveUDPAddr("udp4", addr)
    if err != nil {
        panic(err)
    }

    conn, err := net.ListenUDP("udp4", laddr)
    if err != nil {
        panic(err)
    }
    defer conn.Close()
    
    buffer := make([]byte, 1024)
    for {
        _, addr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            panic(err)
        }
        data := string(buffer)
        
        // fmt.Println(n)
        // fmt.Println(addr) // who to send a response to, if you want.
        // fmt.Println(data)
        fmt.Printf("%s says: %s\n", addr, data)
    }
}
