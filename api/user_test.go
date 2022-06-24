package api

import (
	mockdb "andre/notesnotes-api/db/mock"
	db "andre/notesnotes-api/db/sqlc"
	"andre/notesnotes-api/token"
	"andre/notesnotes-api/util"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetUserAPI(t *testing.T) {
	user := randomUser(t)

	testCases := []struct {
		name          string
		userID        int32
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name:   "OK",
			userID: user.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyReponseMatchUser(t, recorder.Body, user)
			},
		},
		{
			name:   "NotFound",
			userID: user.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(db.User{}, sql.ErrNoRows)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusNotFound, recorder.Code)
			},
		},
		{
			name:   "InternalError",
			userID: user.ID,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Eq(user.ID)).
					Times(1).
					Return(db.User{}, sql.ErrConnDone)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusInternalServerError, recorder.Code)
			},
		},
		{
			name:   "BadRequestInvalidID",
			userID: 0,
			setupAuth: func(t *testing.T, request *http.Request, tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, authorizationTypeBearer, user.Username, time.Minute)
			},
			buildStubs: func(store *mockdb.MockStore) {
				store.EXPECT().
					GetUser(gomock.Any(), gomock.Any()).
					Times(0)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusBadRequest, recorder.Code)
			},
		},
		// add more cases
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// start test server
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			url := fmt.Sprintf("/user/%d", tc.userID)
			request, error := http.NewRequest(http.MethodGet, url, nil)
			require.NoError(t, error)

			tc.setupAuth(t, request, server.tokenMaker)
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

type eqCreateUserParamMatcher struct {
	arg      db.CreateUsersParams
	password string
}

func (e eqCreateUserParamMatcher) Matches(x interface{}) bool {
	arg, ok := x.(db.CreateUsersParams)
	if !ok {
		return false
	}
	err := util.CheckPassword(e.password, arg.Password)
	if err != nil {
		return false
	}

	e.arg.Password = arg.Password
	return reflect.DeepEqual(e.arg, arg)
}

func (e eqCreateUserParamMatcher) String() string {
	return fmt.Sprintf("matches arg %v and password %v", e.arg, e.password)
}

func EqCreateUserParam(arg db.CreateUsersParams, password string) gomock.Matcher {
	return eqCreateUserParamMatcher{arg, password}
}

func TestCreateUserAPI(t *testing.T) {
	user := randomUser(t)

	testCases := []struct {
		name          string
		body          gin.H
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			body: gin.H{
				"first_name": user.FirstName,
				"last_name":  user.LastName.String,
				"username":   user.Username,
				"email":      user.Email,
				"password":   user.Password,
			},
			buildStubs: func(store *mockdb.MockStore) {
				arg := db.CreateUsersParams{
					FullName:  user.FullName,
					FirstName: user.FirstName,
					LastName:  user.LastName,
					Username:  user.Username,
					Email:     user.Email,
					Password:  user.Password,
				}
				store.EXPECT().
					CreateUsers(gomock.Any(), EqCreateUserParam(arg, arg.Password)).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyReponseMatchUser(t, recorder.Body, user)
			},
		},
		// add more cases
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)
			tc.buildStubs(store)

			// start test server
			server := newTestServer(t, store)
			recorder := httptest.NewRecorder()

			data, err := json.Marshal(tc.body)
			require.NoError(t, err)

			url := "/user"
			request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}

func randomUser(t *testing.T) db.User {
	hashedPassword, err := util.HashPassword(util.RandomString(6))
	require.NoError(t, err)

	firstName := util.RandomString(10)
	lastName := sql.NullString{String: util.RandomString(10), Valid: true}
	fullName := firstName + " " + lastName.String

	return db.User{
		ID:        util.RandomInt(1, 1000),
		FullName:  fullName,
		FirstName: firstName,
		LastName:  lastName,
		Username:  util.RandomString(10),
		Email:     util.RandomString(20) + "@email.com",
		Password:  hashedPassword,
	}
}

func requireBodyReponseMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := ioutil.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)

	response := newUserResponse(user)

	require.Equal(t, user.FullName, response.FullName)
	require.Equal(t, user.FirstName, response.FirstName)
	require.Equal(t, user.LastName, response.LastName)
	require.Equal(t, user.Username, response.Username)
	require.Equal(t, user.Email, response.Email)
	require.Equal(t, user.CreatedAt, response.CreatedAt)
	require.Equal(t, user.UpdatedAt, response.UpdatedAt)
	require.Equal(t, user.NotesCount, response.NotesCount)
}
