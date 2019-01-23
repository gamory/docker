package main

import (
    "fmt"
    "os"
)
var debug_flag bool = false

func main() {
    if os.Args[1] == "--debug" {
        debug_flag = true
        fmt.printf("> Debugging enabled <")
        cmd_args := os.Args[2:]
    } else {
        cmd_args := os.Args[1:]
    }
    return
}
