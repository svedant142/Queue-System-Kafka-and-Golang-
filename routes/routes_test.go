package routes

import (
	"encoding/json"
	"fmt"
	"message-queue-system/db"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

//BenchMark Test
func BenchmarkTest_API(b *testing.B) {
	testCase := fmt.Sprintf("/user/item?user_id=%d&product_name=%s", 1, "abc")
	db.InitMYSQL()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		router := gin.Default()
		InitRoutes(router)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", testCase, nil)
		router.ServeHTTP(w, req)
	}
}


//Integration Test - Success
func TestIntegrationSuccess(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db.InitMYSQL()
	var TestCases [3]string 
	TestCases[0] = fmt.Sprintf("/user/item?user_id=%d&product_name=%s",1,"abc")
	TestCases[1] = fmt.Sprintf("/user/item?user_id=%d&product_name=%s",1,"   abc    ")
	TestCases[2] = fmt.Sprintf("/user/item?user_id=%d&product_name=%s&product_description=%s&product_images=%s&product_price=%f&product_images=%s",1,"soap","a soap","https://upload.wikimedia.org/wikipedia/commons/9/9b/Handmade_soap_cropped_and_simplified.jpg",
	43.43,"https://upload.wikimedia.org/wikipedia/commons/2/23/Shampoo.jpg")
	for _, path := range TestCases {
		router := gin.Default()
		InitRoutes(router)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", path, nil)
		router.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, "successful insertion", response["message"])
		assert.Equal(t, true, response["status"])
	}
}

//Integration Test - Failure
func TestIntegrationFailure(t *testing.T) {
	gin.SetMode(gin.TestMode)
	db.InitMYSQL()
	var TestCases [5]string 
	TestCases[0] = fmt.Sprintf("/user/item?user_id=%d&product_name=%s",0,"abc") //invalid userID
	TestCases[1] = fmt.Sprintf("/user/item?user_id=%d&product_name=%s",1," ") //invalid productName
	TestCases[2] = fmt.Sprintf("/user/item?user_id=%d&product_name=%s",2,"") //invalid productName
	TestCases[3] = fmt.Sprintf("/user/item?user_id=%d&product_name=%s",-1,"") //invalid userID
	TestCases[4] = fmt.Sprintf("/user/item?user_id=%d&product_name=%s",200,"soap") //user does not exist in DB

	for _, path := range TestCases {
		router := gin.Default()
		InitRoutes(router)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", path, nil)
		router.ServeHTTP(w, req)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, false, response["status"])
	}
}