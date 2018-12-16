package api

import (
	"fmt"
	"github.com/badoux/goscraper"
	"golang.org/x/net/context"
	"log"
)

// Server represents the gRPC server
type Server struct {
}

// ParseURL loads the web pages and returns meta data
func (s *Server) ParseURL(ctx context.Context, in *ParseRequest) (response *ParsedURLResponse, err error) {
	log.Printf("Parsing %s", in.Url)
	metaData, err := parseUrlAndReturnData(ctx, in.Url)
	return &metaData, err
}

func parseUrlAndReturnData(ctx context.Context, url string) (md ParsedURLResponse, err error) {
	s, err := goscraper.Scrape(url, 2)

	metaData := ParsedURLResponse{}

	if err != nil {
		errorMsg := fmt.Sprintf("Could not parse %s", url)
		log.Fatal(ctx, errorMsg, err)
		return metaData, err
	}

	metaData.Icon = s.Preview.Icon
	metaData.Name = s.Preview.Name
	metaData.Title = s.Preview.Title
	metaData.Description = s.Preview.Description
	metaData.Images = s.Preview.Images
	return metaData, nil
}
