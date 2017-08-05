package pantsuAPI

import (
	"encoding/json"
	"gopkg.in/resty.v0"
	"time"
)

const pantsuClientTimeout = time.Second * 30
const pantsuUserAgent = "pantsuAPI"

func consumeGet(url string, params map[string]string, token string, response interface{}) error {
	res, err := resty.R().
		SetQueryParams(params).
		SetHeader("Accept", "application/json").
		SetAuthToken(token).
		Get("/search_result")
	if err != nil {
		return err
	}

	err = json.Unmarshal(res.Body(), &response)
	if err != nil {
		return err
	}

	return nil
}

func consumePost(url string, params map[string]string, token string, response interface{}, responseError interface{}) error {
	res, err := resty.R().
		SetBody(params).
		SetResult(&response).
		SetError(&responseError).
		SetAuthToken(token).
		Post(url)
	if err != nil {
		return err
	}

	err = json.Unmarshal(res.Body(), &response)
	if err != nil {
		return err
	}

	return nil
}

// TorrentIndex fetches and returns the Pantsu torrend index
func TorrentIndex() (index TorrentIndexObject, err error) {
	const url = ""
	err = consumeGet(url, nil, "", &index)
	return index, err
}

// func searchTorrents

// func torrentInfo

// func torrentHead // what is this

// func uploadTorrent

// func updateTorrent

// UserLogin logs a user into pantsu and returns the user object for the user
func UserLogin(username, password string) (res LoginResponse, resErr PantsuPostErrors, err error) {
	const url = ""
	var params map[string]string
	params["username"] = username
	params["password"] = password

	err = consumePost(url, params, "", res, resErr)
	return res, resErr, err
}

// UserProfile fetches the user profile for the specified user
func UserProfile(id int) (res User, resErr PantsuPostErrors, err error) {
	const url = ""
	var params map[string]string
	params["id"] = string(id)

	err = consumePost(url, params, "", res, resErr)
	return res, resErr, err
}
