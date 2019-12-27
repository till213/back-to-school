package main

import "fmt"

const blank = 32
const perc = 50
const two = 48
const zero = 37

func urlify(url []rune, len int) {
	for i := 0; i < len; i++ {
		if url[i] == blank {
			copy(url[i+3:], url[i+1:])
			url[i] = perc
			url[i+1] = two
			url[i+2] = zero
			i += 3
			len += 2
		}
	}
}

func urlify2(url []rune, trueLength int) {
	var nofBlanks int

	for i := 0; i < trueLength; i++ {
		if url[i] == blank {
			nofBlanks++
		}
	}

	writePos := trueLength + nofBlanks*2
	if trueLength < len(url) {
		url[trueLength] = 0
	}
	for j := trueLength - 1; j >= 0; j-- {
		if url[j] == blank {
			url[writePos-1] = zero
			url[writePos-2] = two
			url[writePos-3] = perc
			writePos -= 3
		} else {
			url[writePos-1] = url[j]
			writePos--
		}
	}
}

func main() {
	str := "Mr John Smith    "
	url := []rune(str)
	fmt.Println("Original:", url, "Capacity", cap(url))
	urlify(url, 13)
	fmt.Println("Urlified 1:", string(url), "Capacity", cap(url))
	url = []rune(str)
	urlify2(url, 13)
	fmt.Println("Urlified 2:", string(url), "Capacity", cap(url))
}
