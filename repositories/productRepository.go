package repositories

import (
	"database/sql"
	"errors"
	"github.com/jacky-htg/shonet-frontend/models"
	"strconv"
	"strings"
)

func GetTopProduct(flag uint) (TopProductMenu, error) {
	var result TopProductMenu
	var oriLimit = 3
	var curLimit = 0
	var whereNotIn = []string{}

	sqlWords := " SELECT `id`, `name`, `brand_id`, `category_id`, `thumbnail`, `description`, `price`, `site_url`, `must_have`, `top_product`, `created_by`, `created_at`, `updated_at`, `view`, `click` " +
		        " FROM `products` WHERE `products`.`top_product` != 0 AND `products`.`category_id` IN ( " +
		   		" SELECT `categories`.`id` FROM `categories` WHERE `categories`.`parent_id` = " +strconv.Itoa(int(flag))+
		   		" ) ORDER BY `products`.`name` DESC LIMIT " +strconv.Itoa(oriLimit)

	products, err := fetchProducts(db.Query(sqlWords))
	if err != nil {
		err = errors.New(" Error @productRepository:GetTopProduct #fetchProducts#db.Query :: " + err.Error())
		return TopProductMenu{}, err
	}

	if len(products) > 0 {
		result.Top = products
		for _, val := range products {whereNotIn = append(whereNotIn, strconv.Itoa(int(val.ID)))}
		curLimit += len(products)
	} else {
		whereNotIn = append(whereNotIn, strconv.Itoa(0))
	}

	if curLimit < oriLimit {
		sqlWords = " SELECT `id`, `name`, `brand_id`, `category_id`, `thumbnail`, `description`, `price`, `site_url`, `must_have`, `top_product`, `created_by`, `created_at`, `updated_at`, `view`, `click` " +
			       " FROM `products` WHERE `products`.`id` NOT IN ( " +
				   " " +strings.Join(whereNotIn, ", ")+ " ) " +
				   " AND `products`.`category_id` IN ( " +
				   " SELECT `id` FROM `categories` WHERE `categories`.`parent_id` = " +strconv.Itoa(int(flag))+
				   " ) ORDER BY `products`.`view`, `products`.`id` DESC  LIMIT " +strconv.Itoa(oriLimit - curLimit)

		products, err = fetchProducts(db.Query(sqlWords))
		if err != nil {
			err = errors.New(" Error @productRepository:GetTopProduct #fetchProducts#db.Query#2 :: " + err.Error())
			return TopProductMenu{}, err
		}

		result.Products = products
	}

	return result, nil
}

func fetchProducts(rows *sql.Rows, err error) ([]models.Product, error) {
	var products []models.Product

	if err != nil {
		err = errors.New(" Error @productRepository:fetchProducts #db.Query :: " + err.Error())
		return []models.Product{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var product models.Product
		var productNull models.ProductNull

		err = rows.Scan(
				&product.ID,
				&product.Name,
				&product.Brand.ID,
				&product.Category.ID,
				&product.Thumbnail,
				&productNull.Description,
				&product.Price,
				&product.SiteURL,
				&product.MustHave,
				&product.TopProduct,
				&product.CreatedBy,
				&productNull.CreatedAt,
				&productNull.UpdatedAt,
				&productNull.View,
				&productNull.Click,
			)

		if err != nil {
			err = errors.New(" Error @productRepository:fetchProducts #rows.Scan :: " + err.Error())
			return []models.Product{}, err
		}

		product.Description 	= productNull.Description.String
		product.CreatedAt 		= productNull.CreatedAt.Time
		product.UpdatedAt 		= productNull.UpdatedAt.Time
		product.View 			= int(productNull.View.Int64)
		product.Click 			= int(productNull.Click.Int64)

		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		err = errors.New(" Error @productRepository:fetchProducts #rows.Err() :: " + err.Error())
		return []models.Product{}, err
	}

	return products, nil
}