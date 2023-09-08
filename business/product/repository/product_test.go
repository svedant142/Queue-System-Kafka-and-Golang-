package repository

import (
	"context"
	"database/sql"
	"testing"

	"message-queue-system/domain/dto/request"
	"message-queue-system/domain/entity"

	"message-queue-system/db"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockDB struct {
	Db *sql.DB
	mock.Mock
}

func (m *mockDB) Conn(ctx context.Context) (*sql.Conn, error) {
	args := m.Called(ctx)
	return args.Get(0).(*sql.Conn), args.Error(1)
}

func TestInsertProduct_Success(t *testing.T) {
	mockDb := new(mockDB)
	db.InitMYSQL()	
	mockDb.Db = db.DbClient

	productRepo := NewProductRepo(mockDb.Db)
	var input [3]request.InsertProduct
	req1 := request.InsertProduct{
		//base case	
		ProductName: "Sunfeast",	
		ProductDescription: "a packet of biscuit",
		ProductPrice: 43.43,
		ProductImages: []string{"https://upload.wikimedia.org/wikipedia/commons/2/23/Shampoo.jpg"}, 
	}
	req2 := request.InsertProduct{
		// multiple values for Images
		ProductName: "soap",
		ProductDescription: "a packet of soap",
		ProductPrice: 43.43,
		ProductImages: []string{"https://upload.wikimedia.org/wikipedia/commons/2/23/Shampoo.jpg",
														"https://upload.wikimedia.org/wikipedia/commons/2/23/Shampoo.jpg"},	
	}
	req3 := request.InsertProduct{
		//Empty fields except ProductName
		ProductName: "Sunfeast",	
	}
	input[0]=req1
	input[1]=req2
	input[2]=req3
	for _, req := range input {
		_, err := productRepo.InsertProduct(context.Background(), req)
		assert.NoError(t, err)
	}
}

func TestInsertProduct_Failure(t *testing.T) {
	mockDb := new(mockDB)
	db.InitMYSQL()
	mockDb.Db = db.DbClient
	mockDb.On("Conn", mock.Anything).Return(&sql.Conn{}, nil)

	productRepo := NewProductRepo(mockDb.Db)
	var input [3]request.InsertProduct
	req1 := request.InsertProduct{
		//product name is made a required field, otherwise empty data will be entered , i.e, handling empty name for repository separate
		ProductDescription: "a packet of biscuit",	
		ProductPrice: 43.43,
		ProductImages: []string{"https://upload.wikimedia.org/wikipedia/commons/2/23/Shampoo.jpg"},
	}
	//empty values
	req2 := request.InsertProduct{} 
	req3 := request.InsertProduct{
		//out of range price
		ProductName: "soap",
		ProductDescription: "a packet of soap",
		ProductPrice: 99999999999.999999,
		ProductImages: []string{"https://upload.wikimedia.org/wikipedia/commons/2/23/Shampoo.jpg"},
	}
	input[0]=req1
	input[1]=req2
	input[2]=req3
	for _, req := range input {
		_, err := productRepo.InsertProduct(context.Background(), req)
		assert.Error(t, err)
	}
}

func TestGetProductImage_Success(t *testing.T) {
	mockDb := new(mockDB)
	db.InitMYSQL()	
	mockDb.Db = db.DbClient
	mockDb.On("Conn", mock.Anything).Return(&sql.Conn{}, nil)

	productRepo := NewProductRepo(mockDb.Db)
	var input [3]int64
	input[0]=1
	input[1]=2
	input[2]=12
	for _, req := range input {
		_, err := productRepo.GetProductImage(context.Background(), req)
		assert.NoError(t, err)
	}
}

func TestGetProductImage_Failure(t *testing.T) {
	mockDb := new(mockDB)
	db.InitMYSQL()	
	mockDb.Db = db.DbClient
	mockDb.On("Conn", mock.Anything).Return(&sql.Conn{}, nil)
	productRepo := NewProductRepo(mockDb.Db)
	var input [3]int64
	input[0]=0
	input[1]=32312
	input[2]=-1
	for _, req := range input {
		_, err := productRepo.GetProductImage(context.Background(), req)
		assert.Error(t, err)
	}
}


func TestUpdateImagePath(t *testing.T) {
	mockDb := new(mockDB)
	db.InitMYSQL()	
	mockDb.Db = db.DbClient
	mockDb.On("Conn", mock.Anything).Return(&sql.Conn{}, nil)

	productRepo := NewProductRepo(mockDb.Db)
	var reqPath entity.URLS = []string{"E:\\Vedant\\GoPrograaming\\productID-1","some/path"}
	var reqProductID int64 = 12
	err := productRepo.UpdateImagePath(context.Background(), reqPath, reqProductID)
	assert.NoError(t, err)
}

func TestUserExists(t *testing.T) {
	mockDb := new(mockDB)
	db.InitMYSQL()	
	mockDb.Db = db.DbClient
	mockDb.On("Conn", mock.Anything).Return(&sql.Conn{}, nil)

	productRepo := NewProductRepo(mockDb.Db)
	var userID int64 = 1 
	count, err := productRepo.UserExists(context.Background(), userID)
	assert.NoError(t, err)
	assert.Equal(t,1,count) 
}
