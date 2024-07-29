package main

import (
    "fmt"
    "os"
    "syscall"
    "unsafe"
)

func main() {
    fd := int(os.Stdin.Fd())

    // save current terminal state
    var oldState syscall.Termios
    _, _, err := syscall.Syscall6(syscall.SYS_IOCTL,
                                     uintptr(fd),
                                     uintptr(syscall.TCGETS),
                                     uintptr(unsafe.Pointer(&oldState)),
                                     0,
                                     0,
                                     0)
    if err != 0 {
        fmt.Println("Error getting terminal state:", err)
        return
    }

    // set terminal to raw mode
    newState := oldState
    newState.Lflag &^= syscall.ICANON | syscall.ECHO
    _, _, err = syscall.Syscall6(syscall.SYS_IOCTL,
                                     uintptr(fd),
                                     uintptr(syscall.TCSETS),
                                     uintptr(unsafe.Pointer(&newState)),
                                     0,
                                     0,
                                     0)
    if err != 0 {
        fmt.Println("Error setting terminal to raw mode:", err)
        return
    }

    // restore terminal to state when done
    defer func() {
        _, _, err = syscall.Syscall6(syscall.SYS_IOCTL,
                                        uintptr(fd),
                                        uintptr(syscall.TCSETS),
                                        uintptr(unsafe.Pointer(&oldState)),
                                        0,
                                        0,
                                        0)
        if err != 0 {
            fmt.Println("Error restoring terminal state:", err)
        }
    }()

    fmt.Println("Start typing. (Ctrl+C to exit):")
    fmt.Print("> ")

    // read char one by one
    var b []byte = make([]byte, 1)
    for {
        if _, err := os.Stdin.Read(b); err != nil {
            fmt.Println("Error reading from stdin:", err)
            break
        }
        // send b[0] char to udp server to stream input live to it.
        fmt.Printf("You typed: %c\n", b[0])
    }

}
