package mapper

import "time"

type PrimitiveMapper struct{}

func (m *PrimitiveMapper) StringPtrToString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}

func (m *PrimitiveMapper) StringToStringPtr(s string) *string {
	if s != "" {
		return &s
	}
	return nil
}

func (m *PrimitiveMapper) Int32PtrToInt32(i *int32) int32 {
	if i != nil {
		return *i
	}
	return 0
}

func (m *PrimitiveMapper) Int32ToInt32Ptr(i int32) *int32 {
	if i != 0 {
		return &i
	}
	return nil
}

func (m *PrimitiveMapper) BoolPtrToBool(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}

func (m *PrimitiveMapper) BoolToBoolPtr(b bool) *bool {
	if b {
		return &b
	}
	return nil
}

func (m *PrimitiveMapper) TimePtrToTime(t *time.Time) time.Time {
	if t != nil {
		return *t
	}
	return time.Time{}
}

func (m *PrimitiveMapper) TimeToTimePtr(t time.Time) *time.Time {
	if t != (time.Time{}) {
		return &t
	}
	return nil
}

func (m *PrimitiveMapper) UintToInt32(u uint) int32 {
	return int32(u)
}

func (m *PrimitiveMapper) Int32ToUint(i int32) *uint {
	if i != 0 {
		u := uint(i)
		return &u
	}
	return nil
}
