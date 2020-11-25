package handles

import (
	"fmt"
	"github.com/louisevanderlith/blog/api"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"golang.org/x/oauth2"
	"html/template"
	"log"
	"net/http"
)

func GetArticles(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Articles", tmpl, "./views/articles.html")
	pge.AddMenu(FullMenu())
	pge.ChangeTitle("Articles")
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)

		clnt := AuthConfig.Client(r.Context(), &tkn)
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
	pge.ChangeTitle("Articles")
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
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
	pge := mix.PreparePage("Articles View", tmpl, "./views/articleview.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
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

		pge.ChangeTitle(fmt.Sprintf("View Article - %s", result.Title))
		err = mix.Write(w, pge.Create(r, result))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
