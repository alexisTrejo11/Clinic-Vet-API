package mapper

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"fmt"
	"math/big"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

type SqlcFieldMapper struct {
	Primitive     PrimitiveMapper
	ValueObject   ValueObjectMapper
	PgText        PgTextMapper
	PgInt2        PgInt2Mapper
	PgBool        PgBoolMapper
	PgInt4        PgInt4Mapper
	PgDate        PgDateMapper
	PgTimestamptz PgTimestamptzMapper
	PgNumeric     NumericMapper
}

type PgInt4Mapper struct{}

type PgDateMapper struct{}

type NumericMapper struct{}

func NewSqlcFieldMapper() SqlcFieldMapper {
	return SqlcFieldMapper{
		Primitive:     PrimitiveMapper{},
		PgText:        PgTextMapper{},
		PgInt4:        PgInt4Mapper{},
		PgDate:        PgDateMapper{},
		PgTimestamptz: PgTimestamptzMapper{},
		PgNumeric:     NumericMapper{},
		ValueObject:   ValueObjectMapper{},
		PgInt2:        PgInt2Mapper{},
		PgBool:        PgBoolMapper{},
	}
}

type PgInt2Mapper struct{}

func (m *PgInt2Mapper) ToIntPtr(pgType pgtype.Int2) *int {
	if pgType.Valid {
		i := int(pgType.Int16)
		return &i
	}
	return nil
}

func (m *PgInt2Mapper) FromInt(i *int) pgtype.Int2 {
	if i != nil {
		return pgtype.Int2{Int16: int16(*i), Valid: true}
	}
	return pgtype.Int2{Valid: false}
}

func (m *PgInt2Mapper) ToInt(pgType pgtype.Int2) int {
	if pgType.Valid {
		return int(pgType.Int16)
	}
	return 0
}

func (m *PgInt2Mapper) FromIntValue(i int) pgtype.Int2 {
	if i != 0 {
		return pgtype.Int2{Int16: int16(i), Valid: true}
	}
	return pgtype.Int2{Valid: false}
}

type PgBoolMapper struct{}

func (m *PgBoolMapper) ToBoolPtr(pgType pgtype.Bool) *bool {
	if pgType.Valid {
		return &pgType.Bool
	}
	return nil
}

func (m *PgBoolMapper) FromBool(b *bool) pgtype.Bool {
	if b != nil {
		return pgtype.Bool{Bool: *b, Valid: true}
	}
	return pgtype.Bool{Valid: false}
}

func (m *PgBoolMapper) FromBoolValue(b bool) pgtype.Bool {
	return pgtype.Bool{Bool: b, Valid: true}
}

// Pg Type Mapper
type PgTextMapper struct{}

func (m *PgTextMapper) ToStringPtr(pgType pgtype.Text) *string {
	if pgType.Valid {
		return &pgType.String
	}
	return nil
}

func (m *PgTextMapper) ToString(pgType pgtype.Text) string {
	if pgType.Valid {
		return pgType.String
	}
	return ""
}

func (m *PgTextMapper) FromStringPtr(s *string) pgtype.Text {
	if s != nil {
		return pgtype.Text{String: *s, Valid: true}
	}
	return pgtype.Text{Valid: false}
}

func (m *PgTextMapper) FromString(s string) pgtype.Text {
	if s != "" {
		return pgtype.Text{String: s, Valid: true}
	}
	return pgtype.Text{Valid: false}
}

func (m *PgTextMapper) ToPhoneNumberPtr(t pgtype.Text) *valueobject.PhoneNumber {
	if t.Valid && t.String != "" {
		phone := valueobject.NewPhoneNumberNoErr(t.String)
		return &phone
	}
	return nil
}

func (m *PgInt4Mapper) ToInt32Ptr(pgType pgtype.Int4) *int32 {
	if pgType.Valid {
		return &pgType.Int32
	}
	return nil
}

func (m *PgInt4Mapper) FromUintPtr(i *uint) pgtype.Int4 {
	if i != nil {
		return pgtype.Int4{Int32: int32(*i), Valid: true}
	}
	return pgtype.Int4{Valid: false}
}

func (m *PgInt4Mapper) FromUint(i uint) pgtype.Int4 {
	if i != 0 {
		return pgtype.Int4{Int32: int32(i), Valid: true}
	}
	return pgtype.Int4{Valid: false}
}

func (m *PgInt4Mapper) ToUint(pgType pgtype.Int4) uint {
	if pgType.Valid {
		return uint(pgType.Int32)
	}
	return 0
}

func (m *PgInt4Mapper) ToUserIDPtr(pgType pgtype.Int4) *valueobject.UserID {
	if pgType.Valid {
		id := valueobject.NewUserID(uint(pgType.Int32))
		return &id
	}
	return nil
}

func (m *PgInt4Mapper) FromUserIDPtr(id *valueobject.UserID) pgtype.Int4 {
	if id != nil {
		return pgtype.Int4{Int32: int32(id.Value()), Valid: true}
	}
	return pgtype.Int4{Valid: false}
}

func (m *PgInt4Mapper) ToCustomerIDPtr(pgType pgtype.Int4) *valueobject.CustomerID {
	if pgType.Valid {
		id := valueobject.NewCustomerID(uint(pgType.Int32))
		return &id
	}
	return nil
}

func (m *PgInt4Mapper) ToEmployeeIDPtr(pgType pgtype.Int4) *valueobject.EmployeeID {
	if pgType.Valid {
		id := valueobject.NewEmployeeID(uint(pgType.Int32))
		return &id
	}
	return nil
}

func (m *PgInt4Mapper) ToEmployeeID(pgType pgtype.Int4) valueobject.EmployeeID {
	if pgType.Valid {
		return valueobject.NewEmployeeID(uint(pgType.Int32))
	}
	return valueobject.EmployeeID{}
}

func (m *PgInt4Mapper) ToPetID(pgType pgtype.Int4) valueobject.PetID {
	if pgType.Valid {
		return valueobject.NewPetID(uint(pgType.Int32))
	}
	return valueobject.PetID{}
}

func (m *PgDateMapper) ToTimePtr(pgType pgtype.Date) *time.Time {
	if pgType.Valid {
		return &pgType.Time
	}
	return nil
}

func (m *PgDateMapper) ToTime(pgType pgtype.Date) time.Time {
	if pgType.Valid {
		return pgType.Time
	}
	return time.Time{}
}

func (m *PgDateMapper) FromTimePtr(t *time.Time) pgtype.Date {
	if t != nil && *t != (time.Time{}) {
		return pgtype.Date{Time: *t, Valid: true}
	}
	return pgtype.Date{Valid: false}
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

func (m *SqlcFieldMapper) StringToPgText(s string) pgtype.Text {
	if s != "" {
		return pgtype.Text{String: s, Valid: true}
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

func (m *SqlcFieldMapper) UserIDPtrToInt32(id *valueobject.UserID) pgtype.Int4 {
	if id != nil {
		return pgtype.Int4{Int32: int32(id.Value()), Valid: true}
	}
	return pgtype.Int4{Valid: false}
}

func (m *SqlcFieldMapper) StringPtrToString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func (m *SqlcFieldMapper) Int32ToUserIDPtr(i int32) *valueobject.UserID {
	if i != 0 {
		id := valueobject.NewUserID(uint(i))
		return &id
	}
	return nil
}

type PgTimestamptzMapper struct{}

func (r *PgTimestamptzMapper) FromTime(t time.Time) pgtype.Timestamptz {
	if t != (time.Time{}) {
		return pgtype.Timestamptz{Time: t, Valid: true}
	}
	return pgtype.Timestamptz{Valid: false}
}

func (r *PgTimestamptzMapper) FromTimePtr(t *time.Time) pgtype.Timestamptz {
	if t != nil && *t != (time.Time{}) {
		return pgtype.Timestamptz{Time: *t, Valid: true}
	}
	return pgtype.Timestamptz{Valid: false}
}

func (r *PgTimestamptzMapper) ToTime(t pgtype.Timestamptz) time.Time {
	if t.Valid {
		return t.Time
	}
	return time.Time{}
}

func (r *PgTimestamptzMapper) ToTimePtr(t pgtype.Timestamptz) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

func (r *NumericMapper) ToMoney(num pgtype.Numeric, currency string) (valueobject.Money, error) {
	if !num.Valid {
		return valueobject.Money{}, fmt.Errorf("numeric value is null")
	}

	var amount float64
	if err := num.Scan(&amount); err != nil {
		return valueobject.Money{}, fmt.Errorf("failed to convert numeric to amount: %w", err)
	}

	return valueobject.NewMoney(valueobject.NewDecimalFromFloat(amount), currency), nil
}

func (r *NumericMapper) FromMoneyPtr(money *valueobject.Money) pgtype.Numeric {
	if money != nil {
		amount := money.Amount()
		return pgtype.Numeric{
			Int:   big.NewInt(int64(amount.Float64())),
			Valid: true,
		}
	}
	return pgtype.Numeric{Valid: false}
}

func (r *NumericMapper) FromMoney(money valueobject.Money) pgtype.Numeric {
	amount := money.Amount()
	return pgtype.Numeric{
		Int:   big.NewInt(int64(amount.Float64())),
		Valid: true,
	}
}

func (r *NumericMapper) ToNumeric(money valueobject.Money) pgtype.Numeric {
	amount := money.Amount()
	return pgtype.Numeric{
		Int:   big.NewInt(int64(amount.Float64())),
		Valid: true,
	}
}
