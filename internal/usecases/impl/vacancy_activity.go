package impl

import (
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/repository"
	"HeadHunter/pkg/errorHandler"
)

type VacancyActivityService struct {
	vacancyActivityRep repository.VacancyActivityRepository
	userRep            repository.UserRepository
	vacancyRep         repository.VacancyRepository
	notificationRep    repository.NotificationRepository
}

func NewVacancyActivityService(vacancyActivityRepos repository.VacancyActivityRepository, _userRep repository.UserRepository,
	_vacancyRep repository.VacancyRepository, _notificationRep repository.NotificationRepository) *VacancyActivityService {
	return &VacancyActivityService{vacancyActivityRep: vacancyActivityRepos, userRep: _userRep,
		vacancyRep: _vacancyRep, notificationRep: _notificationRep}
}

func (vas *VacancyActivityService) GetAllVacancyApplies(vacancyId uint) ([]*models.VacancyActivityPreview, error) {
	return vas.vacancyActivityRep.GetAllVacancyApplies(vacancyId)
}

func (vas *VacancyActivityService) ApplyForVacancy(email string, vacancyId uint, input *models.VacancyActivity) (*models.Notification, error) {
	user, getErr := vas.userRep.GetUserByEmail(email)
	if getErr != nil {
		return nil, getErr
	}
	if user.UserType != "applicant" {
		return nil, errorHandler.InvalidUserType
	}

	input.UserAccountId = user.ID
	input.VacancyId = vacancyId
	err := vas.vacancyActivityRep.ApplyForVacancy(input)
	if err != nil {
		return nil, err
	}

	vacancy, getVacancyErr := vas.vacancyRep.GetById(vacancyId)
	if getVacancyErr != nil {
		return nil, getVacancyErr
	}

	notification := &models.Notification{
		UserFromID: user.ID,
		UserToID:   vacancy.PostedByUserId,
		Type:       models.AllowedNotificationTypes[0], //apply
	}
	return notification, nil
}

func (vas *VacancyActivityService) GetAllUserApplies(userId uint) ([]*models.VacancyActivityPreview, error) {

	return vas.vacancyActivityRep.GetAllUserApplies(userId)
}

func (vas *VacancyActivityService) DeleteUserApply(email string, id uint) error {
	user, getErr := vas.userRep.GetUserByEmail(email)
	if getErr != nil {
		return getErr
	}
	userId := user.ID
	return vas.vacancyActivityRep.DeleteUserApply(userId, id)
}
