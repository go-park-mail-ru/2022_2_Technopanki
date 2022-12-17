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

// GetAllUsers mocks base method.
func (m *MockUserRepository) GetAllUsers(conditions []string, filterValues []interface{}, flag string) ([]*models.UserAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers", conditions, filterValues, flag)
	ret0, _ := ret[0].([]*models.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockUserRepositoryMockRecorder) GetAllUsers(conditions, filterValues, flag interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserRepository)(nil).GetAllUsers), conditions, filterValues, flag)
}

// GetBestApplicantForEmployer mocks base method.
func (m *MockUserRepository) GetBestApplicantForEmployer() ([]*models.UserAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBestApplicantForEmployer")
	ret0, _ := ret[0].([]*models.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBestApplicantForEmployer indicates an expected call of GetBestApplicantForEmployer.
func (mr *MockUserRepositoryMockRecorder) GetBestApplicantForEmployer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBestApplicantForEmployer", reflect.TypeOf((*MockUserRepository)(nil).GetBestApplicantForEmployer))
}

// GetBestVacanciesForApplicant mocks base method.
func (m *MockUserRepository) GetBestVacanciesForApplicant() ([]*models.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBestVacanciesForApplicant")
	ret0, _ := ret[0].([]*models.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBestVacanciesForApplicant indicates an expected call of GetBestVacanciesForApplicant.
func (mr *MockUserRepositoryMockRecorder) GetBestVacanciesForApplicant() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBestVacanciesForApplicant", reflect.TypeOf((*MockUserRepository)(nil).GetBestVacanciesForApplicant))
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

// UpdateUser mocks base method.
func (m *MockUserRepository) UpdateUser(newUser *models.UserAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", newUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserRepositoryMockRecorder) UpdateUser(newUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserRepository)(nil).UpdateUser), newUser)
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

// AddVacancyToFavorites mocks base method.
func (m *MockVacancyRepository) AddVacancyToFavorites(user *models.UserAccount, vacancy *models.Vacancy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddVacancyToFavorites", user, vacancy)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddVacancyToFavorites indicates an expected call of AddVacancyToFavorites.
func (mr *MockVacancyRepositoryMockRecorder) AddVacancyToFavorites(user, vacancy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddVacancyToFavorites", reflect.TypeOf((*MockVacancyRepository)(nil).AddVacancyToFavorites), user, vacancy)
}

// CheckFavoriteVacancy mocks base method.
func (m *MockVacancyRepository) CheckFavoriteVacancy(userId, vacancyId uint) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckFavoriteVacancy", userId, vacancyId)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckFavoriteVacancy indicates an expected call of CheckFavoriteVacancy.
func (mr *MockVacancyRepositoryMockRecorder) CheckFavoriteVacancy(userId, vacancyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckFavoriteVacancy", reflect.TypeOf((*MockVacancyRepository)(nil).CheckFavoriteVacancy), userId, vacancyId)
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
func (m *MockVacancyRepository) Delete(userId, vacancyId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userId, vacancyId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockVacancyRepositoryMockRecorder) Delete(userId, vacancyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVacancyRepository)(nil).Delete), userId, vacancyId)
}

// DeleteVacancyFromFavorites mocks base method.
func (m *MockVacancyRepository) DeleteVacancyFromFavorites(user *models.UserAccount, vacancy *models.Vacancy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVacancyFromFavorites", user, vacancy)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteVacancyFromFavorites indicates an expected call of DeleteVacancyFromFavorites.
func (mr *MockVacancyRepositoryMockRecorder) DeleteVacancyFromFavorites(user, vacancy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVacancyFromFavorites", reflect.TypeOf((*MockVacancyRepository)(nil).DeleteVacancyFromFavorites), user, vacancy)
}

// GetAll mocks base method.
func (m *MockVacancyRepository) GetAll(conditions []string, filterValues []interface{}) ([]*models.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", conditions, filterValues)
	ret0, _ := ret[0].([]*models.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockVacancyRepositoryMockRecorder) GetAll(conditions, filterValues interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockVacancyRepository)(nil).GetAll), conditions, filterValues)
}

// GetAllFilter mocks base method.
func (m *MockVacancyRepository) GetAllFilter(filter string) ([]*models.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFilter", filter)
	ret0, _ := ret[0].([]*models.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFilter indicates an expected call of GetAllFilter.
func (mr *MockVacancyRepositoryMockRecorder) GetAllFilter(filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFilter", reflect.TypeOf((*MockVacancyRepository)(nil).GetAllFilter), filter)
}

// GetById mocks base method.
func (m *MockVacancyRepository) GetById(vacancyId uint) (*models.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", vacancyId)
	ret0, _ := ret[0].(*models.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockVacancyRepositoryMockRecorder) GetById(vacancyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockVacancyRepository)(nil).GetById), vacancyId)
}

// GetByUserId mocks base method.
func (m *MockVacancyRepository) GetByUserId(userId uint) ([]*models.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserId", userId)
	ret0, _ := ret[0].([]*models.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserId indicates an expected call of GetByUserId.
func (mr *MockVacancyRepositoryMockRecorder) GetByUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserId", reflect.TypeOf((*MockVacancyRepository)(nil).GetByUserId), userId)
}

// GetPreviewVacanciesByEmployer mocks base method.
func (m *MockVacancyRepository) GetPreviewVacanciesByEmployer(userId uint) ([]*models.VacancyPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPreviewVacanciesByEmployer", userId)
	ret0, _ := ret[0].([]*models.VacancyPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPreviewVacanciesByEmployer indicates an expected call of GetPreviewVacanciesByEmployer.
func (mr *MockVacancyRepositoryMockRecorder) GetPreviewVacanciesByEmployer(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPreviewVacanciesByEmployer", reflect.TypeOf((*MockVacancyRepository)(nil).GetPreviewVacanciesByEmployer), userId)
}

// GetUserFavoriteVacancies mocks base method.
func (m *MockVacancyRepository) GetUserFavoriteVacancies(user *models.UserAccount) ([]*models.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserFavoriteVacancies", user)
	ret0, _ := ret[0].([]*models.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserFavoriteVacancies indicates an expected call of GetUserFavoriteVacancies.
func (mr *MockVacancyRepositoryMockRecorder) GetUserFavoriteVacancies(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserFavoriteVacancies", reflect.TypeOf((*MockVacancyRepository)(nil).GetUserFavoriteVacancies), user)
}

// Update mocks base method.
func (m *MockVacancyRepository) Update(userId, vacancyId uint, oldVacancy, updates *models.Vacancy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userId, vacancyId, oldVacancy, updates)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockVacancyRepositoryMockRecorder) Update(userId, vacancyId, oldVacancy, updates interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVacancyRepository)(nil).Update), userId, vacancyId, oldVacancy, updates)
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

// DeleteUserApply mocks base method.
func (m *MockVacancyActivityRepository) DeleteUserApply(userId, applyId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserApply", userId, applyId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserApply indicates an expected call of DeleteUserApply.
func (mr *MockVacancyActivityRepositoryMockRecorder) DeleteUserApply(userId, applyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserApply", reflect.TypeOf((*MockVacancyActivityRepository)(nil).DeleteUserApply), userId, applyId)
}

// GetAllUserApplies mocks base method.
func (m *MockVacancyActivityRepository) GetAllUserApplies(userId uint) ([]*models.VacancyActivityPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserApplies", userId)
	ret0, _ := ret[0].([]*models.VacancyActivityPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserApplies indicates an expected call of GetAllUserApplies.
func (mr *MockVacancyActivityRepositoryMockRecorder) GetAllUserApplies(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserApplies", reflect.TypeOf((*MockVacancyActivityRepository)(nil).GetAllUserApplies), userId)
}

// GetAllVacancyApplies mocks base method.
func (m *MockVacancyActivityRepository) GetAllVacancyApplies(vacancyId uint) ([]*models.VacancyActivityPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllVacancyApplies", vacancyId)
	ret0, _ := ret[0].([]*models.VacancyActivityPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllVacancyApplies indicates an expected call of GetAllVacancyApplies.
func (mr *MockVacancyActivityRepositoryMockRecorder) GetAllVacancyApplies(vacancyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllVacancyApplies", reflect.TypeOf((*MockVacancyActivityRepository)(nil).GetAllVacancyApplies), vacancyId)
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

// GetAllResumes mocks base method.
func (m *MockResumeRepository) GetAllResumes(conditions []string, filterValues []interface{}) ([]*models.Resume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllResumes", conditions, filterValues)
	ret0, _ := ret[0].([]*models.Resume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllResumes indicates an expected call of GetAllResumes.
func (mr *MockResumeRepositoryMockRecorder) GetAllResumes(conditions, filterValues interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllResumes", reflect.TypeOf((*MockResumeRepository)(nil).GetAllResumes), conditions, filterValues)
}

// GetEmployerIdByVacancyActivity mocks base method.
func (m *MockResumeRepository) GetEmployerIdByVacancyActivity(id uint) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmployerIdByVacancyActivity", id)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEmployerIdByVacancyActivity indicates an expected call of GetEmployerIdByVacancyActivity.
func (mr *MockResumeRepositoryMockRecorder) GetEmployerIdByVacancyActivity(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmployerIdByVacancyActivity", reflect.TypeOf((*MockResumeRepository)(nil).GetEmployerIdByVacancyActivity), id)
}

// GetPreviewResumeByApplicant mocks base method.
func (m *MockResumeRepository) GetPreviewResumeByApplicant(userId uint) ([]*models.ResumePreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPreviewResumeByApplicant", userId)
	ret0, _ := ret[0].([]*models.ResumePreview)
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
