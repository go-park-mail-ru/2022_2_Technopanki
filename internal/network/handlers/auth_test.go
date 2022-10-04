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

			if len(recorder.Result().Cookies()) > 0 {
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

func Test_SignIn(t *testing.T) {
	testTable := []struct {
		name                 string
		inputBody            string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Success case",
			inputBody: `{
    			"email": "example@mail.ru",
    			"password": "123456!!a"
			}`,
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		{
			name: "user does not exist",
			inputBody: `{
    			"email": "examplee@mail.ru",
    			"password": "123456!!a"
			}`,
			expectedStatusCode:   401,
			expectedResponseBody: "{\"error\":\"пользователя с таким email не существует\"}",
		},
		{
			name: "invalid json",
			inputBody: `{
    			"email": "examplee@mail.ru",
    			"password": "123456!!a",
			}`,
			expectedStatusCode:   400,
			expectedResponseBody: "",
		},
	}
	for _, testCase := range testTable {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			router := gin.New()
			router.POST("/auth/sign-in", SignIn, middleware.ErrorHandler())
			recorder := httptest.NewRecorder()

			req, reqErr := http.NewRequest("POST", "/auth/sign-in",
				bytes.NewBufferString(tc.inputBody))
			require.NoError(t, reqErr)

			router.ServeHTTP(recorder, req)

			if len(recorder.Result().Cookies()) > 0 {
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
