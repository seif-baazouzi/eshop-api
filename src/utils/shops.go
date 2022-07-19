package utils

import (
	"database/sql"

	"gitlab.com/seif-projects/e-shop/api/src/models"
)

func GetShopsRating(conn *sql.DB, shopsList []models.Shop) error {
	for index := range shopsList {
		rows, err := conn.Query(
			"SELECT sum(rate) as sum, count(*) as count FROM shops S, shopsRates R WHERE S.shopName = R.shopName AND R.shopName = $1",
			shopsList[index].ShopName,
		)

		if err != nil {
			return err
		}

		if rows.Next() {
			var sum uint64
			var count uint64
			rows.Scan(&sum, &count)

			if count == 0 {
				shopsList[index].ShopRate = 0
			} else {
				shopsList[index].ShopRate = sum / count
			}
		}
	}

	return nil
}
