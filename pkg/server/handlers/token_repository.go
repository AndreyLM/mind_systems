package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/andreylm/mind_systems/pkg/ent/token"

	"github.com/andreylm/mind_systems/pkg/ent"
)

type ctxVal string

var (
	contextURL          = ctxVal("url")
	contextURLShortened = ctxVal("urlShortened")
)

// SaveToken - saves url with shortened url
func SaveToken(ctx context.Context, client *ent.Client) (*ent.Token, error) {
	cToken := client.Token.Create()
	if url, ok := ctx.Value(contextURL).(string); ok {
		log.Println(url)
		cToken.SetURL(url)
	}
	if urlShortened, ok := ctx.Value(contextURLShortened).(string); ok {
		log.Println(urlShortened)
		cToken.SetURLShortened(urlShortened)
	}
	t, err := cToken.Save(ctx)
	log.Println(t)
	if err != nil {
		return nil, fmt.Errorf("failed creating token: %v", err)
	}

	return t, nil
}

// QueryToken - gets token
func QueryToken(ctx context.Context, client *ent.Client) (*ent.Token, error) {
	urlShortened, ok := ctx.Value(contextURLShortened).(string)
	if !ok {
		return nil, fmt.Errorf("please set shortened url for query")
	}
	log.Println(urlShortened)
	t, err := client.Token.Query().Where(token.URLShortenedEQ(urlShortened)).Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying token: %v", err)
	}

	return t, nil
}
