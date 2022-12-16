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

func (np *NotificationPostgres) notificationApplyQuery() *gorm.DB {
	/*select notifications.id, notifications.type, notifications.user_to_id, notifications.user_from_id, company.company_name,
	applicant.applicant_name, vacancies.title, notifications.object_id, notifications.is_viewed
	from notifications
	left join user_accounts as company on
	notifications.user_to_id = company.id
	left join vacancies on notifications.object_id = vacancies.id
	left join user_accounts as applicant on notifications.user_from_id = applicant.id
	where notifications.user_to_id = 1;
	*/
	return np.db.Table("notifications").
		Select("notifications.id,notifications.type, notifications.user_to_id, notifications.user_from_id, applicant.applicant_name, " +
			"notifications.is_viewed, company.company_name, vacancies.title, notifications.object_id").
		Joins("left join user_accounts as applicant on notifications.user_from_id =applicant.id").
		Joins("left join user_accounts as company on notifications.user_to_id = company.id").
		Joins("left join vacancies on vacancies.id = notifications.object_id")

}

func (np *NotificationPostgres) notificationDownloadPDFQuery() *gorm.DB {
	return np.db.Table("notifications").
		Select("notifications.id,notifications.type, notifications.user_to_id, notifications.user_from_id, notifications.is_viewed " +
			"applicant.applicant_name, company.company_name, resumes.title, resumes.id as object_id").
		Joins("left join user_accounts as applicant on notifications.user_to_id =applicant.id").
		Joins("left join user_accounts as company on notifications.user_from_id = company.id").
		Joins("left join resumes on resumes.user_account_id = applicant.id")
}

func (np *NotificationPostgres) GetNotificationPreviewApply(id uint) (*models.NotificationPreview, error) {
	var result *models.NotificationPreview
	query := np.notificationApplyQuery().Where("notifications.id = ?", id).Scan(&result)
	return result, QueryValidation(query, "notification")
}

func (np *NotificationPostgres) GetNotificationPreviewDownloadPDF(id uint) (*models.NotificationPreview, error) {
	var result *models.NotificationPreview
	query := np.notificationDownloadPDFQuery().Where("notifications.id = ?", id).Scan(&result)
	return result, QueryValidation(query, "notification")
}

func (np *NotificationPostgres) GetApplyNotificationsByUser(id uint) ([]*models.NotificationPreview, error) {
	var result []*models.NotificationPreview
	query := np.notificationApplyQuery().Where("user_to_id = ?", id).Scan(&result)
	return result, QueryValidation(query, "notification")
}

func (np *NotificationPostgres) GetDownloadPDFNotificationsByUser(id uint) ([]*models.NotificationPreview, error) {
	var result []*models.NotificationPreview
	query := np.notificationDownloadPDFQuery().Where("user_to_id = ?", id).Scan(&result)
	return result, QueryValidation(query, "notification")

}

func (np *NotificationPostgres) CreateNotification(notification *models.Notification) error {
	creatingErr := np.db.Create(notification).Save(notification).Error

	if creatingErr != nil {
		return fmt.Errorf("cannot create notification: %w", creatingErr)
	}

	return nil
}
func (np *NotificationPostgres) ReadNotification(id uint) error {
	return np.db.Model(&models.Notification{ID: id}).Update("is_viewed", true).Error
}

func (np *NotificationPostgres) GetNotification(id uint) (*models.Notification, error) {
	var result *models.Notification
	query := np.db.Where("id = ?", id).Find(&result)
	return result, QueryValidation(query, "notification")
}

func (np *NotificationPostgres) DeleteNotificationsFromUser(userId uint) error {
	return np.db.Where("user_to_id = ?", userId).Delete(&models.Notification{}).Error
}
