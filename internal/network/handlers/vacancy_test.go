package handlers

//func Test_GetVacancies(t *testing.T) {
//	testTable := []struct {
//		name                 string
//		inputQuery           int
//		expectedStatusCode   int
//		expectedResponseBody string
//	}{
//		{
//			name:                 "Success case",
//			inputQuery:           5,
//			expectedStatusCode:   200,
//			expectedResponseBody: "",
//		},
//	}
//	for _, testCase := range testTable {
//		tc := testCase
//		t.Run(tc.name, func(t *testing.T) {
//			t.Parallel()
//			router := gin.New()
//			router.POST("/auth/sign-up", SignUp, middleware.ErrorHandler())
//			recorder := httptest.NewRecorder()
//
//			req, reqErr := http.NewRequest("POST", "/auth/sign-up?id=",
//				bytes.NewBufferString(""))
//			require.NoError(t, reqErr)
//
//			router.ServeHTTP(recorder, req)
//
//			if len(recorder.Result().Cookies()) > 0 {
//				_, sessionTokenErr := uuid.Parse(recorder.Result().Cookies()[0].Value)
//				assert.NoError(t, sessionTokenErr)
//			} else {
//				assert.Equal(t, []*http.Cookie{}, recorder.Result().Cookies())
//			}
//
//			assert.Equal(t, tc.expectedStatusCode, recorder.Code)
//			assert.Equal(t, tc.expectedResponseBody, recorder.Body.String())
//		})
//	}
//}
