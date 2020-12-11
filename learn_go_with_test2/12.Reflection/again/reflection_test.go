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
        } {expected}

        walk(x, func(input string) {
            got = append(got, input)
        })

        if got[0] != expected {
            t.Errorf("got %s want %s", got[0], expected)
        }
    })

    cases := []TestCase {
        {
            "Struct with one string field",
            struct {
                Name string
            }{"Chris"},
            []string{"Chris"},
        },
        {
            "Struct with two string fields",
            struct {
                Name string
                City string
            }{"Chris","London"},
            []string{"Chris","London"},
        },
        {
            "Struct with one string field",
            struct {
                Name string
                Age  int
            } {"Chris", 33},
            []string{"Chris"},
        },
        {
            "Nested fields",
            Person {"Chris", Profile{33, "London" }},
            []string{"Chris","London"},
        },
        {
            "Pointers to things",
            &Person{"Chris",
                Profile{33, "London"},
            },
            []string{"Chris", "London"},
        },
        {
            "Slices",
            []Profile {
                {33, "London"},
                {33, "Seoul"},
            },
            []string{"London", "Seoul"},
        },
        {
            "Slices",
            [2]Profile {
                {33, "London"},
                {33, "Seoul"},
            },
            []string{"London", "Seoul"},
        },
        {
            "Map",
            map[string]string {
                "Foo": "Bar",
                "Baz": "Boz",
            },
            []string{"Bar", "Boz"},
        },
    }

    for _, c := range cases {
        t.Run(c.TestName, func(t *testing.T) {
            var got []string
            walk(c.Input, func(input string) {
                got = append(got, input)
            })

            // error can occur because of ordering
            if !reflect.DeepEqual(got, c.ExpectedCalls) {
                t.Errorf("testname(%s) got %v want %v", c.TestName, got, c.ExpectedCalls)
            }
        })
    }

    t.Run("channel", func(t *testing.T) {
        aChannel := make(chan Profile)

        go func() {
            aChannel <- Profile{33, "Chris"}
            aChannel <- Profile{34, "Cloudrain"}
            close(aChannel)
        }()

        want := []string{"Chris", "Cloudrain"}
        var got []string
        walk(aChannel, func(input string) {
            got = append(got, input)
        })

        if !reflect.DeepEqual(got,want) {
            t.Errorf("got %v want %v", got, want)
        }
    })

    t.Run("function", func(t *testing.T) {
        f := func() (Profile, Profile) {
            return Profile{33, "Chris"}, Profile{34, "Cloudrain"}
        }

        var got []string
        want := []string{"Chris", "Cloudrain"}

        walk(f, func(input string) {
            got = append(got, input)
        })

        if !reflect.DeepEqual(got,want) {
            t.Errorf("got %v want %v", got, want)
        }
    })
}

type TestCase struct {
    TestName string
    Input interface{}
    ExpectedCalls []string
}
