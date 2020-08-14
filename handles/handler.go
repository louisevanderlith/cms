package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, securityUrl, managerUrl string) http.Handler {
	r := mux.NewRouter()

	cnt := kong.ResourceMiddleware(http.DefaultClient, "cms.content.view", scrt, securityUrl, managerUrl, DisplayContent)
	r.HandleFunc("/display", cnt).Methods(http.MethodGet)

	get := kong.ResourceMiddleware(http.DefaultClient, "cms.content.search", scrt, securityUrl, managerUrl, GetContent)
	r.HandleFunc("/content", get).Methods(http.MethodGet)

	view := kong.ResourceMiddleware(http.DefaultClient, "cms.content.view", scrt, securityUrl, managerUrl, ViewContent)
	r.HandleFunc("/content/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	srch := kong.ResourceMiddleware(http.DefaultClient, "cms.content.search", scrt, securityUrl, managerUrl, SearchContent)
	r.HandleFunc("/content/{pagesize:[A-Z][0-9]+}", srch).Methods(http.MethodGet)
	r.HandleFunc("/content/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srch).Methods(http.MethodGet)

	create := kong.ResourceMiddleware(http.DefaultClient, "cms.content.create", scrt, securityUrl, managerUrl, CreateContent)
	r.HandleFunc("/content", create).Methods(http.MethodPost)

	update := kong.ResourceMiddleware(http.DefaultClient, "cms.content.update", scrt, securityUrl, managerUrl, UpdateContent)
	r.HandleFunc("/content", update).Methods(http.MethodPut)

	lst, err := kong.Whitelist(http.DefaultClient, securityUrl, "cms.content.view", scrt)

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
