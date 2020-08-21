package handles

import (
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/kong"
	"github.com/rs/cors"
	"net/http"
)

func SetupRoutes(scrt, securityUrl, managerUrl string) http.Handler {
	r := mux.NewRouter()
	ins := kong.NewResourceInspector(http.DefaultClient, securityUrl, managerUrl)
	cnt := ins.Middleware("cms.content.view", scrt, DisplayContent)
	r.HandleFunc("/display", cnt).Methods(http.MethodGet)

	get := ins.Middleware("cms.content.search", scrt, GetContent)
	r.HandleFunc("/content", get).Methods(http.MethodGet)

	view := ins.Middleware("cms.content.view", scrt, ViewContent)
	r.HandleFunc("/content/{key:[0-9]+\\x60[0-9]+}", view).Methods(http.MethodGet)

	srch := ins.Middleware("cms.content.search", scrt, SearchContent)
	r.HandleFunc("/content/{pagesize:[A-Z][0-9]+}", srch).Methods(http.MethodGet)
	r.HandleFunc("/content/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", srch).Methods(http.MethodGet)

	create := ins.Middleware("cms.content.create", scrt, CreateContent)
	r.HandleFunc("/content", create).Methods(http.MethodPost)

	update := ins.Middleware("cms.content.update", scrt, UpdateContent)
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
