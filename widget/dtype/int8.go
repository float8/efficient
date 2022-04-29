package dtype

import (
	"database/sql/driver"
	"errors"
	"strconv"
)

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
	return n.Int8, nil
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
	return n.Uint8, nil
}
