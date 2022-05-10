package database

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/whf-sky/efficient/widget/tools/numeric"
	"strconv"
	"strings"
	"time"
)

// Bit is an implementation of a bool for the MySQL type BIT(1).
// This type allows you to avoid wasting an entire byte for MySQL's boolean type TINYINT.
type Bit bool

// Value implements the driver.Valuer interface,
// and turns the BitBool into a bitfield (BIT(1)) for MySQL storage.
func (b Bit) Value() (driver.Value, error) {
	if b {
		return []byte{1}, nil
	} else {
		return []byte{0}, nil
	}
}

// Scan implements the sql.Scanner interface,
// and turns the bitfield incoming from MySQL into a BitBool
func (b *Bit) Scan(src interface{}) error {
	v, ok := src.([]byte)
	if !ok {
		return errors.New("bad []byte type assertion")
	}
	*b = v[0] == 1
	return nil
}


// NullBit represents a Bit that may be null.
// NullBit implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullBit struct {
	Bit  Bit
	Valid bool // Valid is true if Bit is not NULL
}

// Scan implements the Scanner interface.
func (n *NullBit) Scan(value interface{}) error {
	if value == nil {
		n.Bit, n.Valid = false, false
		return nil
	}
	n.Valid = true
	v, ok := value.([]byte)
	if !ok {
		return errors.New("bad []byte type assertion")
	}
	n.Bit = v[0] == 1
	return nil
}

// Value implements the driver Valuer interface.
func (n NullBit) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	if n.Bit {
		return []byte{1}, nil
	}
	return []byte{0}, nil
}



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



type Date struct {
	time.Time
}

func (d *Date) Scan(src interface{}) error {
	v, ok := src.([]byte)
	if !ok {
		return errors.New("bad string type assertion")
	}
	time, _ := time.ParseInLocation("2006-01-02", string(v), time.Local)
	d.Time = time
	return nil
}

func (d Date) Value() (driver.Value, error) {
	str := d.Format("2006-01-02")
	if str == "0001-01-01" {
		return "0000-00-00", nil
	}
	return str, nil
}

func (d Date) String() string {
	str := d.Format("2006-01-02")
	if str == "0001-01-01" {
		return "0000-00-00"
	}
	return str
}


type NullDate struct {
	Date time.Time
	Valid bool // Valid is true if NullTime is not NULL
}

func (n *NullDate) Scan(src interface{}) error {
	if src == nil {
		n.Date, n.Valid = time.Time{}, false
		return nil
	}
	n.Valid = true
	v, ok := src.([]byte)
	if !ok {
		return errors.New("bad string type assertion")
	}
	time, _ := time.ParseInLocation("2006-01-02", string(v), time.Local)
	n.Date = time
	return nil
}

func (n NullDate) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	str := n.Date.Format("2006-01-02")
	if str == "0001-01-01" {
		return "0000-00-00", nil
	}
	return str, nil
}

func (n NullDate) String() string {
	str := n.Date.Format("2006-01-02")
	if str == "0001-01-01" {
		return "0000-00-00"
	}
	return str
}



// Nullint8 represents a int8 that may be null.
// Nullint8 implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullInt8 struct {
	Int8  int8
	Valid bool // Valid is true if uint8 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullInt8) Scan(value interface{}) error {
	if value == nil {
		n.Int8, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	v, ok := value.([]byte)
	if !ok {
		return errors.New("bad []byte type assertion")
	}
	num, _ := strconv.Atoi(string(v))
	n.Int8 = int8(num)
	return nil
}

// Value implements the driver Valuer interface.
func (n NullInt8) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return int64(n.Int8), nil
}

// NullUint8 represents a uint8 that may be null.
// NullUint8 implements the Scanner interface so
// it can be used as a scan destination, similar to Nullint8.
type NullUint8 struct {
	Uint8 uint8
	Valid bool // Valid is true if uint8 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullUint8) Scan(value interface{}) error {
	if value == nil {
		n.Uint8, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	v, ok := value.([]byte)
	if !ok {
		return errors.New("bad []byte type assertion")
	}
	num, _ := strconv.Atoi(string(v))
	n.Uint8 = uint8(num)
	return nil
}

// Value implements the driver Valuer interface.
func (n NullUint8) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return int64(n.Uint8), nil
}



// NullUint16 represents a uint8 that may be null.
// NullUint16 implements the Scanner interface so
// it can be used as a scan destination, similar to NullInt16.
type NullUint16 struct {
	Uint16 uint16
	Valid bool // Valid is true if uint8 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullUint16) Scan(value interface{}) error {
	if value == nil {
		n.Uint16, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	v, ok := value.([]byte)
	if !ok {
		return errors.New("bad []byte type assertion")
	}
	num, _ := strconv.Atoi(string(v))
	n.Uint16 = uint16(num)
	return nil
}

// Value implements the driver Valuer interface.
func (n NullUint16) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return int64(n.Uint16), nil
}


// NullUint32 represents a uint8 that may be null.
// NullUint32 implements the Scanner interface so
// it can be used as a scan destination, similar to NullInt32.
type NullUint32 struct {
	Uint32 uint32
	Valid bool // Valid is true if NullUint32 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullUint32) Scan(value interface{}) error {
	if value == nil {
		n.Uint32, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	v, ok := value.([]byte)
	if !ok {
		return errors.New("bad []byte type assertion")
	}
	num, _ := strconv.Atoi(string(v))
	n.Uint32 = uint32(num)
	return nil
}

// Value implements the driver Valuer interface.
func (n NullUint32) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return int64(n.Uint32), nil
}


// NullUint64 represents a uint8 that may be null.
// NullUint64 implements the Scanner interface so
// it can be used as a scan destination, similar to NullInt64.
type NullUint64 struct {
	Uint64 uint64
	Valid bool // Valid is true if uint8 is not NULL
}

// Scan implements the Scanner interface.
func (n *NullUint64) Scan(value interface{}) error {
	if value == nil {
		n.Uint64, n.Valid = 0, false
		return nil
	}
	n.Valid = true
	v, ok := value.([]byte)
	if !ok {
		return errors.New("bad []byte type assertion")
	}
	num, _ := strconv.ParseInt(string(v), 10, 64)
	n.Uint64 = uint64(num)
	return nil
}

// Value implements the driver Valuer interface.
func (n NullUint64) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Uint64, nil
}


type Set []string

func (s Set) Value() (driver.Value, error) {
	return strings.Join(s, ","), nil
}

func (s *Set) Scan(src interface{}) error {
	v, ok := src.([]byte)
	if !ok {
		return errors.New("bad string type assertion")
	}
	*s = strings.Split(string(v), ",")
	return nil
}

func (s Set) String() string {
	return strings.Join(s, ",")
}

// NullSet represents a Set that may be null.
// NullSet implements the Scanner interface so
// it can be used as a scan destination, similar to NullString.
type NullSet struct {
	Set  Set
	Valid bool // Valid is true if Set is not NULL
}

// Scan implements the Scanner interface.
func (n *NullSet) Scan(value interface{}) error {
	if value == nil {
		n.Set, n.Valid = []string{}, false
		return nil
	}
	n.Valid = true
	v, ok := value.([]byte)
	if !ok {
		return errors.New("bad Set type assertion")
	}
	n.Set = strings.Split(string(v), ",")
	return nil
}

// Value implements the driver Valuer interface.
func (n NullSet) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return strings.Join(n.Set, ","), nil
}


type Time struct {
	time.Time
}

func (t *Time) Scan(src interface{}) error {
	v, ok := src.([]byte)
	if !ok {
		return errors.New("bad string type assertion")
	}
	time, _ := time.ParseInLocation("2006-01-02 15:04:05", string(v), time.Local)
	t.Time = time
	return nil
}

func (t Time) Value() (driver.Value, error) {
	str := t.Format("2006-01-02 15:04:05")
	if str == "0001-01-01 00:00:00" {
		return "0000-00-00 00:00:00", nil
	}
	return str, nil
}

func (t Time) String() string {
	str := t.Format("2006-01-02 15:04:05")
	if str == "0001-01-01 00:00:00" {
		return "0000-00-00 00:00:00"
	}
	return str
}

type NullTime struct {
	Time time.Time
	Valid bool // Valid is true if NullTime is not NULL
}


func (n *NullTime) Scan(src interface{}) error {
	if src == nil {
		n.Time, n.Valid = time.Time{}, false
		return nil
	}
	n.Valid = true
	v, ok := src.([]byte)
	if !ok {
		return errors.New("bad string type assertion")
	}
	time, _ := time.ParseInLocation("2006-01-02 15:04:05", string(v), time.Local)
	n.Time = time
	return nil
}

func (n NullTime) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}

	str := n.Time.Format("2006-01-02 15:04:05")
	if str == "0001-01-01 00:00:00" {
		return "0000-00-00 00:00:00", nil
	}
	return str, nil
}

func (n NullTime) String() string {
	str := n.Time.Format("2006-01-02 15:04:05")
	if str == "0001-01-01 00:00:00" {
		return "0000-00-00 00:00:00"
	}
	return str
}
