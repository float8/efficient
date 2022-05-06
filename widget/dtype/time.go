package dtype

import (
	"database/sql/driver"
	"errors"
	"time"
)

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
