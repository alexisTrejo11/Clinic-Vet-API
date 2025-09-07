package utils

import (
	"database/sql"
	"time"
)

func StringPtrFromNullString(nullString sql.NullString) *string {
	if nullString.Valid {
		return &nullString.String
	}
	return nil
}

func TimePtrFromNullTime(nullTime sql.NullTime) *time.Time {
	if nullTime.Valid {
		return &nullTime.Time
	}
	return nil
}

func Float64FromNullInt64(nullInt sql.NullInt64) float64 {
	if nullInt.Valid {
		return float64(nullInt.Int64)
	}
	return 0
}

func Float64PtrFromNullInt64(nullInt sql.NullInt64) *float64 {
	if nullInt.Valid {
		val := float64(nullInt.Int64)
		return &val
	}
	return nil
}
