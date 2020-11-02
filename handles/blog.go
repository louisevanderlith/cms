package handles

import (
	"github.com/louisevanderlith/blog/api"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"html/template"
	"log"
	"net/http"
)

func GetArticles(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Articles", tmpl, "./views/articles.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		clnt := CredConfig.Client(r.Context())
		result, err := api.FetchLatestArticles(clnt, Endpoints["blog"], "A10")

		if err != nil {
			log.Println("Fetch Articles Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchArticles(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Articles", tmpl, "./views/articles.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		clnt := CredConfig.Client(r.Context())
		result, err := api.FetchLatestArticles(clnt, Endpoints["blog"], drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch Articles Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewArticle(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Articles View", tmpl, "./views/articlesview.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		clnt := CredConfig.Client(r.Context())
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		result, err := api.FetchArticle(clnt, Endpoints["blog"], key)

		if err != nil {
			log.Println("Fetch Article Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
