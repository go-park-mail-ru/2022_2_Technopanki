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
			query:                "",
			expectedStatusCode:   200,
			expectedResponseBody: "{\n    \"1\": {\n        \"title\": \"Android-разработчик\",\n        \"description\": \"Мы разрабатываем сложную рекомендательную систему, делаем приложения под Android и iOS, интегрируем модули Дзена в другие сервисы и пилим свой редактор видео. Все, чтобы авторы нашли аудиторию, а пользователи — то, что им интересно.\",\n        \"minimal_salary\": 100000,\n        \"maximum_salary\": 150000,\n        \"date\": \"0001-01-01T00:00:00Z\"\n    },\n    \"2\": {\n        \"title\": \"Golang разработчик\",\n        \"description\": \"Мы приглашаем в свою команду разработчика с отличным знанием Go (golang) для разработки высоконагруженных сервисов, которые помогают тысячам клиентов в области транспорта решать их профессиональные задачи\",\n        \"minimal_salary\": 50000,\n        \"maximum_salary\": 150000,\n        \"date\": \"0001-01-01T00:00:00Z\"\n    },\n    \"3\": {\n        \"title\": \"Middle Golang Developer\",\n        \"description\": \"Группа Компаний «Flow inc.» открывает набор на позицию Middle Golang разработчик, в сильный и амбициозный коллектив. Мы - финтех компания, которая разработала платежный шлюз, работающий напрямую с Qiwi, разрабатываем высоконагруженные сервисы с отказоустойстойчивой архитектурой;\",\n        \"minimal_salary\": 0,\n        \"maximum_salary\": 350000,\n        \"date\": \"0001-01-01T00:00:00Z\"\n    },\n    \"4\": {\n        \"title\": \"Онлайн-репетитор по математике для подготовки к ЕГЭ/ОГЭ\",\n        \"description\": \"В онлайн-школу требуется преподаватель математики по подготовке к ЕГЭ, ОГЭ, олимпиадам для проведения занятий с помощью современных образовательных технологий: онлайн-платформы и интерактивных инструментов.\",\n        \"minimal_salary\": 48000,\n        \"maximum_salary\": 96000,\n        \"date\": \"0001-01-01T00:00:00Z\"\n    },\n    \"5\": {\n        \"title\": \"IT-наставник для детей по разработке игр (Minecraft, Roblox, Unity, Construct 3, Unreal Engine)\",\n        \"description\": \"В онлайн-школу требуется преподаватель компьютерных курсов для детей по разработке игр\",\n        \"minimal_salary\": 0,\n        \"maximum_salary\": 80000,\n        \"date\": \"0001-01-01T00:00:00Z\"\n    }\n}",
		},
		{
			name:                 "Success case(query)",
			query:                "?id=1",
			expectedStatusCode:   200,
			expectedResponseBody: "{\n    \"title\": \"Android-разработчик\",\n    \"description\": \"Мы разрабатываем сложную рекомендательную систему, делаем приложения под Android и iOS, интегрируем модули Дзена в другие сервисы и пилим свой редактор видео. Все, чтобы авторы нашли аудиторию, а пользователи — то, что им интересно.\",\n    \"minimal_salary\": 100000,\n    \"maximum_salary\": 150000,\n    \"date\": \"0001-01-01T00:00:00Z\"\n}",
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
