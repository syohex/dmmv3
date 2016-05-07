package actress

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Response is response from actress searching
type Response struct {
	Request struct {
		Parameters struct {
			APIID       string `json:"api_id"`
			AffiliateID string `json:"affiliate_id"`
			Initial     string `json:"initial"`
			Keyword     string `json:"keyword"`
			Bust        string `json:"bust"`
			Waist       string `json:"waist"`
			Hip         string `json:"hip"`
			Height      string `json:"height"`
			Birthday    string `json:"birthday"`
			Hits        string `json:"hits"`
			Offset      string `json:"offset"`
			Sort        string `json:"sort"`
			Output      string `json:"output"`
		} `json:"parameters"`
	} `json:"request"`

	Result struct {
		Status        string `json:"status"`
		Result        int    `json:"result_count"`
		TotalCount    string `json:"total_count"`
		FirstPosition int    `json:"first_position"`
		Actresses     []struct {
			ID          string            `json:"id"`
			Name        string            `json:"name"`
			Ruby        string            `json:"ruby"`
			Bust        string            `json:"bust"`
			Cup         string            `json:"cup"`
			Waist       string            `json:"hip"`
			Hip         string            `json:"hip"`
			Height      string            `json:"height"`
			Birthday    string            `json:"birthday"`
			BloodType   string            `json:"blood_type"`
			Hobby       string            `json:"hobby"`
			Prefectures string            `json:"prefectures"`
			ImageURL    map[string]string `json:"imageURL"`
			ListURL     map[string]string `json:"listURL"`
		} `json:"actress"`
	} `json:"result"`
}

// Actress is actress instance
type Actress struct {
	ID         string
	Name       string
	Ruby       string
	Bust       int
	Cup        string
	Waist      int
	Hip        int
	Height     int
	Birthday   *time.Time
	BloodType  string
	Hobbies    []string
	BirthPlace string
	ImageURL   map[string]string
	ListURL    map[string]string
}

var birthdayRe = regexp.MustCompile(`^(\d+)-(\d+)-(\d+)$`)

func parseBirthday(date string) (*time.Time, error) {
	m := birthdayRe.FindStringSubmatch(date)
	if m == nil {
		return nil, fmt.Errorf("invalid date '%s'", date)
	}

	year, err := strconv.ParseInt(m[1], 10, 32)
	if err != nil {
		return nil, err
	}

	month, err := strconv.ParseInt(m[2], 10, 32)
	if err != nil {
		return nil, err
	}

	day, err := strconv.ParseInt(m[3], 10, 32)
	if err != nil {
		return nil, err
	}

	ret := time.Date(int(year), time.Month(month), int(day), 0, 0, 0, 0, time.Local)
	return &ret, nil
}

func strToInt(str *string) (int, error) {
	ret := 0
	s := *str
	if s != "" {
		v, err := strconv.Atoi(s)
		if err != nil {
			return 0, err
		}
		ret = v
	}

	return ret, nil
}

// ParseResponse parses API response and return actresses information
func ParseResponse(res io.Reader) ([]Actress, error) {
	d := json.NewDecoder(res)

	var r Response
	if err := d.Decode(&r); err != nil {
		return nil, err
	}

	// XXX Error check

	var actresses []Actress
	for _, a := range r.Result.Actresses {
		bust, err := strToInt(&a.Bust)
		if err != nil {
			return nil, err
		}

		waist, err := strToInt(&a.Waist)
		if err != nil {
			return nil, err
		}

		hip, err := strToInt(&a.Hip)
		if err != nil {
			return nil, err
		}

		height, err := strToInt(&a.Height)
		if err != nil {
			return nil, err
		}

		var birthday *time.Time
		if a.Birthday != "" {
			birthday, err = parseBirthday(a.Birthday)
			if err != nil {
				return nil, err
			}
		}

		hobbies := strings.Split(a.Hobby, "、")

		actress := Actress{
			ID:         a.ID,
			Name:       a.Name,
			Ruby:       a.Ruby,
			Bust:       bust,
			Cup:        a.Cup,
			Waist:      waist,
			Hip:        hip,
			Height:     height,
			Birthday:   birthday,
			BloodType:  a.BloodType,
			Hobbies:    hobbies,
			BirthPlace: a.Prefectures,
			ImageURL:   a.ImageURL,
			ListURL:    a.ListURL,
		}

		actresses = append(actresses, actress)
	}

	return actresses, nil
}
