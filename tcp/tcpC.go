package main

import (
    "fmt"
    "net"
    "os"
    "strings"
    "bufio"
)

func main() {
    args := os.Args
    if len(args) == 1 {
        fmt.Println("Please provide host:port")
        return 
    }

    CONN := args[1]
    c, err := net.Dial("tcp", CONN)
    if err != nil {
        fmt.Println(err)
        return
    }

    for {
        reader := bufio.NewReader(os.Stdin)
        fmt.Print(">> ")
        text, _ := reader.ReadString('\n')
        fmt.Fprintf(c, text+"\n")

        message, _ := bufio.NewReader(c).ReadString('\n')
        fmt.Print("->: " + message)
        if strings.TrimSpace(string(text)) == "STOP" {
            fmt.Println("TCP client exiting...")
            return
        }
    }

}
