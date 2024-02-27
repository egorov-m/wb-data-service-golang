package utils

import "time"

const threeDays = time.Hour * 24 * 3

func CheckRefreshExpirationTime(tokenExpiration time.Time) bool {
	return tokenExpiration.Add(threeDays).After(time.Now())
}
