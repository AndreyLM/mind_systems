package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/andreylm/mind_systems/pkg/cache"
	"github.com/andreylm/mind_systems/pkg/ent"

	"github.com/gorilla/mux"
)

// URLResolver - shortener handler
func URLResolver(db *ent.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		urlToken := vars["urlToken"]

		if url := getURL(urlToken, db); url != "" {
			http.Redirect(w, r, url, http.StatusFound)
			return
		}

		http.NotFound(w, r)
	}
}

func getURL(urlToken string, db *ent.Client) string {
	if url := cache.Storage.Get(urlToken); url != "" {
		return url
	}

	ctx := context.WithValue(context.Background(), contextURLShortened, urlToken)
	t, err := QueryToken(ctx, db)
	if err != nil {
		log.Println(err)
		return ""
	}

	return t.URL
}
