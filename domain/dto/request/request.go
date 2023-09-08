package request

import (
	"errors"
	"message-queue-system/domain/entity"
	"strings"
)

type InsertProduct struct {
	UserID             int64      `form:"user_id"`
	ProductName        string   `form:"name"`
	ProductDescription string   `form:"description"`
	ProductImages      entity.URLS 		`form:"images"`
	ProductPrice       float64  `form:"price"`
}

func (req *InsertProduct) Validate() error {
	if(req.UserID <= 0) {
		return errors.New("invalid userID")
	}
	req.ProductName = strings.Trim(req.ProductName, " ")
	if(len(req.ProductName) == 0) {
		return errors.New("invalid product name")
	}
	//validations on parameters can be added as per need 
	return nil
}
