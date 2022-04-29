package dtype

import (
	"database/sql/driver"
	"errors"
	"strconv"
)

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
	return n.Uint32, nil
}