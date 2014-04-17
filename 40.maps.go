package main

import (
    "code.google.com/p/go-tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    maps := make(map[string]int)
    fields := strings.Fields(s)

    for _, value := range fields {
    	maps[value]++
    }
    
    return maps
}

func main() {
    wc.Test(WordCount)
}
