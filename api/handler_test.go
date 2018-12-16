package api

import (
	"golang.org/x/net/context"
	"testing"
)

func TestParseURL(t *testing.T) {
	s := Server{}
	req := &ParseRequest{Url: "http://www.w3school.com"}
	resp, err := s.ParseURL(context.Background(), req)
	if err != nil {
		t.Errorf("ParseURL(%s) got unexpected error", req.Url)
	}
	if resp.Name == "" {
		t.Errorf("Meta data of the web page should have contained a name for the site")
	}
}
