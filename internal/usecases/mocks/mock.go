// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock_usecases is a generated GoMock package.
package mock_usecases

import (
	models "HeadHunter/internal/entity/models"
	multipart "mime/multipart"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// AuthCheck mocks base method.
func (m *MockUser) AuthCheck(email string) (*models.UserAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AuthCheck", email)
	ret0, _ := ret[0].(*models.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AuthCheck indicates an expected call of AuthCheck.
func (mr *MockUserMockRecorder) AuthCheck(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AuthCheck", reflect.TypeOf((*MockUser)(nil).AuthCheck), email)
}

// DeleteUserImage mocks base method.
func (m *MockUser) DeleteUserImage(user *models.UserAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserImage", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserImage indicates an expected call of DeleteUserImage.
func (mr *MockUserMockRecorder) DeleteUserImage(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserImage", reflect.TypeOf((*MockUser)(nil).DeleteUserImage), user)
}

// GetUser mocks base method.
func (m *MockUser) GetUser(id uint) (*models.UserAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", id)
	ret0, _ := ret[0].(*models.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser.
func (mr *MockUserMockRecorder) GetUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUser)(nil).GetUser), id)
}

// GetUserByEmail mocks base method.
func (m *MockUser) GetUserByEmail(email string) (*models.UserAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByEmail", email)
	ret0, _ := ret[0].(*models.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByEmail indicates an expected call of GetUserByEmail.
func (mr *MockUserMockRecorder) GetUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByEmail", reflect.TypeOf((*MockUser)(nil).GetUserByEmail), email)
}

// GetUserId mocks base method.
func (m *MockUser) GetUserId(email string) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserId", email)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserId indicates an expected call of GetUserId.
func (mr *MockUserMockRecorder) GetUserId(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserId", reflect.TypeOf((*MockUser)(nil).GetUserId), email)
}

// GetUserSafety mocks base method.
func (m *MockUser) GetUserSafety(id uint) (*models.UserAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserSafety", id)
	ret0, _ := ret[0].(*models.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserSafety indicates an expected call of GetUserSafety.
func (mr *MockUserMockRecorder) GetUserSafety(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserSafety", reflect.TypeOf((*MockUser)(nil).GetUserSafety), id)
}

// Logout mocks base method.
func (m *MockUser) Logout(token string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Logout", token)
	ret0, _ := ret[0].(error)
	return ret0
}

// Logout indicates an expected call of Logout.
func (mr *MockUserMockRecorder) Logout(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockUser)(nil).Logout), token)
}

// SignIn mocks base method.
func (m *MockUser) SignIn(input *models.UserAccount) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignIn", input)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignIn indicates an expected call of SignIn.
func (mr *MockUserMockRecorder) SignIn(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignIn", reflect.TypeOf((*MockUser)(nil).SignIn), input)
}

// SignUp mocks base method.
func (m *MockUser) SignUp(input *models.UserAccount) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignUp", input)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignUp indicates an expected call of SignUp.
func (mr *MockUserMockRecorder) SignUp(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignUp", reflect.TypeOf((*MockUser)(nil).SignUp), input)
}

// UpdateUser mocks base method.
func (m *MockUser) UpdateUser(input *models.UserAccount) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", input)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser.
func (mr *MockUserMockRecorder) UpdateUser(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUser)(nil).UpdateUser), input)
}

// UploadUserImage mocks base method.
func (m *MockUser) UploadUserImage(user *models.UserAccount, fileHeader *multipart.FileHeader) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UploadUserImage", user, fileHeader)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UploadUserImage indicates an expected call of UploadUserImage.
func (mr *MockUserMockRecorder) UploadUserImage(user, fileHeader interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UploadUserImage", reflect.TypeOf((*MockUser)(nil).UploadUserImage), user, fileHeader)
}

// MockVacancy is a mock of Vacancy interface.
type MockVacancy struct {
	ctrl     *gomock.Controller
	recorder *MockVacancyMockRecorder
}

// MockVacancyMockRecorder is the mock recorder for MockVacancy.
type MockVacancyMockRecorder struct {
	mock *MockVacancy
}

// NewMockVacancy creates a new mock instance.
func NewMockVacancy(ctrl *gomock.Controller) *MockVacancy {
	mock := &MockVacancy{ctrl: ctrl}
	mock.recorder = &MockVacancyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVacancy) EXPECT() *MockVacancyMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockVacancy) Create(arg0 uint, arg1 *models.Vacancy) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0, arg1)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockVacancyMockRecorder) Create(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVacancy)(nil).Create), arg0, arg1)
}

// Delete mocks base method.
func (m *MockVacancy) Delete(arg0 uint, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockVacancyMockRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVacancy)(nil).Delete), arg0, arg1)
}

// GetAll mocks base method.
func (m *MockVacancy) GetAll() ([]*models.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*models.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockVacancyMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockVacancy)(nil).GetAll))
}

// GetById mocks base method.
func (m *MockVacancy) GetById(arg0 int) (*models.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0)
	ret0, _ := ret[0].(*models.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockVacancyMockRecorder) GetById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockVacancy)(nil).GetById), arg0)
}

// GetByUserId mocks base method.
func (m *MockVacancy) GetByUserId(arg0 int) ([]*models.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserId", arg0)
	ret0, _ := ret[0].([]*models.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserId indicates an expected call of GetByUserId.
func (mr *MockVacancyMockRecorder) GetByUserId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserId", reflect.TypeOf((*MockVacancy)(nil).GetByUserId), arg0)
}

// Update mocks base method.
func (m *MockVacancy) Update(arg0 uint, arg1 int, arg2 *models.Vacancy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockVacancyMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVacancy)(nil).Update), arg0, arg1, arg2)
}

// MockVacancyActivity is a mock of VacancyActivity interface.
type MockVacancyActivity struct {
	ctrl     *gomock.Controller
	recorder *MockVacancyActivityMockRecorder
}

// MockVacancyActivityMockRecorder is the mock recorder for MockVacancyActivity.
type MockVacancyActivityMockRecorder struct {
	mock *MockVacancyActivity
}

// NewMockVacancyActivity creates a new mock instance.
func NewMockVacancyActivity(ctrl *gomock.Controller) *MockVacancyActivity {
	mock := &MockVacancyActivity{ctrl: ctrl}
	mock.recorder = &MockVacancyActivityMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVacancyActivity) EXPECT() *MockVacancyActivityMockRecorder {
	return m.recorder
}

// ApplyForVacancy mocks base method.
func (m *MockVacancyActivity) ApplyForVacancy(arg0 uint, arg1 int, arg2 *models.VacancyActivity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyForVacancy", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyForVacancy indicates an expected call of ApplyForVacancy.
func (mr *MockVacancyActivityMockRecorder) ApplyForVacancy(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyForVacancy", reflect.TypeOf((*MockVacancyActivity)(nil).ApplyForVacancy), arg0, arg1, arg2)
}

// GetAllUserApplies mocks base method.
func (m *MockVacancyActivity) GetAllUserApplies(arg0 int) ([]*models.VacancyActivity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserApplies", arg0)
	ret0, _ := ret[0].([]*models.VacancyActivity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserApplies indicates an expected call of GetAllUserApplies.
func (mr *MockVacancyActivityMockRecorder) GetAllUserApplies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserApplies", reflect.TypeOf((*MockVacancyActivity)(nil).GetAllUserApplies), arg0)
}

// GetAllVacancyApplies mocks base method.
func (m *MockVacancyActivity) GetAllVacancyApplies(arg0 int) ([]*models.VacancyActivity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllVacancyApplies", arg0)
	ret0, _ := ret[0].([]*models.VacancyActivity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllVacancyApplies indicates an expected call of GetAllVacancyApplies.
func (mr *MockVacancyActivityMockRecorder) GetAllVacancyApplies(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllVacancyApplies", reflect.TypeOf((*MockVacancyActivity)(nil).GetAllVacancyApplies), arg0)
}

// MockResume is a mock of Resume interface.
type MockResume struct {
	ctrl     *gomock.Controller
	recorder *MockResumeMockRecorder
}

// MockResumeMockRecorder is the mock recorder for MockResume.
type MockResumeMockRecorder struct {
	mock *MockResume
}

// NewMockResume creates a new mock instance.
func NewMockResume(ctrl *gomock.Controller) *MockResume {
	mock := &MockResume{ctrl: ctrl}
	mock.recorder = &MockResumeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResume) EXPECT() *MockResumeMockRecorder {
	return m.recorder
}

// CreateResume mocks base method.
func (m *MockResume) CreateResume(resume *models.Resume, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateResume", resume, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateResume indicates an expected call of CreateResume.
func (mr *MockResumeMockRecorder) CreateResume(resume, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateResume", reflect.TypeOf((*MockResume)(nil).CreateResume), resume, email)
}

// DeleteResume mocks base method.
func (m *MockResume) DeleteResume(id uint, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteResume", id, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteResume indicates an expected call of DeleteResume.
func (mr *MockResumeMockRecorder) DeleteResume(id, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteResume", reflect.TypeOf((*MockResume)(nil).DeleteResume), id, email)
}

// GetPreviewResumeByApplicant mocks base method.
func (m *MockResume) GetPreviewResumeByApplicant(userId uint, email string) ([]*models.ResumePreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPreviewResumeByApplicant", userId, email)
	ret0, _ := ret[0].([]*models.ResumePreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPreviewResumeByApplicant indicates an expected call of GetPreviewResumeByApplicant.
func (mr *MockResumeMockRecorder) GetPreviewResumeByApplicant(userId, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPreviewResumeByApplicant", reflect.TypeOf((*MockResume)(nil).GetPreviewResumeByApplicant), userId, email)
}

// GetResume mocks base method.
func (m *MockResume) GetResume(id uint, email string) (*models.Resume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResume", id)
	ret0, _ := ret[0].(*models.Resume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResume indicates an expected call of GetResume.
func (mr *MockResumeMockRecorder) GetResume(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResume", reflect.TypeOf((*MockResume)(nil).GetResume), id)
}

// GetResumeByApplicant mocks base method.
func (m *MockResume) GetResumeByApplicant(userId uint, email string) ([]*models.Resume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResumeByApplicant", userId, email)
	ret0, _ := ret[0].([]*models.Resume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResumeByApplicant indicates an expected call of GetResumeByApplicant.
func (mr *MockResumeMockRecorder) GetResumeByApplicant(userId, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResumeByApplicant", reflect.TypeOf((*MockResume)(nil).GetResumeByApplicant), userId, email)
}

// UpdateResume mocks base method.
func (m *MockResume) UpdateResume(id uint, resume *models.Resume, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateResume", id, resume, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateResume indicates an expected call of UpdateResume.
func (mr *MockResumeMockRecorder) UpdateResume(id, resume, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateResume", reflect.TypeOf((*MockResume)(nil).UpdateResume), id, resume, email)
}
