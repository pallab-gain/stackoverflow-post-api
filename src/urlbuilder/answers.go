package urlbuilder

import (
	"errors"
	"fmt"
)

type Answer struct {
	Items []struct {
		Link               string `json:"link"`
		Last_activity_date int64  `json:"last_activity_date"`
		AnswerID           int64  `json:"answer_id"`
	} `json:"items"`
}

/*
	Make the API endpoint. It will take userID
	and page count as input

	Return the stackoverflow answer API endpoint, or error
*/
func (a *Answer) APIEndpoint(userID, pageCount int64) (string, error) {
	apiEndpoint := fmt.Sprintf("https://api.stackexchange.com/2.2/users/%d/answers?order=desc&sort=activity&site=stackoverflow&page=%d", userID, pageCount)

	var uid, pcount int64
	n, err := fmt.Sscanf(apiEndpoint, "https://api.stackexchange.com/2.2/users/%d/answers?order=desc&sort=activity&site=stackoverflow&page=%d", &uid, &pcount)
	if err != nil || n != 2 {
		return "", errors.New("failed to parse api endpoint")
	}
	return apiEndpoint, nil
}

/*
	Get answer link URL from answer ID. Take answer id

	Return answer link URL, or error

*/
func (a *Answer) GetLink(answerID int64) (string, error) {
	linkURL := fmt.Sprintf("https://stackoverflow.com/a/%d", answerID)

	var aid int64
	n, err := fmt.Sscanf(linkURL, "https://stackoverflow.com/a/%d", &aid)
	if err != nil || n != 2 {
		return "", errors.New("failed to parse api endpoint")
	}
	return linkURL, nil
}
