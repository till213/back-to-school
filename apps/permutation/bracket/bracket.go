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

// DistributeParentheses distributes n opening and closing
// parentheses ( and ) in such a way that all valid combinations
// are returned
// E.g. n = 2 : ()(), (())
//      n = 3 : ()()(), (()()), (())(), ()(()), ((())) [but e.g. not )()()(]
// http://blog.gainlo.co/index.php/2016/12/23/uber-interview-questions-permutations-parentheses/
// Note: this does not seem to work. The idea seems right, but:
//       * size of array seems to short: tot = 2*n only (with n=2 we get ()()(())) already, which
//         has len = 8 > 2 * 2 = 4
//       * the second if left < right-condition does not adjust pos, so if left > 0 and left < right
//         then arr[pos] is overwritten      
func DistributeParentheses(left, right, tot int, arr []rune) {
	if left == 0 && right == 0 {
		return
	}
	pos := tot - left - right
	if left > 0 {
		arr[pos] = rune('(')
		DistributeParentheses(left -1, right, tot, arr)
	}
	if left < right {
		arr[pos] = rune(')')
		DistributeParentheses(left, right-1, tot, arr)
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

		tot := 2*n
		arr := make([]rune, tot)
		DistributeParentheses(n, n, tot, arr)
		fmt.Println(arr)

	} else {
		fmt.Println("Parse error: ", err)
	}	
}
