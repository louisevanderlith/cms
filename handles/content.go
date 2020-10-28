package handles

import (
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/folio/api"
	"github.com/louisevanderlith/husk/keys"
	"html/template"
	"log"
	"net/http"
)

func GetAllContent(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Content", tmpl, "./views/content.html")
	pge.AddMenu(FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {
		clnt := CredConfig.Client(r.Context())
		result, err := api.FetchAllContent(clnt, Endpoints["folio"], "A10")

		if err != nil {
			log.Println("Fetch All Content Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchContent(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Content", tmpl, "./views/content.html")
	pge.AddMenu(FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {
		clnt := CredConfig.Client(r.Context())
		result, err := api.FetchAllContent(clnt, Endpoints["folio"], drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch All Content Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewContent(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Content View", tmpl, "./views/contentview.html")
	pge.AddMenu(FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {

		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		clnt := CredConfig.Client(r.Context())
		result, err := api.FetchContent(clnt, Endpoints["folio"], key)

		if err != nil {
			log.Println("Fetch Content Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
