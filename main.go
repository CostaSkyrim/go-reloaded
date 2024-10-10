package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	filename := os.Args[1]
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file")
		return
	}
	var words []string
	words = splitWhiteSpaces(string(content))

	for i := 0; i < len(words); i++ {
		if words[i] == "(up)" {
			words[i-1] = upper(words[i-1])
			words[i] = ""
		} else if words[i] == "(low)" {
			words[i-1] = lower(words[i-1])
			words[i] = ""
		} else if words[i] == "(cap)" {
			words[i-1] = cap(words[i-1])
			words[i] = ""
		} else if words[i] == "(hex)" {
			words[i-1] = hex(words[i-1])
			words[i] = ""
		} else if words[i] == "(bin)" {
			words[i-1] = bin(words[i-1])
			words[i] = ""
		}
	}

	fmt.Println(words)
}

func lower(s string) string {
	b := []byte(s)
	for i, c := range b {
		if c >= 'A' && c <= 'Z' {
			b[i] = c + ('a' - 'A')
		}
	}
	return string(b)
}

func upper(s string) string {
	b := []byte(s)
	for i, c := range b {
		if c >= 'a' && c <= 'z' {
			b[i] = c - ('a' - 'A')
		}
	}
	return string(b)
}

func cap(s string) string {
	var str string
	first := true
	for i := 0; i < len(s); i++ {
		if isLower(string(s[i])) || isUpper(string(s[i])) {
			if first {
				str += upper(string(s[i]))
				first = false
			} else {
				str += lower(string(s[i]))
			}
		} else {
			str += string(s[i])
			first = !isNumeric(string(s[i]))
		}
	}
	return str
}

func isNumeric(s string) bool {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] < '0' || slice[i] > '9' {
			return false
		}
	}
	return true
}

func isLower(s string) bool {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] < 'a' || slice[i] > 'z' {
			return false
		}
	}
	return true
}

func isUpper(s string) bool {
	slice := []rune(s)
	for i := 0; i < len(slice); i++ {
		if slice[i] < 'A' || slice[i] > 'Z' {
			return false
		}
	}
	return true
}

func splitWhiteSpaces(s string) []string {
	var words []string
	wordStart := -1
	for i := 0; i < len(s); i++ {
		if isSeparator(s[i]) {
			if wordStart != -1 {
				words = append(words, s[wordStart:i])
				wordStart = -1
			}
		} else if wordStart == -1 {
			wordStart = i
		}
		if i == len(s)-1 && wordStart != -1 {
			words = append(words, s[wordStart:i+1])
		}
	}
	return words
}

func isSeparator(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n'
}

func hex(s string) string {
	dec, _ := strconv.ParseInt(s, 16, 64)
	return strconv.FormatInt(dec, 10)
}

func bin(s string) string {
	bin, _ := strconv.ParseInt(s, 2, 64)
	return strconv.FormatInt(bin, 10)
}
