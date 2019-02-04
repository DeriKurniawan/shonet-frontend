package repositories

import (
	"database/sql"
	"github.com/jacky-htg/shonet-frontend/models"
)

func GetCategoryForMenu(query []string) ([]models.Category, error) {
	sqlWords :=  " SELECT * FROM `categories` WHERE `categories`.`" +query[0]+ "` = '" +query[1]+ "' ORDER BY `categories`.`title` ASC"

	return fetchCategories(db.Query(sqlWords))
}

func fetchCategories(rows *sql.Rows, err error) ([]models.Category, error) {
	var categories []models.Category

	if err != nil {
		return []models.Category{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var category models.Category
		var categoryNull models.CategoryNull

		err = rows.Scan(
				&category.ID,
				&categoryNull.ParentID,
				&category.Group,
				&category.Title,
				&categoryNull.CreatedAt,
				&categoryNull.UpdatedAt,
				&categoryNull.Icon,
			)

		if err != nil {
			return []models.Category{}, err
		}

		category.CreatedAt = categoryNull.CreatedAt.Time
		category.UpdatedAt = categoryNull.UpdatedAt.Time
		category.Icon	   = categoryNull.Icon.String
		category.ParentID  = uint(categoryNull.ParentID.Int64)

		categories = append(categories, category)
	}

	if err = rows.Err(); err != nil {
		return []models.Category{}, err
	}

	return categories, nil
}