package repository

import (
	"context"
	"database/sql"
	"errors"
	"message-queue-system/domain/dto/request"
	"message-queue-system/domain/entity"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Masterminds/squirrel"
)

type IProductRepo interface {
	InsertProduct(ctx context.Context, req request.InsertProduct) (int64, error)
	GetProductImage(ctx context.Context, productID int64) (entity.URLS, error) 
	UpdateImagePath(ctx context.Context, path entity.URLS, productID int64) (error)
	UserExists(ctx context.Context, userID int64) (int, error)
}

type ProductRepo struct {
	Db *sql.DB
}

func NewProductRepo(Db *sql.DB) IProductRepo {
	return &ProductRepo{Db}
}

func (prepo *ProductRepo)  InsertProduct(ctx context.Context, req request.InsertProduct) (int64, error) {
	if len(req.ProductName)==0 {
		return 0, errors.New("product name empty")
	}
	qBuilder := squirrel.Insert("products").Columns("name","description","images","price").
							Values(req.ProductName, req.ProductDescription, req.ProductImages, req.ProductPrice)
	query, qargs, err := qBuilder.ToSql()
	if err != nil {
		return 0, errors.New("error generating query")
	}
	conn, err := prepo.Db.Conn(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	res, err := conn.ExecContext(ctx, query, qargs...)
	if err != nil {
		return 0, err
	}
	productID, err := res.LastInsertId() 
	if err != nil {
		return 0, err
	}
	return productID, nil
}

func (prepo *ProductRepo) GetProductImage(ctx context.Context, productID int64) (entity.URLS, error) {
	qBuilder := squirrel.Select("images").From("products").
						Where(squirrel.Eq{"id": productID})
	query, qargs, err := qBuilder.ToSql()
	if err != nil {
		return nil, errors.New("error generating query")
	}
	conn, err := prepo.Db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	row := conn.QueryRowContext(ctx,query,qargs...)
	var productImage entity.URLS
	err = row.Scan(&productImage)
	if err != nil {
		return nil, err
	}
	return productImage, nil
}

func (prepo *ProductRepo)  UpdateImagePath(ctx context.Context, path entity.URLS, productID int64) (error) {
	qBuilder := squirrel.Update("products").Set("compressed_product_images",path).Where(squirrel.Eq{"id": productID})
	query, qargs, err := qBuilder.ToSql()
	if err != nil {
		return errors.New("error generating query")
	}
	conn, err := prepo.Db.Conn(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	_, err = conn.ExecContext(ctx, query, qargs...)
	if err != nil {
		return err
	}
	return nil
}

func (prepo *ProductRepo) UserExists(ctx context.Context, userID int64) (int, error) {
	qBuilder := squirrel.Select("count(1)").From("users").
						Where(squirrel.Eq{"id": userID})
	query, qargs, err := qBuilder.ToSql()
	if err != nil {
		return 0, errors.New("error generating query")
	}
	conn, err := prepo.Db.Conn(ctx)
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	row := conn.QueryRowContext(ctx,query,qargs...)
	var count int
	err = row.Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}