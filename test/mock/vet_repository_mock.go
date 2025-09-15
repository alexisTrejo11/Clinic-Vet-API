package mock

import (
	context "context"
	reflect "reflect"

	vetDtos "clinic-vet-api/app/veterinarians/application/dtos"
	vetDomain "clinic-vet-api/app/veterinarians/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockVeterinarianRepository is a mock of VeterinarianRepository interface.
type MockVeterinarianRepository struct {
	ctrl     *gomock.Controller
	recorder *MockVeterinarianRepositoryMockRecorder
}

// MockVeterinarianRepositoryMockRecorder is the mock recorder for MockVeterinarianRepository.
type MockVeterinarianRepositoryMockRecorder struct {
	mock *MockVeterinarianRepository
}

// NewMockVeterinarianRepository creates a new mock instance.
func NewMockVeterinarianRepository(ctrl *gomock.Controller) *MockVeterinarianRepository {
	mock := &MockVeterinarianRepository{ctrl: ctrl}
	mock.recorder = &MockVeterinarianRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVeterinarianRepository) EXPECT() *MockVeterinarianRepositoryMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *MockVeterinarianRepository) Delete(ctx context.Context, id int, isSoftDelete bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, id, isSoftDelete)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockVeterinarianRepositoryMockRecorder) Delete(ctx context.Context, id int) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVeterinarianRepository)(nil).Delete), ctx, id)
}

// Exists mocks base method.
func (m *MockVeterinarianRepository) Exists(ctx context.Context, vetId int) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", ctx, vetId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Exists indicates an expected call of Exists.
func (mr *MockVeterinarianRepositoryMockRecorder) Exists(ctx, vetId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockVeterinarianRepository)(nil).Exists), ctx, vetId)
}

// GetByID mocks base method.
func (m *MockVeterinarianRepository) GetByID(ctx context.Context, id int) (vetDomain.Veterinarian, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(vetDomain.Veterinarian)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockVeterinarianRepositoryMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockVeterinarianRepository)(nil).GetByID), ctx, id)
}

// GetByUserID mocks base method.
func (m *MockVeterinarianRepository) GetByUserID(ctx context.Context, id int) (vetDomain.Veterinarian, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserID", ctx, id)
	ret0, _ := ret[0].(vetDomain.Veterinarian)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserID indicates an expected call of GetByUserID.
func (mr *MockVeterinarianRepositoryMockRecorder) GetByUserID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserID", reflect.TypeOf((*MockVeterinarianRepository)(nil).GetByUserID), ctx, id)
}

// List mocks base method.
func (m *MockVeterinarianRepository) List(ctx context.Context, searchParams vetDtos.VetSearchParams) ([]vetDomain.Veterinarian, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", ctx, searchParams)
	ret0, _ := ret[0].([]vetDomain.Veterinarian)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockVeterinarianRepositoryMockRecorder) List(ctx, searchParams interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockVeterinarianRepository)(nil).List), ctx, searchParams)
}

// Save mocks base method.
func (m *MockVeterinarianRepository) Save(ctx context.Context, pet *vetDomain.Veterinarian) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, pet)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockVeterinarianRepositoryMockRecorder) Save(ctx, pet interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockVeterinarianRepository)(nil).Save), ctx, pet)
}
