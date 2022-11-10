// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	models "HeadHunter/internal/entity/models"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUserRepository) CreateUser(user *models.UserAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUserRepositoryMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserRepository)(nil).CreateUser), user)
}

// GetUser mocks base method.
func (m *MockUserRepository) GetUser(id uint) (*models.UserAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", id)
	ret0, _ := ret[0].(*models.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserRepositoryMockRecorder) GetUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserRepository)(nil).GetUser), id)
}

// GetUserByEmail mocks base method.
func (m *MockUserRepository) GetUserByEmail(email string) (*models.UserAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", email)
	ret0, _ := ret[0].(*models.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserRepositoryMockRecorder) GetUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUserRepository)(nil).GetUserByEmail), email)
}

// GetUserSafety mocks base method.
func (m *MockUserRepository) GetUserSafety(id uint, safeFields []string) (*models.UserAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSafety", id, safeFields)
	ret0, _ := ret[0].(*models.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSafety indicates an expected call of GetUserSafety.
func (mr *MockUserRepositoryMockRecorder) GetUserSafety(id, safeFields interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSafety", reflect.TypeOf((*MockUserRepository)(nil).GetUserSafety), id, safeFields)
}

// IsUserExist mocks base method.
func (m *MockUserRepository) IsUserExist(email string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsUserExist", email)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsUserExist indicates an expected call of IsUserExist.
func (mr *MockUserRepositoryMockRecorder) IsUserExist(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsUserExist", reflect.TypeOf((*MockUserRepository)(nil).IsUserExist), email)
}

// UpdateUser mocks base method.
func (m *MockUserRepository) UpdateUser(oldUser, newUser *models.UserAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", oldUser, newUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserRepositoryMockRecorder) UpdateUser(oldUser, newUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepository)(nil).UpdateUser), oldUser, newUser)
}

// UpdateUserField mocks base method.
func (m *MockUserRepository) UpdateUserField(oldUser, newUser *models.UserAccount, field ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{oldUser, newUser}
	for _, a := range field {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateUserField", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUserField indicates an expected call of UpdateUserField.
func (mr *MockUserRepositoryMockRecorder) UpdateUserField(oldUser, newUser interface{}, field ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{oldUser, newUser}, field...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUserField", reflect.TypeOf((*MockUserRepository)(nil).UpdateUserField), varargs...)
}

// MockVacancyRepository is a mock of VacancyRepository interface.
type MockVacancyRepository struct {
	ctrl     *gomock.Controller
	recorder *MockVacancyRepositoryMockRecorder
}

// MockVacancyRepositoryMockRecorder is the mock recorder for MockVacancyRepository.
type MockVacancyRepositoryMockRecorder struct {
	mock *MockVacancyRepository
}

// NewMockVacancyRepository creates a new mock instance.
func NewMockVacancyRepository(ctrl *gomock.Controller) *MockVacancyRepository {
	mock := &MockVacancyRepository{ctrl: ctrl}
	mock.recorder = &MockVacancyRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVacancyRepository) EXPECT() *MockVacancyRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockVacancyRepository) Create(vacancy *models.Vacancy) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", vacancy)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockVacancyRepositoryMockRecorder) Create(vacancy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVacancyRepository)(nil).Create), vacancy)
}

// Delete mocks base method.
func (m *MockVacancyRepository) Delete(arg0 uint, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockVacancyRepositoryMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVacancyRepository)(nil).Delete), arg0, arg1)
}

// GetAll mocks base method.
func (m *MockVacancyRepository) GetAll() ([]*models.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*models.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockVacancyRepositoryMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockVacancyRepository)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockVacancyRepository) GetById(arg0 int) (*models.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(*models.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockVacancyRepositoryMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockVacancyRepository)(nil).GetById), arg0)
}

// GetByUserId mocks base method.
func (m *MockVacancyRepository) GetByUserId(arg0 int) ([]*models.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserId", arg0)
	ret0, _ := ret[0].([]*models.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserId indicates an expected call of GetByUserId.
func (mr *MockVacancyRepositoryMockRecorder) GetByUserId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserId", reflect.TypeOf((*MockVacancyRepository)(nil).GetByUserId), arg0)
}

// Update mocks base method.
func (m *MockVacancyRepository) Update(arg0 int, arg1 *models.Vacancy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockVacancyRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVacancyRepository)(nil).Update), arg0, arg1)
}

// MockVacancyActivityRepository is a mock of VacancyActivityRepository interface.
type MockVacancyActivityRepository struct {
	ctrl     *gomock.Controller
	recorder *MockVacancyActivityRepositoryMockRecorder
}

// MockVacancyActivityRepositoryMockRecorder is the mock recorder for MockVacancyActivityRepository.
type MockVacancyActivityRepositoryMockRecorder struct {
	mock *MockVacancyActivityRepository
}

// NewMockVacancyActivityRepository creates a new mock instance.
func NewMockVacancyActivityRepository(ctrl *gomock.Controller) *MockVacancyActivityRepository {
	mock := &MockVacancyActivityRepository{ctrl: ctrl}
	mock.recorder = &MockVacancyActivityRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVacancyActivityRepository) EXPECT() *MockVacancyActivityRepositoryMockRecorder {
	return m.recorder
}

// ApplyForVacancy mocks base method.
func (m *MockVacancyActivityRepository) ApplyForVacancy(arg0 *models.VacancyActivity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyForVacancy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyForVacancy indicates an expected call of ApplyForVacancy.
func (mr *MockVacancyActivityRepositoryMockRecorder) ApplyForVacancy(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyForVacancy", reflect.TypeOf((*MockVacancyActivityRepository)(nil).ApplyForVacancy), arg0)
}

// GetAllUserApplies mocks base method.
func (m *MockVacancyActivityRepository) GetAllUserApplies(arg0 int) ([]models.VacancyActivity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserApplies", arg0)
	ret0, _ := ret[0].([]models.VacancyActivity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserApplies indicates an expected call of GetAllUserApplies.
func (mr *MockVacancyActivityRepositoryMockRecorder) GetAllUserApplies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserApplies", reflect.TypeOf((*MockVacancyActivityRepository)(nil).GetAllUserApplies), arg0)
}

// GetAllVacancyApplies mocks base method.
func (m *MockVacancyActivityRepository) GetAllVacancyApplies(arg0 int) ([]*models.VacancyActivity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllVacancyApplies", arg0)
	ret0, _ := ret[0].([]*models.VacancyActivity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllVacancyApplies indicates an expected call of GetAllVacancyApplies.
func (mr *MockVacancyActivityRepositoryMockRecorder) GetAllVacancyApplies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllVacancyApplies", reflect.TypeOf((*MockVacancyActivityRepository)(nil).GetAllVacancyApplies), arg0)
}

// MockResumeRepository is a mock of ResumeRepository interface.
type MockResumeRepository struct {
	ctrl     *gomock.Controller
	recorder *MockResumeRepositoryMockRecorder
}

// MockResumeRepositoryMockRecorder is the mock recorder for MockResumeRepository.
type MockResumeRepositoryMockRecorder struct {
	mock *MockResumeRepository
}

// NewMockResumeRepository creates a new mock instance.
func NewMockResumeRepository(ctrl *gomock.Controller) *MockResumeRepository {
	mock := &MockResumeRepository{ctrl: ctrl}
	mock.recorder = &MockResumeRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResumeRepository) EXPECT() *MockResumeRepositoryMockRecorder {
	return m.recorder
}

// CreateResume mocks base method.
func (m *MockResumeRepository) CreateResume(resume *models.Resume, userId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateResume", resume, userId)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateResume indicates an expected call of CreateResume.
func (mr *MockResumeRepositoryMockRecorder) CreateResume(resume, userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateResume", reflect.TypeOf((*MockResumeRepository)(nil).CreateResume), resume, userId)
}

// DeleteResume mocks base method.
func (m *MockResumeRepository) DeleteResume(id uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResume", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteResume indicates an expected call of DeleteResume.
func (mr *MockResumeRepositoryMockRecorder) DeleteResume(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResume", reflect.TypeOf((*MockResumeRepository)(nil).DeleteResume), id)
}

// GetPreviewResumeByApplicant mocks base method.
func (m *MockResumeRepository) GetPreviewResumeByApplicant(userId uint) ([]*models.Resume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPreviewResumeByApplicant", userId)
	ret0, _ := ret[0].([]*models.Resume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPreviewResumeByApplicant indicates an expected call of GetPreviewResumeByApplicant.
func (mr *MockResumeRepositoryMockRecorder) GetPreviewResumeByApplicant(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPreviewResumeByApplicant", reflect.TypeOf((*MockResumeRepository)(nil).GetPreviewResumeByApplicant), userId)
}

// GetResume mocks base method.
func (m *MockResumeRepository) GetResume(id uint) (*models.Resume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResume", id)
	ret0, _ := ret[0].(*models.Resume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResume indicates an expected call of GetResume.
func (mr *MockResumeRepositoryMockRecorder) GetResume(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResume", reflect.TypeOf((*MockResumeRepository)(nil).GetResume), id)
}

// GetResumeByApplicant mocks base method.
func (m *MockResumeRepository) GetResumeByApplicant(userId uint) ([]*models.Resume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResumeByApplicant", userId)
	ret0, _ := ret[0].([]*models.Resume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResumeByApplicant indicates an expected call of GetResumeByApplicant.
func (mr *MockResumeRepositoryMockRecorder) GetResumeByApplicant(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResumeByApplicant", reflect.TypeOf((*MockResumeRepository)(nil).GetResumeByApplicant), userId)
}

// UpdateResume mocks base method.
func (m *MockResumeRepository) UpdateResume(id uint, resume *models.Resume) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateResume", id, resume)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateResume indicates an expected call of UpdateResume.
func (mr *MockResumeRepositoryMockRecorder) UpdateResume(id, resume interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResume", reflect.TypeOf((*MockResumeRepository)(nil).UpdateResume), id, resume)
}