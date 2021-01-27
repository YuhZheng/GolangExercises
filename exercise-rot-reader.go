package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 rot13Reader) Read(b []byte) (int, error) {
	buffer := make([]byte, len(b))
	cnt, err := rot13.r.Read(buffer)
	if err == nil {
		for i:= range buffer {
			if buffer[i] >= 'a' && buffer[i] <= 'z' {
				if buffer[i] + 13 <= 'z'{
					b[i] = buffer[i] + 13
				} else {
					b[i] = buffer[i] - 13
				}
			}else if buffer[i] >= 'A' && buffer[i] <= 'Z'{
				if buffer[i] + 13 <= 'Z'{
					b[i] = buffer[i] + 13
				} else {
					b[i] = buffer[i] - 13
				}
			}else{
				b[i] = buffer[i]
			}
		}
	}
	return cnt, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

