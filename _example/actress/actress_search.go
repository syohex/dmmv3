package main

import (
	"fmt"
	"log"
	"os"

	"github.com/syohex/dmmv3"
	"github.com/syohex/dmmv3/actress"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: actress_search.go keyword\n")
		os.Exit(1)
	}

	keyword := os.Args[1]
	req := &actress.Request{
		APIID:       os.Getenv("DMM_APIID"),
		AffiliateID: os.Getenv("DMM_AFFILIATE_ID"),
		Keyword:     keyword,
	}

	actresses, err := dmmv3.SearchByActress(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Search result: keyword=%s\n", keyword)
	for _, actress := range actresses {
		fmt.Printf("%s[%s]\n", actress.Name, actress.Ruby)
	}
}
