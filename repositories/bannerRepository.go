package repositories

import (
	"errors"
	"github.com/jacky-htg/shonet-frontend/models"
)

func GetBannerForFront() (models.Banner, error) {
	sqlWord := "SELECT * FROM `banners` WHERE `banners`.`page` = 'profile' AND `banners`.`start_date` <= NOW() AND `banners`.`end_date` >= NOW() LIMIT 1"
	var banner models.Banner
	var bannerNull models.BannerNull

	rows, err := db.Query(sqlWord)
	if err != nil {
		err = errors.New("Error @bannerRepository:GetBannerForFront #db.Query :: " + err.Error())
		return banner, err
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(
				&banner.ID,
				&banner.Banner,
				&banner.Title,
				&banner.Page,
				&bannerNull.Url,
				&banner.StartDate,
				&bannerNull.EndDate,
				&bannerNull.CreatedAt,
				&bannerNull.UpdatedAt,
			)

		if err != nil {
			err = errors.New("Error @bannerRepository:GetBannerForFront #rows.Scan :: " + err.Error())
			return banner, err
		}

		banner.Url =bannerNull.Url.String
		banner.EndDate = bannerNull.EndDate.Time
		banner.CreatedAt = bannerNull.CreatedAt.Time
		banner.UpdatedAt = bannerNull.UpdatedAt.Time
	}

	if err = rows.Err(); err !=nil {
		err = errors.New("Error @bannerRepository:GetBannerForFront #rows.Err() :: " + err.Error())
		return banner, err
	}

	return banner, nil
}