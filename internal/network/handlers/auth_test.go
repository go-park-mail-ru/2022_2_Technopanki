package handlers

import (
	"HeadHunter/internal/network/middleware"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_SignUp(t *testing.T) {
	testTable := []struct {
		name                 string
		inputBody            string
		expectedStatusCode   int
		expectedResponseBody string
		isValid              bool
	}{
		{
			name: "Success case",
			inputBody: `{
    			"email": "example2@mail.ru",
    			"name": "Zahar",
    			"surname": "Urvancev",
    			"password": "111abc!!",
    			"role": "applicant"
			}`,
			expectedStatusCode:   200,
			expectedResponseBody: "",
			isValid:              true,
		},
		{
			name: "Invalid name",
			inputBody: `{
    			"email": "invalid_example@mail.ru",
    			"name": "Z",
    			"surname": "Urvancev",
    			"password": "111abc!!",
    			"role": "applicant"
			}`,
			expectedStatusCode:   400,
			expectedResponseBody: "{\"error\":\"длина имени должна быть между 3 и 20 символами\"}",
			isValid:              false,
		},
		{
			name: "Invalid surname",
			inputBody: `{
    			"email": "invalid_example@mail.ru",
    			"name": "Zahar",
    			"surname": "U",
    			"password": "111abc!!",
    			"role": "applicant"
			}`,
			expectedStatusCode:   400,
			expectedResponseBody: "{\"error\":\"длина фамилии должна быть между 3 и 20 символами\"}",
			isValid:              false,
		},
		{
			name: "Invalid password 1",
			inputBody: `{
    			"email": "invalid_example@mail.ru",
    			"name": "Zahar",
    			"surname": "Urvancev",
    			"password": "1a!",
    			"role": "applicant"
			}`,
			expectedStatusCode:   400,
			expectedResponseBody: "{\"error\":\"длина пароля должна быть между 8 и 20 символами\"}",
			isValid:              false,
		},
		{
			name: "Invalid password 2",
			inputBody: `{
    			"email": "invalid_example@mail.ru",
    			"name": "Zahar",
    			"surname": "Urvancev",
    			"password": "111asdwwqv",
    			"role": "applicant"
			}`,
			expectedStatusCode:   400,
			expectedResponseBody: "{\"error\":\"пароль должен содержать буквы латиницы, цифры и спецсимволы(!#%^$)\"}",
			isValid:              false,
		},
	}
	for _, testCase := range testTable {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			router := gin.New()
			router.POST("/auth/sign-up", SignUp, middleware.ErrorHandler())
			recorder := httptest.NewRecorder()

			req, reqErr := http.NewRequest("POST", "/auth/sign-up",
				bytes.NewBufferString(tc.inputBody))
			require.NoError(t, reqErr)

			router.ServeHTTP(recorder, req)

			if tc.isValid && len(recorder.Result().Cookies()) > 0 {
				_, sessionTokenErr := uuid.Parse(recorder.Result().Cookies()[0].Value)
				assert.NoError(t, sessionTokenErr)
			} else {
				assert.Equal(t, []*http.Cookie{}, recorder.Result().Cookies())
			}

			assert.Equal(t, tc.expectedStatusCode, recorder.Code)
			assert.Equal(t, tc.expectedResponseBody, recorder.Body.String())
		})
	}
}
