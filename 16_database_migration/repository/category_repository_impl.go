package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_restful_api/helper"
	"go_restful_api/models/domain"
)

type CategoryRepositoryImpl struct {}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (*CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, 
category domain.Category) domain.Category {	
	sqlScript := "INSERT INTO categories(name) VALUES(?)"
	result, err := tx.ExecContext(ctx, sqlScript, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	
	return category
}

func (*CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, 
categoryId int) (domain.Category, error) {
	sqlScript := "SELECT id, name FROM categories WHERE id = ?"
	row := tx.QueryRowContext(ctx, sqlScript, categoryId)

	category := domain.Category{}
	err := row.Scan(&category.Id, &category.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return category, errors.New("category is not found")
		}
		return category, err
	}

	return category, nil
}

func (*CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx, 
) []domain.Category {
	sqlScript := "SELECT id, name FROM categories"
	rows, err := tx.QueryContext(ctx, sqlScript)
	helper.PanicIfError(err)
	defer rows.Close()

	categories := []domain.Category{}
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}

func (*CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, 
category domain.Category) domain.Category {
	sqlScript := "UPDATE categories SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, sqlScript, category.Name, category.Id)
	helper.PanicIfError(err)
	
	return category
}

func (*CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, 
category domain.Category) {
	sqlScript := "DELETE FROM categories WHERE id = ?"
	_, err := tx.ExecContext(ctx, sqlScript, category.Id)
	helper.PanicIfError(err)
}
