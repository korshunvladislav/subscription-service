package utils

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

const DateLayout = "02-01-2006"
const monthYearLayout = "01-2006"

type MonthYear struct {
	time.Time
}

func (m *MonthYear) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse(monthYearLayout, s)
	if err != nil {
		return err
	}
	m.Time = t
	return nil
}

func (m MonthYear) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", m.Format(monthYearLayout))), nil
}

func (m MonthYear) Value() (driver.Value, error) {
	return m.Time, nil
}

func (m *MonthYear) Scan(value interface{}) error {
	if val, ok := value.(time.Time); ok {
		m.Time = val
		return nil
	}
	return fmt.Errorf("cannot scan value %v into MonthYear", value)
}
