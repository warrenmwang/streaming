package main
import (
    "bufio"
    "fmt"
    "net"
    "os"
    "strconv"
    "strings"
)

func handleConnection(c net.Conn, count int) {
    defer c.Close()
    fmt.Print(".")
    for {
        netData, err := bufio.NewReader(c).ReadString('\n')
        if err != nil {
            fmt.Println(err)
            return
        }

        tmp := strings.TrimSpace(string(netData))
        if tmp == "STOP" {
            break
        }
        fmt.Println(tmp)
        counter := strconv.Itoa(count) + "\n"
        _, err = c.Write([]byte(string(counter)))
        if err != nil {
            fmt.Println(err)
            return
        }
    }
}

func main() {
    var count = 0

    args := os.Args
    if len(args) == 1 {
        fmt.Println("Please provide a port number!")
        return
    }

    PORT := ":" + args[1]
    l, err := net.Listen("tcp4", PORT)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer l.Close()

    for {
        c, err := l.Accept()
        if err != nil {
            fmt.Println(err)
            return
        }

        go handleConnection(c, count)
        count++
    }
}
