package iteration

func Repeat(a string, count int) string {
    r := ""
    for i:=0; i<count; i++ {
        r += a
    }
    return r
}
