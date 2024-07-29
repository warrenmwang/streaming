package main

import (
    "fmt"
    "os"
    "strconv"
    "strings"
    "net"
    "regexp"
)

func main() {
    args := os.Args
    if len(args) == 1 {
        fmt.Println("Please provide a port number")
        return
    }
    PORT := ":" + args[1]

    s, err := net.ResolveUDPAddr("udp4", PORT)
    if err != nil {
        fmt.Println(err)
        return
    }

    conn, err := net.ListenUDP("udp4", s)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()
    buffer := make([]byte, 1024)

    for {
        n, addr, err := conn.ReadFromUDP(buffer)
        if err != nil {
            fmt.Println(err)
            return
        }
        str := string(buffer[0:n])
        fmt.Print("-> ", str[0:n-1])


        if strings.TrimSpace(str) == "STOP" {
            fmt.Println("Exiting UDP server...")
            return
        }

        var res string
        // Check format of the given string
        r, err := regexp.Compile("[0-9],[0-9]")
        if err != nil {
            fmt.Println(err)
            return
        }

        if ! r.MatchString(strings.TrimSpace(str)) {
            fmt.Println("Unexpected input not of format <0-9>,<0-9>")
            res = "invalid input"
        } else {
            // Parse the given data and add the two numbers together
            tmp := strings.Split(strings.TrimSpace(str), ",")

            a, err := strconv.Atoi(tmp[0])
            if err != nil {
                fmt.Println(err)
                return
            }
            b, err := strconv.Atoi(tmp[1])
            if err != nil {
                fmt.Println(err)
                return
            }
            res = fmt.Sprintf("%d", a + b)
        }

        // Send result
        data := []byte(res)
        fmt.Printf("data: %s\n", string(data))
        _, err = conn.WriteToUDP(data, addr)
        if err != nil {
            fmt.Println(err)
            return
        }
    }
}
