package utils

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"gitlab.com/seif-projects/e-shop/api/src/models"
)

const MIN_USER_RATE_COUNT_FOR_CACHING = 1000

func GetSingleShopRating(conn *sql.DB, redisClient redis.Conn, shopName string) (uint64, error) {
	shopRateKey := shopName + "Rate"

	// check the cache
	res, err := redisClient.Do("GET", shopRateKey)

	if err != nil {
		return 0, err
	}

	if res != nil {
		resStr := fmt.Sprintf("%s", res)
		rate, _ := strconv.Atoi(resStr)
		return uint64(rate), nil
	}

	// get the rate
	rows, err := conn.Query(
		"SELECT sum(rate) as sum, count(*) as count FROM shops S, shopsRates R WHERE S.shopName = R.shopName AND R.shopName = $1",
		shopName,
	)

	if err != nil {
		return 0, err
	}

	defer rows.Close()

	if rows.Next() {
		var sum uint64
		var count uint64
		rows.Scan(&sum, &count)
		var rate uint64

		if count == 0 {
			rate = 0
		} else {
			rate = sum / count
		}

		if count > MIN_USER_RATE_COUNT_FOR_CACHING {
			redisClient.Do("SET", shopRateKey, rate, "EX", "60")
		}

		return rate, nil
	}

	return 0, nil
}

func GetShopsRating(conn *sql.DB, redisClient redis.Conn, shopsList []models.Shop) error {
	for index := range shopsList {
		rate, err := GetSingleShopRating(conn, redisClient, shopsList[index].ShopName)

		if err != nil {
			return err
		}

		shopsList[index].ShopRate = rate
	}

	return nil
}
