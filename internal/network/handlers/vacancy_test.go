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

func Test_GetVacancies(t *testing.T) {
	testTable := []struct {
		name                 string
		query                string
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:                 "Success case",
			query:                "?id=1",
			expectedStatusCode:   200,
			expectedResponseBody: "{\n    \"title\": \"Android-разработчик\",\n    \"description\": \"Мы разрабатываем сложную рекомендательную систему, делаем приложения под Android и iOS, интегрируем модули Дзена в другие сервисы и пилим свой редактор видео. Все, чтобы авторы нашли аудиторию, а пользователи — то, что им интересно.\",\n    \"minimal_salary\": 100000,\n    \"maximum_salary\": 150000,\n    \"employer_name\": \"VK\",\n    \"city\": \"Moscow\",\n    \"date\": \"0001-01-01T00:00:00Z\"\n}",
		},
		{
			name:                 "Invalid query",
			query:                "?id=asda",
			expectedStatusCode:   400,
			expectedResponseBody: "{\"error\":\"invalid query\"}",
		},
		{
			name:                 "Vacancy not found",
			query:                "?id=1000",
			expectedStatusCode:   404,
			expectedResponseBody: "{\"error\":\"vacancy not found\"}",
		},
	}
	for _, testCase := range testTable {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			router := gin.New()
			router.Use(func(c *gin.Context) {
				c.Set("userEmail", "testVal")
			})
			router.GET("/api/vacancy", GetVacancies, middleware.ErrorHandler())
			recorder := httptest.NewRecorder()

			req, reqErr := http.NewRequest("GET", "/api/vacancy"+tc.query,
				bytes.NewBufferString(""))
			require.NoError(t, reqErr)

			router.ServeHTTP(recorder, req)

			assert.Equal(t, tc.expectedStatusCode, recorder.Code)
			assert.Equal(t, tc.expectedResponseBody, recorder.Body.String())
		})
	}
}
