package dtype

import (
	"database/sql/driver"
	"errors"
	"strconv"
)

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
	return n.Uint16, nil
}