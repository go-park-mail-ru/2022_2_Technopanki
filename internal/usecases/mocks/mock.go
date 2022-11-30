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

// ConfirmUser mocks base method.
func (m *MockUser) ConfirmUser(code, email string) (*models.UserAccount, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConfirmUser", code, email)
	ret0, _ := ret[0].(*models.UserAccount)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ConfirmUser indicates an expected call of ConfirmUser.
func (mr *MockUserMockRecorder) ConfirmUser(code, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConfirmUser", reflect.TypeOf((*MockUser)(nil).ConfirmUser), code, email)
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

// GetAllApplicants mocks base method.
func (m *MockUser) GetAllApplicants(filters models.UserFilter) ([]*models.UserAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllApplicants", filters)
	ret0, _ := ret[0].([]*models.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllApplicants indicates an expected call of GetAllApplicants.
func (mr *MockUserMockRecorder) GetAllApplicants(filters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllApplicants", reflect.TypeOf((*MockUser)(nil).GetAllApplicants), filters)
}

// GetAllEmployers mocks base method.
func (m *MockUser) GetAllEmployers(filters models.UserFilter) ([]*models.UserAccount, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllEmployers", filters)
	ret0, _ := ret[0].([]*models.UserAccount)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllEmployers indicates an expected call of GetAllEmployers.
func (mr *MockUserMockRecorder) GetAllEmployers(filters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllEmployers", reflect.TypeOf((*MockUser)(nil).GetAllEmployers), filters)
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

// UpdatePassword mocks base method.
func (m *MockUser) UpdatePassword(code, email, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePassword", code, email, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatePassword indicates an expected call of UpdatePassword.
func (mr *MockUserMockRecorder) UpdatePassword(code, email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePassword", reflect.TypeOf((*MockUser)(nil).UpdatePassword), code, email, password)
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
func (m *MockVacancy) Create(email string, input *models.Vacancy) (uint, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", email, input)
	ret0, _ := ret[0].(uint)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockVacancyMockRecorder) Create(email, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockVacancy)(nil).Create), email, input)
}

// Delete mocks base method.
func (m *MockVacancy) Delete(email string, vacancyId uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", email, vacancyId)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockVacancyMockRecorder) Delete(email, vacancyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockVacancy)(nil).Delete), email, vacancyId)
}

// GetAll mocks base method.
func (m *MockVacancy) GetAll(filters models.VacancyFilter) ([]*models.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", filters)
	ret0, _ := ret[0].([]*models.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockVacancyMockRecorder) GetAll(filters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockVacancy)(nil).GetAll), filters)
}

// GetById mocks base method.
func (m *MockVacancy) GetById(vacancyId uint) (*models.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", vacancyId)
	ret0, _ := ret[0].(*models.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockVacancyMockRecorder) GetById(vacancyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockVacancy)(nil).GetById), vacancyId)
}

// GetByUserId mocks base method.
func (m *MockVacancy) GetByUserId(userId uint) ([]*models.Vacancy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserId", userId)
	ret0, _ := ret[0].([]*models.Vacancy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserId indicates an expected call of GetByUserId.
func (mr *MockVacancyMockRecorder) GetByUserId(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserId", reflect.TypeOf((*MockVacancy)(nil).GetByUserId), userId)
}

// GetPreviewVacanciesByEmployer mocks base method.
func (m *MockVacancy) GetPreviewVacanciesByEmployer(userId uint) ([]*models.VacancyPreview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPreviewVacanciesByEmployer", userId)
	ret0, _ := ret[0].([]*models.VacancyPreview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPreviewVacanciesByEmployer indicates an expected call of GetPreviewVacanciesByEmployer.
func (mr *MockVacancyMockRecorder) GetPreviewVacanciesByEmployer(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPreviewVacanciesByEmployer", reflect.TypeOf((*MockVacancy)(nil).GetPreviewVacanciesByEmployer), userId)
}

// Update mocks base method.
func (m *MockVacancy) Update(email string, vacancyId uint, updates *models.Vacancy) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", email, vacancyId, updates)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockVacancyMockRecorder) Update(email, vacancyId, updates interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockVacancy)(nil).Update), email, vacancyId, updates)
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
func (m *MockVacancyActivity) ApplyForVacancy(email string, vacancyId uint, input *models.VacancyActivity) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyForVacancy", email, vacancyId, input)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyForVacancy indicates an expected call of ApplyForVacancy.
func (mr *MockVacancyActivityMockRecorder) ApplyForVacancy(email, vacancyId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyForVacancy", reflect.TypeOf((*MockVacancyActivity)(nil).ApplyForVacancy), email, vacancyId, input)
}

// DeleteUserApply mocks base method.
func (m *MockVacancyActivity) DeleteUserApply(email string, apply uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUserApply", email, apply)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUserApply indicates an expected call of DeleteUserApply.
func (mr *MockVacancyActivityMockRecorder) DeleteUserApply(email, apply interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUserApply", reflect.TypeOf((*MockVacancyActivity)(nil).DeleteUserApply), email, apply)
}

// GetAllUserApplies mocks base method.
func (m *MockVacancyActivity) GetAllUserApplies(userid uint) ([]*models.VacancyActivity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUserApplies", userid)
	ret0, _ := ret[0].([]*models.VacancyActivity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUserApplies indicates an expected call of GetAllUserApplies.
func (mr *MockVacancyActivityMockRecorder) GetAllUserApplies(userid interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUserApplies", reflect.TypeOf((*MockVacancyActivity)(nil).GetAllUserApplies), userid)
}

// GetAllVacancyApplies mocks base method.
func (m *MockVacancyActivity) GetAllVacancyApplies(vacancyId uint) ([]*models.VacancyActivity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllVacancyApplies", vacancyId)
	ret0, _ := ret[0].([]*models.VacancyActivity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllVacancyApplies indicates an expected call of GetAllVacancyApplies.
func (mr *MockVacancyActivityMockRecorder) GetAllVacancyApplies(vacancyId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllVacancyApplies", reflect.TypeOf((*MockVacancyActivity)(nil).GetAllVacancyApplies), vacancyId)
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

// GetAllResumes mocks base method.
func (m *MockResume) GetAllResumes(filters models.ResumeFilter) ([]*models.Resume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllResumes", filters)
	ret0, _ := ret[0].([]*models.Resume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllResumes indicates an expected call of GetAllResumes.
func (mr *MockResumeMockRecorder) GetAllResumes(filters interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllResumes", reflect.TypeOf((*MockResume)(nil).GetAllResumes), filters)
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
	ret := m.ctrl.Call(m, "GetResume", id, email)
	ret0, _ := ret[0].(*models.Resume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResume indicates an expected call of GetResume.
func (mr *MockResumeMockRecorder) GetResume(id, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResume", reflect.TypeOf((*MockResume)(nil).GetResume), id, email)
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
