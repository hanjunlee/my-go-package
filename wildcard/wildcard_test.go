package wildcard

import (
    "testing"
    "log"
)

func TestMatch(t *testing.T) {
    cases := []struct {
        pattern string
        str string
        ok bool
    } {
        { "", "", true },
        { "", "a", false },
        { "baaa?ab", "baaabab", true },
        { "ba*a?", "baaabab", true },
        { "a*ab", "baaabab", false },
        { "a", "aa", false },
        { "*****ba*****ab", "baaabab", true },
        { "?*", "ab", true },
    }

    for _, c := range cases {
        log.Println(c)
        if match := Match(c.pattern, c.str);
        (match && c.ok) || (!match && !c.ok) {
            continue
        }

        t.Error(c)
    }
}

