package dmmv3

import (
	"net/http"

	"github.com/syohex/dmmv3/actress"
)

type Requester interface {
	Request() (*http.Request, error)
	Validate() error
}

func SearchByActress(r Requester) ([]actress.Actress, error) {
	if err := r.Validate(); err != nil {
		return nil, err
	}

	req, err := r.Request()
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	actresses, err := actress.ParseResponse(resp.Body)
	if err != nil {
		return nil, err
	}

	return actresses, nil
}
