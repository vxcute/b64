package b64

import "strings"

var base64Table = []byte{
	'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L',
	'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X',
	'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j',
	'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v',
	'w', 'x', 'y', 'z', '0', '1', '2', '3', '4', '5', '6', '7',
	'8', '9', '+', '/',
}

func GetB64Index(b byte) int {
	if b >= 'A' && b <= 'Z' {
		return int(b - 'A')
	} else if b >= 'a' && b <= 'z' {
		return int((b - 'a') + 26)
	} else if b >= '0' && b <= '9' {
		return int((b - '0') + 52)
	} else if b == '+' {
		return 62
	} else if b == '/' {
		return 63
	}

	return -1
}

func Base64Encode(s string) string {
	var encoded strings.Builder

	for i := 0; i < len(s); i += 3 {
		if i == len(s)-1 {
			b := s[i]
			c := uint16(b) << 4
			encoded.WriteByte(base64Table[(c>>6)&0x3F])
			encoded.WriteByte(base64Table[(c>>0)&0x3F])
			encoded.WriteString("==")
			break
		}

		if i == len(s)-2 {
			b1 := s[i]
			b2 := s[i+1]
			c := ((uint32(b1) << 8) | uint32(b2)) << 2
			encoded.WriteByte(base64Table[(c>>12)&0x3F])
			encoded.WriteByte(base64Table[(c>>6)&0x3F])
			encoded.WriteByte(base64Table[(c>>0)&0x3F])
			encoded.WriteByte('=')
			break
		}

		b1 := s[i]
		b2 := s[i+1]
		b3 := s[i+2]
		c := (uint32(b1) << 16) | (uint32(b2) << 8) | (uint32(b3) << 0)

		encoded.WriteByte(base64Table[(c>>18)&0x3F])
		encoded.WriteByte(base64Table[(c>>12)&0x3F])
		encoded.WriteByte(base64Table[(c>>6)&0x3F])
		encoded.WriteByte(base64Table[(c>>0)&0x3F])
	}

	return encoded.String()
}

func Base64Decode(s string) string {
	var decoded strings.Builder

	for i := 0; i < len(s); i += 4 {
		b1 := s[i]
		b2 := s[i+1]
		b3 := s[i+2]
		b4 := s[i+3]

		if b3 != '=' && b4 != '=' {
			b1Idx := GetB64Index(b1)
			b2Idx := GetB64Index(b2)
			b3Idx := GetB64Index(b3)
			b4Idx := GetB64Index(b4)

			c := (uint32(b1Idx) << 18) |
				(uint32(b2Idx) << 12) |
				(uint32(b3Idx) << 6) |
				(uint32(b4Idx) << 0)

			decoded.WriteByte(byte((c >> 16) & 0xFF))
			decoded.WriteByte(byte((c >> 8) & 0xFF))
			decoded.WriteByte(byte((c >> 0) & 0xFF))
		} else if b3 == '=' && b4 == '=' {
			b1Idx := GetB64Index(b1)
			b2Idx := GetB64Index(b2)
			c := ((uint16(b1Idx) << 6) | uint16(b2Idx)) << 4
			decoded.WriteByte(byte(c>>8) & 0xFF)
		} else if b4 == '=' {
			b1Idx := GetB64Index(b1)
			b2Idx := GetB64Index(b2)
			b3Idx := GetB64Index(b3)
			c := ((uint32(b1Idx) << 12) | (uint32(b2Idx) << 6) | (uint32(b3Idx) << 0)) << 4
			decoded.WriteByte(byte(c>>14) & 0xFF)
			decoded.WriteByte(byte(c>>6) & 0xFF)
		}
	}

	return decoded.String()
}
