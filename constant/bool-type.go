package constant

import "database/sql/driver"

type BoolType string

const (
	True  BoolType = "True"
	False BoolType = "False"
)

func (ct *BoolType) Scan(value interface{}) error {
	*ct = BoolType(value.([]byte))
	return nil
}

func (ct *BoolType) Value() (driver.Value, error) {
	return string(*ct), nil
}
