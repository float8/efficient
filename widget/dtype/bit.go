package dtype

import (
	"database/sql/driver"
	"errors"
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

