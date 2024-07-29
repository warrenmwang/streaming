package main

import (
    "os"
    "fmt"
    "net"
    "syscall"
    "unsafe"
)

func main() {
    args := os.Args
    if len(args) == 1 {
        fmt.Println("Please provide host:port")
        return
    }

    fd := int(os.Stdin.Fd())
    // save current terminal state
    var oldState syscall.Termios
    _, _, termErr := syscall.Syscall6(syscall.SYS_IOCTL,
                                     uintptr(fd),
                                     uintptr(syscall.TCGETS),
                                     uintptr(unsafe.Pointer(&oldState)),
                                     0,
                                     0,
                                     0)
    if termErr != 0 {
        fmt.Println("Error getting terminal state:", termErr)
        return
    }

    // set terminal to raw mode
    newState := oldState
    newState.Lflag &^= syscall.ICANON | syscall.ECHO
    _, _, termErr = syscall.Syscall6(syscall.SYS_IOCTL,
                                     uintptr(fd),
                                     uintptr(syscall.TCSETS),
                                     uintptr(unsafe.Pointer(&newState)),
                                     0,
                                     0,
                                     0)
    if termErr != 0 {
        fmt.Println("Error setting terminal to raw mode:", termErr)
        return
    }

    // restore terminal to state when done
    defer func() {
        _, _, termErr = syscall.Syscall6(syscall.SYS_IOCTL,
                                        uintptr(fd),
                                        uintptr(syscall.TCSETS),
                                        uintptr(unsafe.Pointer(&oldState)),
                                        0,
                                        0,
                                        0)
        if termErr != 0 {
            fmt.Println("Error restoring terminal state:", termErr)
        }
    }()


    // Get the UDP connection ready
    addr := args[1]
    raddr, err := net.ResolveUDPAddr("udp4", addr)
    if err != nil {
        panic(err)
    }

    conn, err := net.DialUDP("udp4", nil, raddr)
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    fmt.Println("Live streaming your input now.")
    fmt.Print("> ")
    var b []byte = make([]byte, 1)
    for {
        if _, err := os.Stdin.Read(b); err != nil {
            fmt.Println("Error reading from stdin:", err)
            break
        }
        // send b[0] char to udp server to stream input live to it.
        inputStr := fmt.Sprintf("%c", b[0])
        data := []byte(inputStr)
        _, err := conn.Write(data)
        if err != nil {
            fmt.Println(err)
            return 
        }
        // print typed char to stdout to give feedback
        fmt.Print(inputStr)
        if inputStr == "\n" {
            fmt.Printf("> ")
        }
    }
}
