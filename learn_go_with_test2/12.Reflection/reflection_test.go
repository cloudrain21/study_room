package reflection

import (
    "reflect"
    "testing"
)

type Person struct {
    Name string
    Profile Profile
}

type Profile struct {
    Age int
    City string
}

func TestWalk(t *testing.T) {
    t.Run("test1", func(t *testing.T) {
        expected := "Chris"
        var got []string

        x := struct {
            Name string
        }{expected}

        walk(x, func(input string) {
            got = append(got, input)
        })

        if got[0] != expected {
            t.Errorf("got : %s, want : %s", got, expected)
        }
    })

    cases := []struct {
        Name string
        Input interface{}
        ExpectedCalls []string
    } {
        {
            "Struct with one string field",
            struct {
                Name string
            } {"Chris"},
            []string{"Chris"},
        },
        {
            "Struct with two string field",
            struct {
                Name string
                City string
            }{"Chris", "London"},
            []string{"Chris", "London"},
        },
        {
            "Struct with non string field",
            struct {
                Name string
                Age int
            }{"Chris", 33},
            []string{"Chris"},
        },
        {
            "Nested fields",
            struct {
                Name string
                Profile struct {
                    Age int
                    City string
                }
            } { "Chris",
                struct {
                    Age int
                    City string
                }{33, "London"},
            },
            []string{"Chris", "London"},
        },
        {
            "Nested fields - simple struct",
            Person {
                "Chris",
                Profile { 33, "London" },
            },
            []string{"Chris", "London"},
        },
    }

    for _, test := range cases {
        t.Run(test.Name, func(t *testing.T) {
            var got []string
            walk(test.Input, func(input string) {
                got = append(got, input)
            })

            if !reflect.DeepEqual(got, test.ExpectedCalls) {
                t.Errorf("got %v expected %v", got, test.ExpectedCalls)
            }
        })
    }
}

func TestWalk2(t *testing.T) {
    t.Run("pointer", func(t *testing.T) {
        cases := []struct {
            Name string
            Input interface{}
            ExpectedCalls []string
        } {
            {
                "string",
                "Kimpo",
                []string{"Kimpo"},
            },
            {
                "structure",
                Profile {48, "Kimpo"},
                []string{"Kimpo"},
            },
            {
                "Pointers to things",
                &Person {
                    "Chris",
                    Profile{33, "London"},
                },
                []string{"Chris", "London"},
            },
            {
                "Slices",
                []string{"Chris", "London"},
                []string{"Chris", "London"},
            },
            {
                "Slice of struct",
                []Profile{
                    {33,"Seoul"},
                    {34,"Kimpo"},
                },
                []string{"Seoul", "Kimpo"},
            },
            {
                "Arrays",
                [3]Profile{
                    {33,"Seoul"},
                    {34,"Kimpo"},
                    {35,"Pankyo"},
                },
                []string{"Seoul", "Kimpo", "Pankyo"},
            },
        }

        for _, c := range cases {
            got := []string{}

            walk(c.Input, func(input string) {
                got = append(got, input)
            })

            if ! reflect.DeepEqual(got, c.ExpectedCalls) {
                t.Errorf("got %v want %v", got, c.ExpectedCalls)
            }
        }
    })

    t.Run("map test", func(t *testing.T) {
        aMap := map[string]string {
            "aaa":"AAAA",
            "bbb":"BBBB",
        }

        got := []string{}

        walk(aMap, func(input string) {
            got = append(got, input)
        })

        assertContains(t, got, "AAAA")
        assertContains(t, got, "BBBB")
    })

    t.Run("channel test", func(t *testing.T) {
        aChannel := make(chan Profile)

        go func() {
            aChannel <- Profile{33, "Chris"}
            aChannel <- Profile{34, "Cloudrain"}
            close(aChannel)
        }()

        got := []string{}
        want := []string{"Chris", "Cloudrain"}

        walk(aChannel, func(input string) {
            got = append(got, input)
        })

        if ! reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    })

    t.Run("function test", func(t *testing.T) {
        aFunc := func() []Profile {
            return []Profile{
                Profile{33, "Chris"},
                Profile{34, "Cloudrain"},
            }
        }

        got := []string{}
        want := []string{"Chris", "Cloudrain"}

        walk(aFunc, func(input string) {
            got = append(got, input)
        })

        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    })
}

func assertContains(t *testing.T, haystack []string, needle string)  {
    t.Helper()
    contains := false
    for _, x := range haystack {
        if x == needle {
            contains = true
        }
    }
    if !contains {
        t.Errorf("expected %+v to contain %q but it didn't", haystack, needle)
    }
}
