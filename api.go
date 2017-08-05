package pantsuAPI

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

const pantsuClientTimeout = time.Second * 30
const pantsuUserAgent = "pantsuAPI"

func torrentIndex() (index TorrentIndex) {
	const url = ""

	client := http.Client{Timeout: pantsuClientTimeout}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	logErrorFatal(err)

	req.Header.Set("User-Agent", pantsuUserAgent)

	res, err := client.Do(req)
	logErrorFatal(err)

	body, err := ioutil.ReadAll(res.Body)
	logErrorFatal(err)

	err = json.Unmarshal(body, &index)
	logErrorFatal(err)

	return index
}
