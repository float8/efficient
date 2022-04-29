package dtype

import (
	"database/sql/driver"
	"errors"
	"time"
)

type Date struct {
	time.Time
}

func (d Date) Value() (driver.Value, error) {
	return d.Format("2006-01-02"), nil
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

func (d Date) String() string {
	return d.Format("2006-01-02")
}


type NullDate struct {
	Date time.Time
	Valid bool // Valid is true if NullTime is not NULL
}

func (n NullDate) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return n.Date.Format("2006-01-02"), nil
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

func (t NullDate) String() string {
	return t.Date.Format("2006-01-02")
}

