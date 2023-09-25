// Goal

// Implement a rot13Reader that implements io.Reader
// reads from an io.Reader, modifying the stream by applying
// rot13 substitution cipher to all alphabetical characters.

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// TODO: Refactor to shorten algorithm implementation and use modulo 26

func (rot *rot13Reader) Read(b []byte) (n int, err error) {

	n, err = rot.r.Read(b)

	for i := 0; i < len(b); i++ {

		if b[i] >= []byte("A")[0] && b[i] <= []byte("Z")[0] {

			places := (b[i] + 13) % 91

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

// Rotate by 13 Places

//   - an implementation of the Cesear Substitution Cipher
//     that replaces each letter in the plain text with a letter 13 positions down
//     starting from the beginning if Z is reached

// Caesar Cipher

//   - An ancient cryptographic substitution cipher
//   - Ciphers are algorithms that perform encryption and decryption
//   - substitution ciphers encrypt plain text by replacing the letters
//     with other letters number and symbols
//   - if the plaintext is viewed as a sequence of bits the
//     subtitution evolves replacing the plain text bit pattern with the cipher text bit pattern

// Caesar Cipher Algorithm Expressed Mathematically

//   - Encryption:
//       E = (p + k) % 26

//   - Decryption
//       D = (p - k) % 26
