package internal

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"strings"
)

/* Usage */
//=================================================================================================//
// buf := make([]byte, 100) // length

// // // Ether: 6byte
// // // d0:37:45:50:0b:06 -> 208 55 69 80 11 6
// // dstMac := [8]byte{
// // 	0xd0, 0x37, 0x45, 0x50,
// // 	0x0b, 0x06, 0x00, 0x00,
// // }
// dstMac := RewriteEtherAddr("d0:37:45:50:0b:06")

// // // Ether: 6byte
// // // d0:37:45:2c:e2:35 -> 208 55 69 44 226 53
// // srcMac := [8]byte{
// // 	0xd0, 0x37, 0x45, 0x2c,
// // 	0xe2, 0x35, 0x00, 0x00,
// // }
// srcMac := RewriteEtherAddr("d0:37:45:2c:e2:35")

// binary.BigEndian.PutUint64(buf[0:8], binary.BigEndian.Uint64(dstMac[0:8]))  // 宛先Mac
// binary.BigEndian.PutUint64(buf[6:14], binary.BigEndian.Uint64(srcMac[0:8])) // 送信元Mac

// fmt.Println("buf:", buf[:12])

/***************************************************************/

// buf := make([]byte, 100) // length

// // IPv6: 16byte
// // fe80::4f9:abd9:c136:d26e -> 254 128 0 0 0 0 0 0 4 249 171 217 193 54 210 110
// srcIP := []byte{
// 	0xfe, 0x80, 0x00, 0x00,
// 	0x00, 0x00, 0x00, 0x00,
// 	0x04, 0xf9, 0xab, 0xd9,
// 	0xc1, 0x36, 0xd2, 0x6e,
// }
// srcIP := RewriteIPv6Layer("fe80::4f9:abd9:c136:d26e")

// // IPv6: 16byte
// // fe80::91e:c59e:f04c:2e3 -> 254 128 0 0 0 0 0 0 9 30 197 158 240 76 2 227
// dstIP := []byte{
// 	0xfe, 0x80, 0x00, 0x00,
// 	0x00, 0x00, 0x00, 0x00,
// 	0x09, 0x1e, 0xc5, 0x9e,
// 	0xf0, 0x4c, 0x02, 0xe3,
// }
// dstIP := RewriteIPv6Layer("fe80::91e:c59e:f04c:2e3")

// buf = append(buf[:0], uint128(srcIP)...)  // 送信元IP
// buf = append(buf[:16], uint128(dstIP)...) // 宛先IP

// fmt.Println("buf:", buf[:32])
//=================================================================================================//
func RewriteEtherAddr(s string) []byte {
	replaced := strings.Replace(s, ":", "", -1)
	p, err := hex.DecodeString(replaced)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func RewriteIPv6Layer(s string) []byte {
	ip := net.ParseIP(s)
	ipv6 := ip.To16()

	replaced := fmt.Sprintf("%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x%02x",
		ipv6[0], ipv6[1], ipv6[2], ipv6[3],
		ipv6[4], ipv6[5], ipv6[6], ipv6[7],
		ipv6[8], ipv6[9], ipv6[10], ipv6[11],
		ipv6[12], ipv6[13], ipv6[14], ipv6[15],
	)

	p, err := hex.DecodeString(replaced)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

type Uint128 struct {
	Hi uint64
	Lo uint64
}

func FromBytes(b []byte) Uint128 {
	hi := binary.BigEndian.Uint64(b[:8])
	lo := binary.BigEndian.Uint64(b[8:])
	return Uint128{hi, lo}
}

func (u Uint128) GetBytes() []byte {
	buf := make([]byte, 16)
	binary.BigEndian.PutUint64(buf[:8], u.Hi)
	binary.BigEndian.PutUint64(buf[8:], u.Lo)
	return buf
}

func uint128(b []byte) []byte {
	// https://github.com/cockroachdb/cockroach/blob/master/pkg/util/uint128/uint128.go
	i := FromBytes(b)
	if !bytes.Equal(i.GetBytes(), b) {
		fmt.Println("[ERROR] incorrect bytes representation for num:", i)
	}
	p := i.GetBytes()
	return p
}
