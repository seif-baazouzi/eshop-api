package utils

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"gitlab.com/seif-projects/e-shop/api/src/models"
)

const MIX_USER_RATE_COUNT_FOR_CACHING = 1000

func GetShopsRating(conn *sql.DB, redisClient redis.Conn, shopsList []models.Shop) error {
	for index := range shopsList {
		shopRateKey := shopsList[index].ShopName + "Rate"

		// check the cache
		res, err := redisClient.Do("GET", shopRateKey)

		if err != nil {
			return err
		}

		if res != nil {
			resStr := fmt.Sprintf("%s", res)
			rate, _ := strconv.Atoi(resStr)
			shopsList[index].ShopRate = uint64(rate)

			continue
		}

		// get the rate
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
			var rate uint64

			if count == 0 {
				rate = 0
			} else {
				rate = sum / count
			}

			shopsList[index].ShopRate = rate

			if count > MIX_USER_RATE_COUNT_FOR_CACHING {
				redisClient.Do("SET", shopRateKey, rate, "EX", "60")
			}
		}
	}

	return nil
}
