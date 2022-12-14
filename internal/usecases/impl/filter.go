package impl

func VacancyFilterQueries(filterName string) string {
	switch filterName {
	case "Title":
		return "title ILIKE ?"
	case "Location":
		return "location = ?"
	case "Format":
		return "format = ?"
	case "Experience":
		return "experience = ?"
	case "FirstSalaryValue":
		return "salary BETWEEN ? AND ?"
	default:
		return ""
	}
}

func EmployerFilterQueries(filterName string) string {
	switch filterName {
	case "CompanyName":
		return "company_name ILIKE ? AND user_type = 'employer'"
	case "Location":
		return "location = ? AND user_type = 'employer'"
	case "BusinessType":
		return "business_type = ?"
	case "FirstCompanySizeValue":
		return "company_size BETWEEN ? AND ? AND user_type = 'employer'"
	default:
		return ""
	}
}

func ApplicantFilterQueries(filterName string) string {
	switch filterName {
	case "ApplicantName":
		return "applicant_name ILIKE ? AND user_type = 'applicant'"
	case "ApplicantSurname":
		return "applicant_surname ILIKE ? AND user_type = 'applicant'"
	case "Location":
		return "location = ? AND user_type = 'applicant'"
	case "FirstAgeValue":
		return "age BETWEEN ? AND ?"
	default:
		return ""
	}
}

func ResumeFilterQueries(filterName string) string {
	switch filterName {
	case "Title":
		return "title ILIKE ?"
	case "Location":
		return "location = ?"
	case "ExperienceInYears":
		return "experience_in_years = ?"
	case "FirstSalaryValue":
		return "salary BETWEEN ? AND ?"
	default:
		return ""
	}
}
