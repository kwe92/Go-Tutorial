package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// TODO: Refactor to shorten algorithm

func (rot *rot13Reader) Read(b []byte) (n int, err error) {

	n, err = rot.r.Read(b)

	for i := 0; i < len(b); i++ {

		if b[i] >= []byte("A")[0] && b[i] <= []byte("Z")[0] {

			places := (b[i] + 13) % 92

			if places < 65 {

				b[i] = places + 65

			} else {

				b[i] = places

			}

		} else if b[i] >= []byte("a")[0] && b[i] <= []byte("z")[0] {

			places := (b[i] + 13) % 123

			if places < 110 {

				b[i] = places + 97

			}
			if places >= 110 {
				b[i] = places
			}

		}

	}

	return
}

func main() {

	s := strings.NewReader("Lbh penpxrq gur pbqr!")

	r := rot13Reader{s}

	io.Copy(os.Stdout, &r)
}
