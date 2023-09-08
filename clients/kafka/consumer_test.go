package kafka

import (
	"context"
	"encoding/json"
	"testing"

	"message-queue-system/db"
	"message-queue-system/domain/dto/request"
	"message-queue-system/domain/entity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock struct for the Consumer
type mockConsumer struct {
	mock.Mock
}

type MockIProductRepo struct {
	mock.Mock
}

func (m *MockIProductRepo) UserExists(ctx context.Context, userID int64) (int, error) {
	return 0, nil
}

func (m *MockIProductRepo) InsertProduct(ctx context.Context, req request.InsertProduct) (int64, error) {
	return 0, nil
}

func (m *MockIProductRepo) UpdateImagePath(ctx context.Context, path entity.URLS, productID int64) (error) {
	return  nil
}
func (m *MockIProductRepo) GetProductImage(ctx context.Context, productID int64) (entity.URLS, error) {
	return nil, nil
}

func TestProcessImages(t *testing.T) {
	mockProductRepo := new(MockIProductRepo)
	db.InitMYSQL()
	productID := int64(802)
	value, _ := json.Marshal(productID)

	ctx := context.Background()

	productImage := entity.URLS{"http://example.com/image1.jpg", "http://example.com/image2.jpg"}
	compressedPath := entity.URLS{"path1.png"}

	mockProductRepo.On("GetProductImage", mock.Anything, productID).Return(productImage, nil)
	mockProductRepo.On("UpdateImagePath", mock.Anything, compressedPath, productID).Return(nil)
	
	err := ProcessImages(ctx, value)
	assert.NoError(t, err)
}
