package handlers

//func TestGetVacancies(t *testing.T) {
//	gin.SetMode(gin.TestMode)
//
//	t.Run("Success", func(t *testing.T) {
//		uid, _ := uuid.NewRandom()
//
//		mockUserResp := &model.User{
//			UID:   uid,
//			Email: "bob@bob.com",
//			Name:  "Bobby Bobson",
//		}
//
//		mockUserService := new(mocks.MockUserService)
//		mockUserService.On("Get", mock.AnythingOfType("*gin.Context"), uid).Return(mockUserResp, nil)
//
//		// a response recorder for getting written http response
//		rr := httptest.NewRecorder()
//
//		// use a middleware to set context for test
//		// the only claims we care about in this test
//		// is the UID
//		router := gin.Default()
//		router.Use(func(c *gin.Context) {
//			c.Set("user", &model.User{
//				UID: uid,
//			},
//			)
//		})
//
//		NewHandler(&Config{
//			R:           router,
//			UserService: mockUserService,
//		})
//
//		request, err := http.NewRequest(http.MethodGet, "/me", nil)
//		assert.NoError(t, err)
//
//		router.ServeHTTP(rr, request)
//
//		respBody, err := json.Marshal(gin.H{
//			"user": mockUserResp,
//		})
//		assert.NoError(t, err)
//
//		assert.Equal(t, 200, rr.Code)
//		assert.Equal(t, respBody, rr.Body.Bytes())
//		mockUserService.AssertExpectations(t) // assert that UserService.Get was called
//	})
//
//}
