package main

import (
	"fmt"
	"strconv"
)

func writebuff(buf []byte, writer, count int) (int, []byte) {
	tmp := strconv.AppendInt(nil, int64(count), 10)
	copy(buf[writer:], tmp)
	return writer + len(tmp), buf
}

// stringCompresExtraBuff - use extra buffer
func stringCompresExtraBuff(s []byte) []byte {
	if len(s) == 0 {
		return []byte{}
	}

	buf := make([]byte, len(s)*2)
	writer := 0
	count := 1
	c := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] == c {
			count++
		} else {
			writer, buf = writebuff(buf, writer, count)
			buf[writer] = c
			writer++
			c = s[i]
			count = 1
		}
	}

	// Flush the final run.
	writer, buf = writebuff(buf, writer, count)
	buf[writer] = c
	writer++

	return buf[0:writer]
}

// stringCompress O(n) - ensure original s never grows
// no extra space, 1 two passes
func stringCompress(buf []byte) []byte {
	if len(buf) == 0 {
		return []byte{}
	}
	writer := 0
	count := 1
	c := buf[0]
	for i := 1; i < len(buf); i++ {
		if buf[i] == c {
			count++
			continue
		}
		if buf[i] != c {
			if count > 1 {
				writer, buf = writebuff(buf, writer, count)
				buf[writer] = c
			}
			writer++
			c = buf[i]
			count = 1
		}
	}
	// Flush the final run.
	if count > 1 {
		writer, buf = writebuff(buf, writer, count)
		buf[writer] = c
	}
	writer++

	return buf[0:writer]
}

// stringCompressTwoPass second pass (backwards?)
func stringCompressTwoPass(s []byte) []byte {
	return s
}

func main() {

	// fmt.Printf("%s\n", stringCompresExtraBuff([]byte("abc")))
	// fmt.Printf("%s\n", stringCompresExtraBuff([]byte("aabbcc")))
	// fmt.Printf("%s\n", stringCompresExtraBuff([]byte("aaaabbcc")))

	fmt.Printf("%s\n", stringCompress([]byte("abc")))
	fmt.Printf("%s\n", stringCompress([]byte("aabbcc")))
	fmt.Printf("%s\n", stringCompress([]byte("aaaabbcc")))

}
