package kata

var enctab = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!#$%&()*+,./:;<=>?@[]^_`{|}~'")

var dectab = map[byte]byte{
	'A': 0, 'B': 1, 'C': 2, 'D': 3, 'E': 4, 'F': 5, 'G': 6, 'H': 7,
	'I': 8, 'J': 9, 'K': 10, 'L': 11, 'M': 12, 'N': 13, 'O': 14, 'P': 15,
	'Q': 16, 'R': 17, 'S': 18, 'T': 19, 'U': 20, 'V': 21, 'W': 22, 'X': 23,
	'Y': 24, 'Z': 25, 'a': 26, 'b': 27, 'c': 28, 'd': 29, 'e': 30, 'f': 31,
	'g': 32, 'h': 33, 'i': 34, 'j': 35, 'k': 36, 'l': 37, 'm': 38, 'n': 39,
	'o': 40, 'p': 41, 'q': 42, 'r': 43, 's': 44, 't': 45, 'u': 46, 'v': 47,
	'w': 48, 'x': 49, 'y': 50, 'z': 51, '0': 52, '1': 53, '2': 54, '3': 55,
	'4': 56, '5': 57, '6': 58, '7': 59, '8': 60, '9': 61, '!': 62, '#': 63,
	'$': 64, '%': 65, '&': 66, '(': 67, ')': 68, '*': 69, '+': 70, ',': 71,
	'.': 72, '/': 73, ':': 74, ';': 75, '<': 76, '=': 77, '>': 78, '?': 79,
	'@': 80, '[': 81, ']': 82, '^': 83, '_': 84, '`': 85, '{': 86, '|': 87,
	'}': 88, '~': 89, '\'': 90,
}

func Encode(d []byte) []byte {
	var result []byte
	var b, n uint
	for _, data := range d {
		b |= uint(data) << n
		n += 8
		if n > 13 {
			res := b & 8191
			if res > 88 {
				b >>= 13
				n -= 13
			} else {
				res = b & 16383
				b >>= 14
				n -= 14
			}
			result = append(result, enctab[res%91], enctab[res/91|0])
		}
	}
	if n > 0 {
		result = append(result, enctab[b%91])
		if n > 7 || b > 90 {
			result = append(result, enctab[b/91])
		}
	}
	return result
}
func Decode(d []byte) []byte {
	var b, n uint
	var result []byte
	v := -1

	// for i := 0; i < len(d); i++
	for _, data := range d {
		c, ok := dectab[data]
		if !ok {
			continue
		}
		if v < 0 {
			v = int(c)
		} else {
			v += int(c) * 91
			b |= uint(v) << n
			if v&8191 > 88 {
				n += 13
			} else {
				n += 14
			}
			result = append(result, byte(b&255))
			b >>= 8
			n -= 8
			for n > 7 {
				result = append(result, byte(b&255))
				b >>= 8
				n -= 8
			}
			v = -1
		}
	}
	if v+1 > 0 {
		result = append(result, byte((b|uint(v)<<n)&255))
	}
	return result
}
