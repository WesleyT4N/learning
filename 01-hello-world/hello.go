package main

import "fmt"

const engHelloPrefix = "Hello, "

func Hello(name string) string {
    if name == "" {
        name = "World"
    }
    return engHelloPrefix + name
}

func main()  {
    fmt.Println(Hello("world"))
}
