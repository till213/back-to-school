package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

// DistributeBrackets distributes n opening and closing
// brackets ( and ) in such a way that all valid combinations
// are returned
// E.g. n = 2 : ()(), (())
//      n = 3 : ()()(), (()()), (())(), ()(()), ((())) [but e.g. not )()()(]
func DistributeBrackets(n int) map[string]bool {
	switch n {
	case 0:
		return map[string]bool{}
	case 1:
		return map[string]bool{"()":true}
	default:
		s := DistributeBrackets(n-1)
		r := make(map[string]bool)
		for k := range s {
			r["()" + k] = true
			r["(" + k + ")"] = true
			r[k + "()"] = true
		}
		return r
	}
}

func main() {
	var n int
	var err error
	if len(os.Args) > 1 {
		n, err = strconv.Atoi(os.Args[1])
	} else {
		err = errors.New("No arg")
	}
	if err == nil {
		brackets := DistributeBrackets(n)
		for k := range brackets {
			fmt.Printf("%v, ", k)
		}
		fmt.Println("")
	} else {
		fmt.Println("Parse error: ", err)
	}	
}
