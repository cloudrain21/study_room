package mystrings

import "strings"

func MyJoin(s1, s2 string) string {
    return s1 + s2
}

func MyContains(s, f string) bool {
    return strings.Contains(s, f)
}

func MyIndex(s, f string) int {
    return strings.Index(s,f)
}
