package cmd

import("fmt")

func check(e error) {
    if e != nil {
        fmt.Println(e)
        panic(e)
        
    }
}