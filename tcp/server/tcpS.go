package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
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
    log.Printf("Server is listening on PORT %s\n", PORT)

    // server will only be able to connect to one TCP client, 
    // the first to have a successful connection.
    c, err := l.Accept()
    if err != nil {
        fmt.Println(err)
        return
    }

    log.Printf("Connected to: %s\n", c.RemoteAddr())

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

        userMsg := strings.ToLower(strings.TrimSpace(string(netData)))

        switch(userMsg) {
            case "hi": {
                c.Write([]byte("yo hows it going\n"))
                continue
            }
            case "hello world": {
                c.Write([]byte("touch grass\n"))
                continue
            }
            case "what": {
                c.Write([]byte("yes.\n"))
                continue
            }
            case "STOP": {
                c.Write([]byte("bye! ending connection.\n"))
                return
            }
        }

        if len(strings.TrimSpace(string(netData))) > 5 {
            t := time.Now()
            myTime := "idk what you said but uh it's currently: " + t.Format(time.RFC3339) + "\n"
            c.Write([]byte(myTime))
        } else {
            c.Write([]byte("can you say that again but with more verbosity?\n"))
        }
    }
}
