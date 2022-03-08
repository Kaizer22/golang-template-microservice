package impl

import (
	"context"
	"database/sql"
	"main/logging"
	"main/model/entity"
	"main/repository"
	"main/utils"
)

func NewPgProductRepository(connection *sql.DB) repository.ProductRepository {
	return pgProductRepository{
		db: connection,
	}
}

type pgProductRepository struct {
	db *sql.DB
}

const insQ = "INSERT name, description, category_id INTO products VALUES(?, ?, ?)"

func (r pgProductRepository) InsertProducts(ctx context.Context, products []*entity.Product) error {
	err := utils.RunWithProfiler(repository.TagInsPr,
		func() error {
			tx, err := r.db.Begin()
			if err != nil {
				logging.Error("could not begin a transaction")
				return err
			}
			defer tx.Rollback()

			stmt, err := tx.PrepareContext(ctx, insQ)
			if err != nil {
				logging.Error("could not prepare a statement")
				return err
			}

			for _, product := range products {
				_, err = stmt.ExecContext(ctx,
					product.Name,
					product.Description,
					product.CategoryId,
				)
				if err != nil {
					logging.ErrorFormat("could not insert product %d %s", product.Id, product.Name)
				}
			}

			if err = tx.Commit(); err != nil {
				logging.Error("could not commit a transaction")
				return err
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}

const delQ = "DELETE FROM products WHERE id = ?"

func (r pgProductRepository) DeleteProducts(ctx context.Context, products []*entity.Product) error {
	err := utils.RunWithProfiler(repository.TagInsPr,
		func() error {
			tx, err := r.db.Begin()
			if err != nil {
				logging.Error("could not begin a transaction")
				return err
			}
			defer tx.Rollback()

			stmt, err := tx.PrepareContext(ctx, delQ)
			if err != nil {
				logging.Error("could not prepare a statement")
				return err
			}

			for _, product := range products {
				_, err = stmt.ExecContext(ctx,
					product.Id,
				)
				if err != nil {
					logging.ErrorFormat("could not delete product %d %s", product.Id, product.Name)
				}
			}

			if err = tx.Commit(); err != nil {
				logging.Error("could not commit a transaction")
				return err
			}
			return nil
		})
	if err != nil {
		return err
	}
	return nil
}

const getAllPrQ = "SELECT id, name, category_id, description FROM products"

func (r pgProductRepository) GetAllProducts(ctx context.Context) ([]*entity.Product, error) {
	var res []*entity.Product
	err := utils.RunWithProfiler(repository.TagInsPr,
		func() error {
			tx, err := r.db.Begin()
			if err != nil {
				logging.Error("could not begin a transaction")
				return err
			}
			defer tx.Rollback()

			stmt, err := tx.PrepareContext(ctx, getAllPrQ)
			if err != nil {
				logging.Error("could not prepare a statement")
				return err
			}

			result, err := stmt.QueryContext(ctx)
			if err != nil {
				logging.ErrorFormat("could not get all products")
			}
			for result.Next() {
				p := entity.Product{}
				err := result.Scan(
					&p.Id,
					&p.Name,
					&p.CategoryId,
					&p.Description,
				)
				if err != nil {
					logging.Error("could not read product")
				}
				res = append(res, &p)
			}

			if err = tx.Commit(); err != nil {
				logging.Error("could not commit a transaction")
				return err
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	return res, nil
}

const getPrByIdQ = "SELECT id, name, category_id, description FROM products WHERE id = ?"

func (r pgProductRepository) GetProductById(ctx context.Context, id int) (*entity.Product, error) {
	res := entity.Product{}
	err := utils.RunWithProfiler(repository.TagInsPr,
		func() error {
			tx, err := r.db.Begin()
			if err != nil {
				logging.Error("could not begin a transaction")
				return err
			}
			defer tx.Rollback()

			stmt, err := tx.PrepareContext(ctx, getPrByIdQ)
			if err != nil {
				logging.Error("could not prepare a statement")
				return err
			}

			row := stmt.QueryRowContext(ctx, id)

			err = row.Scan(
				res.Id,
				res.Name,
				res.CategoryId,
				res.Description,
			)
			if err != nil {
				logging.ErrorFormat("could not get product by id %d ", id)
			}

			if err = tx.Commit(); err != nil {
				logging.Error("could not commit a transaction")
				return err
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	return &res, nil
}

const getPrByQueryQ = "SELECT * FROM products p INNER JOIN categories c ON p.category_id = c.id" + "" +
	"WHERE p.product_name LIKE ?" +
	"OR p.description LIKE ?" +
	"OR c.category_name LIKE ?;"

func (r pgProductRepository) GetProductsByQuery(ctx context.Context, q string) ([]*entity.Product, error) {
	var res []*entity.Product
	err := utils.RunWithProfiler(repository.TagInsPr,
		func() error {
			tx, err := r.db.Begin()
			if err != nil {
				logging.Error("could not begin a transaction")
				return err
			}
			defer tx.Rollback()

			stmt, err := tx.PrepareContext(ctx, getPrByQueryQ)
			if err != nil {
				logging.Error("could not prepare a statement")
				return err
			}
			searchQ := "%" + q + "%"
			result, err := stmt.QueryContext(ctx, searchQ, searchQ, searchQ)
			if err != nil {
				logging.ErrorFormat("could not get products by q")
			}
			for result.Next() {
				p := entity.Product{}
				err := result.Scan(
					&p.Id,
					&p.Name,
					&p.CategoryId,
					&p.Description,
				)
				if err != nil {
					logging.Error("could not read product")
				}
				res = append(res, &p)
			}
			if err = tx.Commit(); err != nil {
				logging.Error("could not commit a transaction")
				return err
			}
			return nil
		})
	if err != nil {
		return nil, err
	}
	return res, nil
}