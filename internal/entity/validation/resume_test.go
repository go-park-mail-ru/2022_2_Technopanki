package validation

import (
	"HeadHunter/configs"
	"HeadHunter/internal/entity/models"
	"HeadHunter/internal/errorHandler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ResumeValidaion(t *testing.T) {
	testTable := []struct {
		name        string
		inputResume *models.Resume
		cfg         configs.ValidationConfig
		expected    error
	}{
		{
			name: "ok",
			inputResume: &models.Resume{
				Title:       "Job",
				Description: "Information about my skills in this job",
			},
			cfg: configs.ValidationConfig{
				MinResumeDescriptionLength: 10,
				MinResumeTitleLength:       3,
				MaxResumeTitleLength:       30,
			},
			expected: nil,
		},
		{
			name: "incorrect title",
			inputResume: &models.Resume{
				Title:       "",
				Description: "Information about my skills in this job",
			},
			cfg: configs.ValidationConfig{
				MinResumeDescriptionLength: 10,
				MinResumeTitleLength:       3,
				MaxResumeTitleLength:       30,
			},
			expected: errorHandler.InvalidResumeTitleLength,
		},
		{
			name: "incorrect description",
			inputResume: &models.Resume{
				Title:       "Some job",
				Description: "",
			},
			cfg: configs.ValidationConfig{
				MinResumeDescriptionLength: 10,
				MinResumeTitleLength:       3,
				MaxResumeTitleLength:       30,
			},
			expected: errorHandler.InvalidResumeDescriptionLength,
		},
		{
			name: "try to change user",
			inputResume: &models.Resume{
				Title:       "Some job",
				Description: "Information about my skills in this job",
				UserName:    "Hacker_name",
			},
			cfg: configs.ValidationConfig{
				MinResumeDescriptionLength: 10,
				MinResumeTitleLength:       3,
				MaxResumeTitleLength:       30,
			},
			expected: errorHandler.ErrBadRequest,
		},
	}
	for _, test := range testTable {
		testCase := test
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			result := ResumeValidaion(testCase.inputResume, testCase.cfg)
			assert.Equal(t, testCase.expected, result)
		})
	}
}
