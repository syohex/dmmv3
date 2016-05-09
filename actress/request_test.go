package actress

import (
	"testing"
)

func TestValidate(t *testing.T) {
	req := &Request{
		APIID:       "foo",
		AffiliateID: "bar",
		Initial:     "わ",
		Bust:        "-100",
		Waist:       "90",
		Hip:         "60-100",
		Height:      "-200",
		Hits:        88,
		Offset:      30,
		Sort:        "birthday",
	}

	if err := req.Validate(); err != nil {
		t.Error(err)
	}
}
