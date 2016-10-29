package main

import (
    "log"
)

type Money struct {
    amount uint
    currency string
}

func main() {
    money := &Money{120, "yen"}
    log.Printf("%#V", money)
}

