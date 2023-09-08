package webhttp

import (
	"context"
	"encoding/json"
	"fmt"
	Requests "message-queue-system/domain/dto/request"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockProductUC struct{}

func (m *mockProductUC) InsertProduct(c context.Context, req Requests.InsertProduct) error {
	return nil
}

func TestInsert_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var TestCases [3]string 
	TestCases[0] = fmt.Sprintf("/user/item?user_id=%d&product_name=%s",1,"abc")
	TestCases[1] = fmt.Sprintf("/user/item?user_id=%d&product_name=%s",1,"   abc    ")
	TestCases[2] = fmt.Sprintf("/user/item?user_id=%d&product_name=%s&product_description=%s&product_images=%s&product_price=%f&product_images=%s",1,"soap","a soap","https://upload.wikimedia.org/wikipedia/commons/9/9b/Handmade_soap_cropped_and_simplified.jpg",
	43.43,"https://upload.wikimedia.org/wikipedia/commons/2/23/Shampoo.jpg")
	for _, path := range TestCases {
		router := gin.New()
		productUC := &mockProductUC{}
		Init(router, productUC)
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

func TestInsert_Failure(t *testing.T) {
	gin.SetMode(gin.TestMode)
	var TestCases [4]string 
	TestCases[0] = fmt.Sprintf("/user/item?user_id=%d&product_name=%s",0,"abc")
	TestCases[1] = fmt.Sprintf("/user/item?user_id=%d&product_name=%s",1," ")
	TestCases[2] = fmt.Sprintf("/user/item?user_id=%d&product_name=%s",2,"")
	TestCases[3] = fmt.Sprintf("/user/item?user_id=%d&product_name=%s",-1,"")

	for _, path := range TestCases {
		router := gin.New()
		productUC := &mockProductUC{}
		Init(router, productUC)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", path, nil)
		router.ServeHTTP(w, req)

		var response map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.NoError(t, err)
		assert.Equal(t, false, response["status"])
	}
}
