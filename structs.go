package pantsuAPI

import (
	"time"
)

type PantsuPostErrors struct {
	OK        bool          `json:"ok"` // whether the request succeeded
	Errors    []string      // error messages
	AllErrors []interface{} // error objects
}

type PantsuGetErrors struct {
	Errors []string `json:"errors"` // error messages
}

// Torrent object describing an individual torrent
type Torrent struct {
	ID           int           `json:"id"`            // unique identifier for the torrent
	Name         string        `json:"name"`          // name of the torrent
	Status       int           `json:"status"`        // status code for the torrent
	Hash         string        `json:"hash"`          // hash of the torrent
	Date         time.Time     `json:"date"`          // date the torrent was uploaded
	Filesize     int           `json:"filesize"`      // size of the torrent contents in bytes
	Description  string        `json:"description"`   // HTML-formatted description of the torrent
	Comments     []string      `json:"comments"`      // comments for the torrent // TODO check type
	SubCategory  string        `json:"sub_category"`  // identifier for the subcategory of the torrent`
	Category     string        `json:"category"`      // identifier for the category of the torrent
	AniDBID      string        `json:"anidb_id"`      // identifier for the anidb entry corresponding to the torrent
	UploaderID   int           `json:"uploader_id"`   // user identifier for the user who uploaded the torrent
	UploaderName string        `json:"uploader_name"` // name of the user who uploaded the torrent
	UploaderOld  string        `json:"uploader_old"`  // name of the user who uploaded the torrent for torrents imported from old Nyaa
	WebsiteLink  string        `json:"website_link"`  // external link to relevant content for the torrent
	Languages    []string      `json:"languages"`     // languages the torrent content "provides" (audio?)
	Magnet       string        `json:"magnet"`        // magnet link for the torrent
	Torrent      string        `json:"torrent"`       // link to download the torrent file
	Seeders      int           `json:"seeders"`       // number of seeders on the torrent
	Leechers     int           `json:"leechers"`      // number of leechers on the torrent
	Completed    int           `json:"completed"`     // number of completed downloads of the torrent
	LastScrape   time.Time     `json:"last_scrape"`   // last time the torrent statistics were updated
	FileList     []interface{} `json:"file_list"`     // list of files in the torrent
}

// TorrentIndexObject describing the torrents returned by TorrentIndex()
type TorrentIndexObject struct {
	Torrents         []Torrent `json:"torrents"`         // List of torrents objects
	QueryRecordCount int       `json:"queryRecordCount"` // number of torrents returned
	TotalRecordCount int       `json:"totalRecordCount"` // total number of torrents available
	Errors           []string  `json:"errors"`           // list of errors for when the request fails
}

// TorrentSearchParameters describes the parameters to be passed to SearchTorrents()
type TorrentSearchParameters struct {
	Categories []string  `json:"c"`        // categories to search within
	Query      string    `json:"q"`        // search string (torrent name)
	Page       int       `json:"page"`     // which page number from search results is desired
	Limit      int       `json:"limit"`    // maximum number of results per page
	UserID     int       `json:"userID"`   // user identifier for the user uploading the torrent
	FromID     int       `json:"fromID"`   // DIRECT QUOTE from API docs: "Show results with torrents ID superior to this."
	Status     int       `json:"status"`   // torrent status code
	MaxAge     int       `json:"maxage"`   // search for torrents uploaded in the last <MaxAge> of days
	ToDate     time.Time `json:"toDate"`   // search for torrents uploaded since <ToDate>
	FromDate   time.Time `json:"fromDate"` // search for torrents uploaded in the last <FromDate> <date type>
	MinSize    string    `json:"minSize"`  // search for torrents of at least this size; units specified in SizeType
	MaxSize    string    `json:"maxSize"`  // search for torrents at most of this size; units specified in SizeType
	SizeType   string    `json:"sizeType"` // units for MinSize and MaxSize; b for bytes, k for kilobytes, m for megabytes, g for gigabytes
	Sort       int       `json:"sort"`     // how to sort results; 0 = id, 1 = name, 2 = date, 3 = downloads, 4 = size, 5 = seeders, 6 = leechers, 7 = completed
	Order      bool      `json:"order"`    // whether to use ascending (true) or descending (false) Order
	Languages  []string  `json:"lang"`     // search for torrents with these languages
}

// TorrentSearch response
type TorrentSearch struct {
	torrents []Torrent // torrents matching search criteria
	errors   []string  // errors for when request fails
}

// TODO: torrent head (what is that)

type UploadTorrentResponse struct {
}

type UpdateTorrentResponse struct {
}

// User object describing an individual user
type User struct {
}

type LoginResponse struct {
	OK   bool     `json:"ok"`    // whether request completed sucessfully
	Info []string `json:"infos"` // relevant information
	Data User     // user object
}
