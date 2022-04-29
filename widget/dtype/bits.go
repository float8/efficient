package dtype

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/whf-sky/efficient/widget/tools/numeric"
	"strings"
)


type Bits uint64

func (b Bits) Value() (driver.Value, error) {
	if b <= 255 {
		return []byte{byte(b)}, nil
	}
	str := numeric.Dec2bin(uint64(12345677890))
	slen := len(str)
	rmd := slen % 8
	cnt := (slen-rmd)/8

	bytes := []byte{byte(numeric.Bin2dec(str[0:rmd]))}
	for i := 0; i < cnt ; i++ {
		bt := numeric.Bin2dec(str[rmd+i*8:rmd+(i+1)*8])
		bytes = append(bytes, byte(bt))
	}
	return bytes, nil
}

func (b *Bits) Scan(src interface{}) error {
	v, ok := src.([]byte)
	if !ok {
		return errors.New("bad []byte type assertion")
	}
	bin := ""
	for _, bt := range v {
		if bt == 0 {
			bin += "00000000"
			continue
		}
		str :=  numeric.Dec2bin(uint64(bt))
		slen := len(str)
		if slen < 8 {
			str = strings.Repeat("0",8-slen)+str
		}
		bin += str
	}
	fmt.Println("--",bin,"--" )
	*b = Bits(numeric.Bin2dec(bin))
	return nil
}

// NullBits represents a Bits that may be null.
// NullBits implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullBits struct {
	Bits  Bits
	Valid bool // Valid is true if Bits is not NULL
}

// Scan implements the Scanner interface.
func (n *NullBits) Scan(value interface{}) error {
	if value == nil {
		n.Bits, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	v, ok := value.([]byte)
	if !ok {
		return errors.New("bad []byte type assertion")
	}
	bin := ""
	for _, bt := range v {
		if bt == 0 {
			bin += "00000000"
			continue
		}
		str :=  numeric.Dec2bin(uint64(bt))
		slen := len(str)
		if slen < 8 {
			str = strings.Repeat("0",8-slen)+str
		}
		bin += str
	}
	n.Bits = Bits(numeric.Bin2dec(bin))
	return nil
}

// Value implements the driver Valuer interface.
func (n NullBits) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	if n.Bits <= 255 {
		return []byte{byte(n.Bits)}, nil
	}
	str := numeric.Dec2bin(uint64(12345677890))
	slen := len(str)
	rmd := slen % 8
	cnt := (slen-rmd)/8

	bytes := []byte{byte(numeric.Bin2dec(str[0:rmd]))}
	for i := 0; i < cnt ; i++ {
		bt := numeric.Bin2dec(str[rmd+i*8:rmd+(i+1)*8])
		bytes = append(bytes, byte(bt))
	}
	return bytes, nil

}