package handlers

import (
	"HeadHunter/internal/network/middleware"
	"bytes"
	"github.com/gin-gonic/gin"
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
    			"email": "example@mail.ru",
    			"name": "Zahar",
    			"surname": "Urvancev",
    			"password": "111abc!!",
    			"role": "applicant"
			}`,
			expectedStatusCode:   200,
			expectedResponseBody: "",
		},
		//{
		//	name: "Bad body case",
		//	inputBody: `{
		//		HEELLOO WORLD
		//	}`,
		//	expectedStatusCode: 400,
		//	expectedResponseBody: `{
		//		"error": "bad request"
		//	}`,
		//},
	}
	for _, testCase := range testTable {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			router := gin.New()
			router.POST("/auth/sign-up", SignUp, middleware.ErrorHandler())
			router.Use(middleware.ErrorHandler())
			recorder := httptest.NewRecorder()

			http.SetCookie(recorder, &http.Cookie{Name: "session", Value: "test_value"})
			req, reqErr := http.NewRequest("POST", "/auth/sign-up",
				bytes.NewBufferString(tc.inputBody))
			require.NoError(t, reqErr)

			router.ServeHTTP(recorder, req)
			assert.Equal(t, "test_value", recorder.Result().Cookies()[0].Value)
			assert.Equal(t, tc.expectedStatusCode, recorder.Code)
			assert.Equal(t, tc.expectedResponseBody, recorder.Body.String())
		})
	}
}
