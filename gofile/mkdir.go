package main

import (
    "os"
    "syscall"
)

func main() {
    os.Mkdir("/tmp/go-dir1", 0777) // create 755 dir
    
    defaultUmask := syscall.Umask(0)
    os.Mkdir("/tmp/go-dir2", 0777) // create 755 dir
    syscall.Umask(defaultUmask)

    directory := "/tmp/go-dir3"
    os.Mkdir(directory, os.ModePerm)
    os.Chmod(directory, 0777)
}


