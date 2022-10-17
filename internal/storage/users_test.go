package storage

//
//func Test_FindByEmail(t *testing.T) {
//	testTable := []struct {
//		name          string
//		email         string
//		expectedUser  entity.User
//		expectedError error
//	}{
//		{
//			name:  "user is found",
//			email: "example@mail.ru",
//			expectedUser: entity.User{
//				Name:     "Zakhar",
//				Surname:  "Urvancev",
//				Email:    "example@mail.ru",
//				Password: string(password),
//				Role:     "applicant",
//			},
//			expectedError: nil,
//		},
//		{
//			name:          "user is not found",
//			email:         "example2@mail.ru",
//			expectedUser:  entity.User{},
//			expectedError: errorHandler.ErrUserNotExists,
//		},
//	}
//	for _, testCase := range testTable {
//		tc := testCase
//		t.Run(tc.name, func(t *testing.T) {
//			t.Parallel()
//			result, errResult := UserStorage.FindByEmail(tc.email)
//
//			assert.Equal(t, tc.expectedUser, result)
//			assert.Equal(t, tc.expectedError, errResult)
//		})
//
//	}
//}

//func Test_AddUser(t *testing.T) {
//	testTable := []struct {
//		name string
//		user entity.User
//	}{
//		{
//			name: "default case",
//			user: entity.User{
//				Name:     "Zakhar",
//				Surname:  "Urvancev",
//				Email:    "example@mail.ru",
//				Password: "123456!!a",
//				Role:     "applicant",
//			},
//		},
//	}
//	for _, testCase := range testTable {
//		tc := testCase
//		t.Run(tc.name, func(t *testing.T) {
//			t.Parallel()
//			result := UserStorage.AddUser(tc.user)
//			require.NoError(t, result)
//			_, err := UserStorage.FindByEmail(tc.user.Email)
//			require.Equal(t, err, nil)
//		})
//
//	}
//}
