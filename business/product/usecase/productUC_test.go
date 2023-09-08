package usecase

import (
	"context"
	"testing"

	"message-queue-system/domain/dto/request"
	"message-queue-system/domain/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Define a mock struct for the product repository
type mockProductRepo struct {
	mock.Mock
}

func (m *mockProductRepo) UserExists(ctx context.Context, userID int64) (int, error) {
	args := m.Called(ctx, userID)
	return args.Int(0), args.Error(1)
}

func (m *mockProductRepo) InsertProduct(ctx context.Context, req request.InsertProduct) (int64, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(int64), args.Error(1)
}

func (m *mockProductRepo) UpdateImagePath(ctx context.Context, path entity.URLS, productID int64) (error) {
	return  nil
}
func (m *mockProductRepo) GetProductImage(ctx context.Context, productID int64) (entity.URLS, error) {
	return nil, nil
}
func TestInsertProductUC(t *testing.T) {
	mockRepo := new(mockProductRepo)
	mockRepo.On("UserExists", mock.Anything, mock.Anything).Return(1, nil)
	mockRepo.On("InsertProduct", mock.Anything, mock.Anything).Return(int64(1), nil)
	
	productUC := NewProductUC(mockRepo)

	var input [3]request.InsertProduct
	req1 := request.InsertProduct{
		ProductName: "Sunfeast",
		ProductDescription: "a packet of biscuit",
		ProductPrice: 43.43,
		ProductImages: []string{"https://upload.wikimedia.org/wikipedia/commons/2/23/Shampoo.jpg"},
	}
	req2 := request.InsertProduct{}
	req3 := request.InsertProduct{
		ProductName: "Sunfeast",
	}
	input[0]=req1
	input[1]=req2
	input[2]=req3
	for _, req := range input {
		err := productUC.InsertProduct(context.Background(), req)
		assert.NoError(t, err)
	}
}