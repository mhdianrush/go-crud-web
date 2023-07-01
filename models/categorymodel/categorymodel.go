package categorymodel

import (
	"github.com/mhdianrush/go-crud-web/config"
	"github.com/mhdianrush/go-crud-web/entities"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query(`select * from categories`)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			panic(err)
		}
		categories = append(categories, category)
	}
	return categories
}

func Create(category entities.Category) bool {
	result, err := config.DB.Exec(
		`insert into categories(name, created_at, updated_at) values(?, ?, ?)`,
		category.Name, category.CreatedAt, category.UpdatedAt,
	)
	if err != nil {
		panic(err)
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return lastInsertId > 0
	// if success, the return bool is true
}

func Detail(id int) entities.Category {
	row := config.DB.QueryRow(`select id, name from categories where id = ?`, id)

	var category entities.Category
	err := row.Scan(&category.Id, &category.Name)
	if err != nil {
		panic(err.Error())
	}
	return category
}

func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec(
		`update categories set name = ?, updated_at = ? where id = ?`,
		category.Name, category.UpdatedAt, id,
	)
	if err != nil {
		panic(err)
	}
	result, err := query.RowsAffected()
	// jika data berhasil diubah, RowsAffected() akan me-return 1
	if err != nil {
		panic(err)
	}
	return result > 0
}

func Delete(id int) error {
	_, err := config.DB.Exec(`delete from categories where id = ?`, id)
	return err
}
