package impl

import (
	"HeadHunter/internal/entity/models"
	"fmt"
	"gorm.io/gorm"
)

type NotificationPostgres struct {
	db *gorm.DB
}

func NewNotificationPostgres(db *gorm.DB) *NotificationPostgres {
	return &NotificationPostgres{db: db}
}

func (np *NotificationPostgres) GetNotificationsApply(id uint) ([]*models.NotificationPreview, error) {
	var result []*models.NotificationPreview
	query := np.db.Table("notifications").
		Select("notifications.id,notifications.type, notifications.user_to_id, notifications.user_from_id, applicant.applicant_name,"+
			"company.company_name, resumes.title").
		Joins("left join user_accounts as applicant on notifications.user_from_id =applicant.id").
		Joins("left join user_accounts as company on notifications.user_to_id = company.id").
		Joins("left join resumes on resumes.user_account_id = applicant.id").Where("user_to_id = ?", id).Scan(&result)
	return result, QueryValidation(query, "notification")

}

func (np *NotificationPostgres) GetNotificationsDownloadPDF(id uint) ([]*models.NotificationPreview, error) {
	var result []*models.NotificationPreview
	query := np.db.Table("notifications").
		Select("notifications.id,notifications.type, notifications.user_to_id, notifications.user_from_id, applicant.applicant_name,"+
			"company.company_name, vacancies.title").
		Joins("left join user_accounts as applicant on notifications.user_to_id =applicant.id").
		Joins("left join user_accounts as company on notifications.user_from_id = company.id").
		Joins("left join vacancies on vacancies.posted_by_user_id = company.id").Where("user_to_id = ?", id).Scan(&result)
	return result, QueryValidation(query, "notification")

}

func (np *NotificationPostgres) CreateNotification(notification *models.Notification) error {
	creatingErr := np.db.Create(notification).Save(notification).Error

	if creatingErr != nil {
		return fmt.Errorf("cannot create notification: %w", creatingErr)
	}

	return nil
}
