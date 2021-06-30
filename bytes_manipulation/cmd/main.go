package main

import (
	"bytes"
	"fmt"
)

// bytes manipulation is same as string manipulation

// what is buffer and how to use that ?

func splitter(s, substr string) (bef, mid, aft string) {
	hay, needle := []byte(s), []byte(substr)
	bef = string(hay)
	if idx := bytes.Index(hay, needle); idx > -1 {
		bef = s[:idx]
		mid = s[idx : idx+len(substr)]
		aft = s[idx+len((substr)):]
	}
	return bef, mid, aft
}

func main() {

	a, b := []byte("abc"), []byte("def")
	if !bytes.Equal(a, b) {
	} // bytes are not equal
	if bytes.Contains([]byte("abc"), []byte("a")) {
	} // contains the byte
	fmt.Println(bytes.Index([]byte("i m dead"), []byte("dead")))

	x := []byte("i m dead")[4 : 4+len("dead")]
	fmt.Println(x)

	z := append(a, b...)
	fmt.Println(string(z), string(a))

	bef, mid, aft := splitter("ok zz ok", "zz")
	fmt.Println(bef)
	fmt.Println(mid)
	fmt.Println(aft)

}
