package usecase

import (
	"context"
	"encoding/json"
	"errors"
	productRepo "message-queue-system/business/product/repository"
	"message-queue-system/clients/kafka"
	"message-queue-system/domain/dto/request"
)

type IProductUC interface {
	InsertProduct(ctx context.Context,req request.InsertProduct) error
}

type ProductUC struct {
	repo productRepo.IProductRepo
}

func NewProductUC(r productRepo.IProductRepo) IProductUC {
	return &ProductUC{repo: r}
}

const TopicProductID = "topic.product.id"

func 	(pUC *ProductUC) InsertProduct(ctx context.Context,req request.InsertProduct) error {
	userExists, err := pUC.repo.UserExists(ctx, req.UserID)
	if err != nil {
		return err
	}
	if userExists == 0 {
		return errors.New("user with userid does not exists")
	}
	productID, err := pUC.repo.InsertProduct(ctx,req)
	if err != nil {
		return err
	}
	value, err := json.Marshal(productID)
	if err != nil {
		return err
	}
	err = kafka.Produce(TopicProductID, value)
	if err != nil {
		return err
	}
	return nil
}
