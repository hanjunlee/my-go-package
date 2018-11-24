/*
Match string with pattern

pattern include in only wildcard(*) and question mark(?)
*/

package wildcard

import (
)

var pattern, str []rune

// 
func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

// Match string with pattern
func Match(p, s string) bool {
    // set global variables
    pattern, str = StringToRuneSlice(p), StringToRuneSlice(s)

    return match(0, 0)
}

func match(idxP, idxS int) bool {
    // basis
    //
    // match
    if idxP == len(pattern) && idxS == len(str) {
        return true
    }
    // pattern reached at the end
    if idxP == len(pattern) && idxS != len(str) {
        return false
    }
    // str reached at the end
    // validate all charactors are wildcard
    if idxS == len(str) {
        for _, c := range pattern[idxP:] {
            if c != '*' {
                return false
            }
        }
        return true
    }

    // algorithm
    for idxP < len(pattern) && idxS < len(str) {
        if pattern[idxP] == '*' {
            break
        }
        if pattern[idxP] == '?' {
            idxP++; idxS++
            continue
        }
        if pattern[idxP] != str[idxS] {
            return false
        }

        idxP++; idxS++
    }

    // Reached
    // evaluated at basis
    if idxP == len(pattern) || idxS == len(str) {
        return match(idxP, idxS)
    }

    // wildcard
    ret := false
    for nextS := idxS; nextS <= len(str); nextS++ {
        if match(idxP+1, nextS) {
            ret = true
        }
    }
    return ret
}

func min(i, j int) int {
    if i < j {
        return i
    } else {
        return j
    }
}
