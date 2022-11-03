package constants

// PrivateUserFields Поля, доступ к которым по умолчанию не доступен
var PrivateUserFields = []string{"email", "contact_number",
	"applicant_current_salary"}

// SafeUserFields Поля, доступ к которым доступен всегда
var SafeUserFields = []string{"id", "user_type", "description", "date_of_birth", "image",
	"applicant_name", "applicant_surname", "company_name", "business_type", "company_website_url"}
