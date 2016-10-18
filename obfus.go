// obfus project obfus.go
package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

const (
	EAX = uint8(unsafe.Sizeof(true))
	ONE = "EAX"
)

//Превращаем число в последовательность смещений
func getNumber(n byte) (buf string) {
	var arr []byte
	for n > EAX {
		if n%2 == EAX {
			arr = append(arr, EAX)
		} else {
			arr = append(arr, 0)
		}
		n = n >> EAX
	}

	buf = ONE
	rand.Seed(time.Now().Unix())

	for i := len(arr) - 1; i >= 0; i-- {
		buf = fmt.Sprintf("%s<<%s", buf, ONE)
		if arr[i] == EAX {
			if rand.Intn(2) == 0 {
				buf = fmt.Sprintf("(%s^%s)", buf, ONE)
			} else {
				buf = fmt.Sprintf("(%s|%s)", buf, ONE)
			}
		}
	}
	return buf
}

//Генерим код, вывода строки
func TextToCode(txt string) string {
	b := []byte(txt)
	tmp := "var str []byte\n"
	for _, item := range b {
		tmp = fmt.Sprintf("%s\nstr = append(str, %s)", tmp, getNumber(item))
	}
	tmp += "\nfmt.Println(string(str))"
	return tmp
}

func main() {
	fmt.Println(TextToCode("Author: @GH0st3rs"))
	//fmt.Scanln()
}
