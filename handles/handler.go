package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, secureUrl string) http.Handler {
	r := mux.NewRouter()

	cnt := kong.ResourceMiddleware("cms.content.view", scrt, secureUrl, DisplayContent)
	r.HandleFunc("/display", cnt).Methods(http.MethodGet)

	get := kong.ResourceMiddleware("cms.content.search", scrt, secureUrl, GetContent)
	r.HandleFunc("/content", get).Methods(http.MethodGet)

	view := kong.ResourceMiddleware("cms.content.view", scrt, secureUrl, ViewContent)
	r.HandleFunc("/content/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	srch := kong.ResourceMiddleware("cms.content.search", scrt, secureUrl, SearchContent)
	r.HandleFunc("/content/{pagesize:[A-Z][0-9]+}", srch).Methods(http.MethodGet)
	r.HandleFunc("/content/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srch).Methods(http.MethodGet)

	create := kong.ResourceMiddleware("cms.content.create", scrt, secureUrl, CreateContent)
	r.HandleFunc("/content", create).Methods(http.MethodPost)

	update := kong.ResourceMiddleware("cms.content.update", scrt, secureUrl, UpdateContent)
	r.HandleFunc("/content", update).Methods(http.MethodPut)

	lst, err := kong.Whitelist(http.DefaultClient, secureUrl, "cms.content.view", scrt)

	if err != nil {
		panic(err)
	}

	corsOpts := cors.New(cors.Options{
		AllowedOrigins: lst,
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodOptions,
			http.MethodHead,
		},
		AllowCredentials: true,
		AllowedHeaders: []string{
			"*", //or you can your header key values which you are using in your application
		},
	})

	return corsOpts.Handler(r)
}
