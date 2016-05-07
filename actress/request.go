package actress

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"time"
	"strconv"
)

// Request is request parameters of actress searching
type Request struct {
	APIID       string
	AffiliateID string
	Initial     string
	Keyword     string
	Bust        string
	Waist       string
	Hip         string
	Height      string
	Birthday    *time.Time
	Hits        int
	Offset      int
	Sort        string
}

// Construct http.Request object
func (r *Request) Request() (*http.Request, error) {
	query := url.Values{}
	query.Set("api_id", r.APIID)
	query.Set("affiliate_id", r.AffiliateID)

	if r.Initial != "" {
		query.Set("initial", r.Initial)
	}
	if r.Keyword != "" {
		query.Set("keyword", r.Keyword)
	}
	if r.Bust != "" {
		query.Set("bust", r.Bust)
	}
	if r.Waist != "" {
		query.Set("waist", r.Waist)
	}
	if r.Hip != "" {
		query.Set("hip", r.Hip)
	}
	if r.Height != "" {
		query.Set("height", r.Height)
	}
	if r.Birthday != nil {
		t := r.Birthday
		b := fmt.Sprintf("%d-%0d-%0d", t.Year(), t.Month(), t.Day())
		query.Set("birthday", b)
	}
	if r.Hits != 0 {
		query.Set("hits", strconv.Itoa(r.Hits))
	}
	if r.Offset != 0 {
		query.Set("offset", strconv.Itoa(r.Offset))
	}
	if r.Sort != "" {
		query.Set("sort", r.Sort)
	}

	query.Set("output", "json")

	request, err := http.NewRequest("GET", "https://api.dmm.com/affiliate/v3/ActressSearch", nil)
	if err != nil {
		return nil, err
	}

	request.URL.RawQuery = query.Encode()

	return request, nil
}

var sizeRe = regexp.MustCompile(`^(?:-?[1-9]\d*|[1-9]\d*-[1-9]\d*)$`)
var sortRe = regexp.MustCompile(`^-?(?:name|bust|waist|hip|height|birthday)$`)

// Validate request parameters
func (r *Request) Validate() error {
	if r.APIID == "" {
		return fmt.Errorf("mandatory parameter 'APIID' is not set")
	}
	if r.AffiliateID == "" {
		return fmt.Errorf("mandatory parameter 'AffiliateID' is not set")
	}

	// XXX check initial

	if r.Bust != "" && !sizeRe.MatchString(r.Bust) {
		return fmt.Errorf("invalid 'Bust' parameter")
	}

	if r.Waist != "" && !sizeRe.MatchString(r.Waist) {
		return fmt.Errorf("invalid 'Waist' parameter")
	}
	if r.Hip != "" && !sizeRe.MatchString(r.Hip) {
		return fmt.Errorf("invalid 'Hip' parameter")

	}
	if r.Height != "" && !sizeRe.MatchString(r.Height) {
		return fmt.Errorf("invalid 'Height' parameter")
	}

	if r.Hits != 0 && (r.Hits < 0 || r.Hits > 100) {
		return fmt.Errorf("'Hits' parameter should be 0 < hits <= 100")
	}
	if r.Offset != 0 && r.Offset < 0 {
		return fmt.Errorf("'Offset' parameter should be offset > 0")
	}
	if r.Sort != "" && !sortRe.MatchString(r.Sort) {
		return fmt.Errorf("invalid 'Sort' parameter")
	}

	return nil
}
