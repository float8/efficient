package dtype

import (
	"database/sql/driver"
	"errors"
	"strconv"
)

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