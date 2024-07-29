package main

import (
    "fmt"
    "bufio"
    "net"
    "os"
    "strings"
)

func main() {
    args := os.Args
    if len(args) == 1 {
        fmt.Println("Please provide a host:port string")
        return
    }
    CONN := args[1]

    s, err := net.ResolveUDPAddr("udp4", CONN)
    if err != nil {
        fmt.Println(err)
        return
    }
    c, err := net.DialUDP("udp4", nil, s)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Printf("The UDP Server is %s\n", c.RemoteAddr().String())
    defer c.Close()

    for {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print(">> ")
        text, _ := reader.ReadString('\n')
        data := []byte(text + "\n")
        _, err := c.Write(data)
        if err != nil {
            fmt.Println(err)
            return 
        }

        // Check if stop message
        if strings.TrimSpace(string(data)) == "STOP" {
            fmt.Println("Exiting UDP Client!")
            return
        }

        buffer := make([]byte, 1024)
        n, _, err := c.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println(err)
            return
        }
        fmt.Printf("Reply: %s\n", string(buffer[0:n]))
    }
}
