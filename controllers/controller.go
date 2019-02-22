package controllers

import (
	"encoding/json"
	"errors"
	"github.com/jacky-htg/shonet-frontend/config"
	"github.com/jacky-htg/shonet-frontend/libraries"
	"github.com/jacky-htg/shonet-frontend/models"
	"github.com/jacky-htg/shonet-frontend/repositories"
	"time"
)

type Page struct {
	Title       string
	Description string
	Uri			string
	Routing		string
	Data        interface{}
}

type Fashion struct {
	Categories		[]models.Category
	Brands			[]models.Brand
	TopProducts		repositories.TopProductMenu
}

type Beauty struct {
	Categories		[]models.Category
	Brands			[]models.Brand
	TopProducts		repositories.TopProductMenu
}

type Journal struct {
	Categories		[]models.Category
	HotArticles     repositories.HotArticlesMenu
}

type Menu struct {
	Fashion Fashion
	Beauty  Beauty
	Journal Journal
}


func GetMenu() (Menu, error) {

	var menu Menu
	cacheName := "_front_menu"

	if yes := libraries.RedisExists(config.GetString("database.redis.name") + cacheName); yes {
		getmenu, err := libraries.RedisGet(config.GetString("database.redis.name") + cacheName)
		if err != nil {
			err = errors.New(" Error @controller.go:GetMenu() #cache#redis#getmenu :: " + err.Error())
			return Menu{}, err
		}

		if err = json.Unmarshal(getmenu, &menu); err != nil {
			err = errors.New(" Error @controller.go:GetMenu() #cache#redis#getmenu#json.Unmarshal :: " + err.Error())
			return Menu{}, err
		}

		return menu, nil
	}

	//for fashion menus
	categories, err := repositories.GetCategoryForMenu([]string{"parent_id", "1"})
	if err != nil {
		err = errors.New(" Error @controller.go:GetMenu() #categories#GetCategoryForMenu#1 :: " + err.Error())
		return Menu{}, err
	}
	menu.Fashion.Categories = categories

	brands, err := repositories.GetTopBrands(1)
	if err != nil {
		err = errors.New(" Error @controller.go:GetMenu() #brands#GetTopBrands#1 :: " + err.Error())
		return Menu{}, err
	}
	menu.Fashion.Brands = brands

	topProducts, err := repositories.GetTopProduct(1)
	if err != nil {
		err = errors.New(" Error @controller.go:GetMenu() #topProducts#GetTopProduct#1 :: " + err.Error())
		return Menu{}, err
	}
	menu.Fashion.TopProducts = topProducts

	//for beauty
	categories, err = repositories.GetCategoryForMenu([]string{"parent_id", "2"})
	if err != nil {
		err = errors.New(" Error @controller.go:GetMenu() #categories#GetCategoryForMenu#2 :: " + err.Error())
		return Menu{}, err
	}
	menu.Beauty.Categories = categories

	brands, err = repositories.GetTopBrands(2)
	if err != nil {
		err = errors.New(" Error @controller.go:GetMenu() #brands#GetTopBrands#2 :: " + err.Error())
		return Menu{}, err
	}
	menu.Beauty.Brands = brands

	topProducts, err = repositories.GetTopProduct(2)
	if err != nil {
		err = errors.New(" Error @controller.go:GetMenu() #topProducts#GetTopProduct#2 :: " + err.Error())
		return Menu{}, err
	}
	menu.Beauty.TopProducts = topProducts

	//for Journal articles
	categories, err = repositories.GetCategoryForMenu([]string{"group", "J"})
	if err != nil {
		err = errors.New(" Error @controller.go:GetMenu() #categories#GetCategoryForMenu#J :: " + err.Error())
		return Menu{}, err
	}
	menu.Journal.Categories = categories

	hotArticles, err := repositories.GetHotArticles()
	if err != nil {
		err = errors.New(" Error @controller.go:GetMenu() #hotArticles#GetHotArticles :: " + err.Error())
		return Menu{}, err
	}
	menu.Journal.HotArticles = hotArticles

	//set ton redis
	data, err := json.Marshal(&menu)
	if err != nil {
		err = errors.New(" :: Error @controller.go:GetMenu() #data#json.Marshal#menu :: " + err.Error())
		return menu, err
	}

	if err = libraries.RedisSet(config.GetString("database.redis.name") + cacheName, data, time.Hour*24); err != nil {
		err = errors.New(" Error @controller.go:GetMenu() #RedisSet#database.redis.name :: " + err.Error())
		return Menu{}, err
	}

	return menu, nil
}
