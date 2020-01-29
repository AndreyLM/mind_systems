package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/andreylm/mind_systems/pkg/cache"
	"github.com/andreylm/mind_systems/pkg/ent"

	"github.com/andreylm/mind_systems/pkg/shortener"
)

// ShortenerForm - form for  json post-request
type ShortenerForm struct {
	URL string `json:"url"`
}

// Shortener - shortener handler
func Shortener(host string, db *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		form := ShortenerForm{}
		response := make(map[string]interface{})

		if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
			http.Error(w, "Cannot decode request body", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if form.URL == "" {
			response["status"] = http.StatusBadRequest
			json.NewEncoder(w).Encode(response)
			return
		}
		urlToken := shortener.Generate()
		response["status"] = http.StatusOK
		response["shortened_url"] = fmt.Sprintf("http://%s/%s", host, urlToken)

		json.NewEncoder(w).Encode(response)

		save(form.URL, urlToken, db)
	}
}

func save(url, shortenedURL string, dbCli *ent.Client) error {
	ctx := context.WithValue(
		context.WithValue(context.Background(), contextURL, url),
		contextURLShortened, shortenedURL,
	)
	cache.Storage.Add(shortenedURL, url)
	_, err := SaveToken(ctx, dbCli)

	return err
}
