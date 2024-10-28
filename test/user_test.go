package test

import (
	"fmt"
	"go-clean/config"
	"go-clean/db"
	"go-clean/middlewares"
	"go-clean/modules/user"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func SetUpRouter() *gin.Engine {
	_ = config.LoadConfig()
	middlewares.InitValidator()
	db.InitDB()
	// defer db.CloseDatabaseConnection(db.Data)

	server := gin.Default()

	server.Use(middlewares.CORSMiddleware())

	return server
}

func TestHomepageHandler(t *testing.T) {
	mockResponse := `{"message":"Welcome to my paradise"}`
	r := SetUpRouter()
	r.GET("/modules/user", user.HomepageHandler)
	req, _ := http.NewRequest("GET", "/modules/user", nil)
	fmt.Print(req)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetUserByID(t *testing.T) {
	r := SetUpRouter()

	r.GET("/api/user/:id", user.Controller.GetUsers)

	req, _ := http.NewRequest("GET", "/api/user/2", nil)
	// test case valid
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)

	example := "california"
	// test case non valid
	reqNotFound, _ := http.NewRequest("GET", "/api/user/"+example, nil) // ID 2 does not exist

	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNotFound)
	fmt.Print(w.Code)
	assert.NotEqual(t, http.StatusOK, w.Code)
}

// func TestGetAllUser(t *testing.T) {
// 	r := SetUpRouter()
// 	r.GET("/api/user", user.Controller.GetUsers)
// 	req, _ := http.NewRequest("GET", "/api/user", nil)
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)

// 	var users []user.ModelUser
// 	json.Unmarshal(w.Body.Bytes(), &users)

// 	assert.Equal(t, http.StatusOK, w.Code)
// 	assert.NotEqual(t, []user.ModelUser{}, users)
// }

// func TestCreateUser(t *testing.T) {
// 	r := SetUpRouter()
// 	r.POST("/api/user", user.Controller.CreateUser)
// 	users := user.ModelUser{
// 		Username: "waskito12",
// 		Password: "waskito123",
// 		Role:     "USER",
// 	}
// 	jsonValue, _ := json.Marshal(users)
// 	req, _ := http.NewRequest("POST", "/api/user", bytes.NewBuffer(jsonValue))

// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, req)
// 	assert.Equal(t, http.StatusCreated, w.Code)
// }

// func TestUpdateUser(t *testing.T) {
// 	r := SetUpRouter()
// 	r.PUT("/api/user/:id", user.Controller.UpdateUSer)
// 	var n uint = 2
// 	var idUser string = strconv.FormatUint(uint64(n), 10)
// 	users := user.ModelUser{
// 		ID:       2,
// 		Username: "waskito1233",
// 		Password: "$10$YMqafNr1k8nyRUcWBAEMrOuGutvKvUD2nBXeZCPm0ge6MM4hlTD..",
// 		Role:     "USER",
// 	}
// 	jsonValue, _ := json.Marshal(users)
// 	reqFound, _ := http.NewRequest("PUT", "/api/user/2"+idUser, bytes.NewBuffer(jsonValue))
// 	w := httptest.NewRecorder()
// 	r.ServeHTTP(w, reqFound)
// 	assert.Equal(t, http.StatusOK, w.Code)

// 	reqNotFound, _ := http.NewRequest("PUT", "/api/user/2", bytes.NewBuffer(jsonValue))
// 	w = httptest.NewRecorder()
// 	r.ServeHTTP(w, reqNotFound)
// 	assert.Equal(t, http.StatusNotFound, w.Code)
// }
