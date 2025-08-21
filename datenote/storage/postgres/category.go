package postgres

import (
	"context"
	"datenote/datenote/storage"
)

const insertCategory = `
	INSERT INTO categories(
		title
	) VALUES(
		:title
	)
	RETURNING id;
`

const getCategory = `
	SELECT * FROM categories
	WHERE id=$1;
`
const getAllCategories = `
	SELECT * FROM categories;
`
const updateCategory = `
	UPDATE categories 
	SET title=:title
	WHERE id=:id
	RETURNING id;
`
const deleteCategory = `
	DELETE FROM categories 
	WHERE id=$1
`
const searchCategory = `
	SELECT * FROM categories
	WHERE title ILIKE '%%' || $1 || '%%';
`

func (s *Storage) CreateCategory(ctx context.Context, t storage.Category) (int64, error) {
	stmt, err := s.db.PrepareNamed(insertCategory)
	if err != nil {
		return 0, err
	}

	var id int64
	if err := stmt.Get(&id, t); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *Storage) GetCategory(ctx context.Context, id int64) (storage.Category, error) {
	var category storage.Category
	if err := s.db.Get(&category, getCategory, id); err != nil {
		return storage.Category{}, err
	}
	return category, nil
}

func (s *Storage) GetAllCategories(ctx context.Context) ([]storage.Category, error) {
	var category []storage.Category
	if err := s.db.Select(&category, getAllCategories); err != nil {
		return []storage.Category{}, err
	}
	return category, nil
}

func (s *Storage) UpdateCategory(ctx context.Context, t storage.Category) error {
	stmt, err := s.db.PrepareNamed(updateCategory)
	if err != nil {
		return err
	}

	var id int64
	if err := stmt.Get(&id, t); err != nil {
		return err
	}
	return nil
}

func (s *Storage) DeleteCategory(ctx context.Context, id int64) error {
	res := s.db.MustExec(deleteCategory, id)
	ok, err := res.RowsAffected()
	if err != nil || ok == 0 {
		return err
	}
	return nil
}

func (s *Storage) SearchCategory(ctx context.Context, sq string) ([]storage.Category, error) {
	var category []storage.Category
	if err := s.db.Select(&category, searchCategory, sq); err != nil {
		return []storage.Category{}, err
	}
	return category, nil
}
