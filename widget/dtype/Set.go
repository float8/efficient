package dtype

import (
	"database/sql/driver"
	"errors"
	"strings"
)

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