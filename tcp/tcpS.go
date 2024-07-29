package main

import (
    "fmt"
    "net"
    "os"
    "strings"
    "bufio"
    "time"
)

func main() {
    args := os.Args
    if len(args) == 1 {
        fmt.Println("Please provide port")
        return
    }

    PORT := ":" + args[1]
    l, err := net.Listen("tcp", PORT)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer l.Close()

    // server will only be able to connect to one TCP client, 
    // the first to have a successful connection.
    c, err := l.Accept()
    if err != nil {
        fmt.Println(err)
        return
    }

    for {
        netData, err := bufio.NewReader(c).ReadString('\n')
        if err != nil {
            fmt.Println(err)
            return
        }
        if strings.TrimSpace(string(netData)) == "STOP" {
            fmt.Println("Exiting TCP server...")
            return
        }

        fmt.Print("-> ", string(netData))

        if len(strings.TrimSpace(string(netData))) > 5 {
            t := time.Now()
            myTime := t.Format(time.RFC3339) + "\n"
            c.Write([]byte(myTime))
        } else {
            c.Write([]byte("send longer\n"))
        }
    }
}
