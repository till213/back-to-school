package main

import (
	"fmt"
)

const Dash = '-'

func convert(ch byte) byte {
    if 'a' <= ch && ch <= 'z' {
		return 'A' + (ch - 'a')
	}
    return ch
}

func licenseKeyFormatting(S string, K int) string {
    raw := make([]byte, len(S)*2)
    
    groupSize := 0
	writePos := len(raw) - 1
	doDash := false
    for readPos := len(S)-1; readPos >= 0; readPos-- {
        if groupSize == K {
            doDash = true
			groupSize = 0
		}
        // ignore original dashes
        if S[readPos] != Dash {
			// Do a dash first, if one is pending
			if doDash {
				raw[writePos] = Dash
				writePos--
				doDash = false
			}
            raw[writePos] = convert(S[readPos])
			writePos--
			groupSize++
		}
    }
    
    res := string(raw[writePos+1:len(raw)])    
    return res 
}

func main() {
	orig := "5F3Z-2e-9-w"
	lc := licenseKeyFormatting(orig, 4)
	fmt.Printf("key: '%s'\n", lc)

	orig = "aaaa"
	lc = licenseKeyFormatting(orig, 2)
	fmt.Printf("key: '%s'\n", lc)

	orig = "--a-a-a-a--"
	lc = licenseKeyFormatting(orig, 2)
	fmt.Printf("key: '%s'\n", lc)
}