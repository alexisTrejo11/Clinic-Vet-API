package mapper

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type SqlcFieldMapper struct{}

func (m *SqlcFieldMapper) PgTextToPtr(pgType pgtype.Text) *string {
	if pgType.Valid {
		return &pgType.String
	}
	return nil
}

func (m *SqlcFieldMapper) PgInt4ToPtr(pgType pgtype.Int4) *int32 {
	if pgType.Valid {
		return &pgType.Int32
	}
	return nil
}

func (m *SqlcFieldMapper) PgDateToPtr(pgType pgtype.Date) *time.Time {
	if pgType.Valid {
		return &pgType.Time
	}
	return nil
}

func (m *SqlcFieldMapper) Int32PtrToInt32(i *int32) int32 {
	if i != nil {
		return *i
	}
	return 0
}

func (m *SqlcFieldMapper) DewormIDPtrToPgInt4(id *valueobject.DewormID) pgtype.Int4 {
	if id != nil {
		return pgtype.Int4{Int32: int32(id.Value()), Valid: true}
	}
	return pgtype.Int4{Valid: false}
}

func (m *SqlcFieldMapper) StringPtrToPgText(s *string) pgtype.Text {
	if s != nil {
		return pgtype.Text{String: *s, Valid: true}
	}
	return pgtype.Text{Valid: false}
}

func (m *SqlcFieldMapper) TimeToPgDate(t time.Time) pgtype.Date {
	if t != (time.Time{}) {
		return pgtype.Date{Time: t, Valid: true}
	}
	return pgtype.Date{Valid: false}
}

func (m *SqlcFieldMapper) TimePtrToPgDate(t *time.Time) pgtype.Date {
	if t != nil && *t != (time.Time{}) {
		return pgtype.Date{Time: *t, Valid: true}
	}
	return pgtype.Date{Valid: false}
}

func (m *SqlcFieldMapper) PgTimeStampToTime(t pgtype.Timestamptz) time.Time {
	if t.Valid {
		return t.Time
	}
	return time.Time{}
}

func (m *SqlcFieldMapper) TimePtrToPgTypestamp(t *time.Time) pgtype.Timestamptz {
	if t != nil && *t != (time.Time{}) {
		return pgtype.Timestamptz{Time: *t, Valid: true}
	}
	return pgtype.Timestamptz{Valid: false}
}

func (m *SqlcFieldMapper) UintToPgInt4(i uint) pgtype.Int4 {
	if i != 0 {
		return pgtype.Int4{Int32: int32(i), Valid: true}
	}
	return pgtype.Int4{Valid: false}
}

func (m *SqlcFieldMapper) PetIDPtrToInt32(id *valueobject.PetID) int32 {
	if id != nil {
		return int32(id.Value())
	}
	return 0
}

func (m *SqlcFieldMapper) EmployeeIDPtrToInt32(id *valueobject.EmployeeID) int32 {
	if id != nil {
		return int32(id.Value())
	}
	return 0
}

func (m *SqlcFieldMapper) DewormIDToInt32(id valueobject.DewormID) int32 {
	return int32(id.Value())
}

func (m *SqlcFieldMapper) StringPtrToString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
