package main

import (
	"database/sql"
	"github.com/yookoala/cachedfetcher"
	"log"
	"time"
)

func example1(host string, db *sql.DB) (resp *cachedfetcher.Response, err error) {
	url := host + "/example/1"
	c := cachedfetcher.NewSqlCache(db)
	f := cachedfetcher.New(c)
	ctx := cachedfetcher.Context{
		Str:  "example/1",
		Time: time.Now(),
	}
	resp, err = f.Get(url, ctx)
	if err != nil {
		return
	}

	// log response
	log.Printf("Host:   %s", host)
	log.Printf("URL:    %s", resp.URL)
	log.Printf("Status: %s", resp.Status)
	log.Printf("Size:   %d", resp.ContentLength)
	log.Printf("Body:   \"%s\"", string(resp.Body))
	return
}
