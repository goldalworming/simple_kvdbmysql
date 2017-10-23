package utils

import (
	"github.com/sdming/gosnow"
	"io"
	"crypto/rand"
    "github.com/bwmarrin/snowflake"
)

var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()-_=+,.?/:;{}[]`~")
// Create New mail message use MailFrom and MailUser
func NewId2() (newid int64, err error) {
    node, err := snowflake.NewNode(1)
    if err != nil {
        // fmt.Println(err)
        return 0, err
    }
    nvid := node.Generate()
    return nvid.Int64(),nil
}
func NewId() (newid int64, err error) {
	v,err := gosnow.Default()
	if err != nil {
		return 0, err
	}
	vid,err := v.Next()
	if err != nil {
		return 0, err
	}
	nvid := int64(vid)
	return nvid,nil
}
func rand_char(length int, chars []byte) string {
    new_pword := make([]byte, length)
    random_data := make([]byte, length+(length/4)) // storage for random bytes.
    clen := byte(len(chars))
    maxrb := byte(256 - (256 % len(chars)))
    i := 0
    for {
        if _, err := io.ReadFull(rand.Reader, random_data); err != nil {
            panic(err)
        }
        for _, c := range random_data {
            if c >= maxrb {
                continue
            }
            new_pword[i] = chars[c%clen]
            i++
            if i == length {
                return string(new_pword)
            }
        }
    }
    panic("unreachable")
}

func RandomStr(length int) string {
    return rand_char(length, StdChars)
}

func StringInSlice(str string, list []string) bool {
    for _, v := range list {
        if v == str {
            return true
        }
    }
    return false
}
