package main

import "fmt"

func HelloStr(prefix, suffix string) string {
    if suffix == "" {
        suffix = "World"
    }
    return prefix + ", " + suffix
}

func Hello(s, lang string) string {
    switch lang {
        case "Spanish": 
            return HelloStr("Hola", s)
        default:
            return HelloStr("Hello", s)
    }
}

func main() {
    fmt.Println(Hello("world", "English"))
}
