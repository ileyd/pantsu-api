package pantsuAPI

import (
	"encoding/json"
	"gopkg.in/resty.v0"
	"io"
	"io/ioutil"
	"net/http"
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

func torrentIndex() (index TorrentIndex) {
	const url = ""
	logErrorFatal(consumeGet(url, nil, "", &index))
	return index
}
