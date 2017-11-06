package urlbuilder

import (
	"errors"
	"fmt"
)

type Question struct {
	Items struct {
		Link               string `json:"link"`
		Last_activity_date int64  `json:"last_activity_date"`
	} `json:"items"`
}

/*
	Make the API endpoint. It will take userID
	and page count as input

	Return the stackoverflow question API endpoint, or error
*/
func (q *Question) APIEndpoint(userID, pageCount int64) (string, error) {
	apiEndpoint := fmt.Sprintf("https://api.stackexchange.com/2.2/users/%d/questions?order=desc&sort=activity&site=stackoverflow&page=%d", userID, pageCount)

	var uid, pcount int64
	n, err := fmt.Sscanf(apiEndpoint, "https://api.stackexchange.com/2.2/users/%d/questions?order=desc&sort=activity&site=stackoverflow&page=%d", &uid, &pcount)
	if err != nil || n != 2 {
		return "", errors.New("failed to parse api endpoint")
	}
	return apiEndpoint, nil
}
