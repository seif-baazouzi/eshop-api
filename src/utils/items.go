package utils

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

func GetSingleItemRating(conn *sql.DB, redisClient redis.Conn, itemID uint) (uint64, error) {
	itemRateKey := strconv.FormatUint(uint64(itemID), 10) + "ItemRate"

	// check the cache
	res, err := redisClient.Do("GET", itemRateKey)

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
		"SELECT sum(rate) as sum, count(*) as count FROM items I, itemsRates R WHERE I.itemID = R.itemID AND R.itemID = $1",
		itemID,
	)

	if err != nil {
		return 0, err
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

		if count > MIN_USER_RATE_COUNT_FOR_CACHING {
			redisClient.Do("SET", itemRateKey, rate, "EX", "60")
		}

		return rate, nil
	}

	return 0, nil
}
