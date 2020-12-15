package handles

import (
	"github.com/louisevanderlith/blog/api"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk/keys"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func GetArticles(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)

		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchLatestArticles(clnt, Endpoints["blog"], "A10")

		if err != nil {
			log.Println("Fetch Articles Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, fact.Create(r, "Articles", "./views/articles.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func SearchArticles(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		data, err := api.FetchLatestArticles(clnt, Endpoints["blog"], drx.FindParam(r, "pagesize"))

		if err != nil {
			log.Println("Fetch Articles Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		err = mix.Write(w, fact.Create(r, "Articles", "./views/articles.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}

func ViewArticle(fact mix.MixerFactory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tkn := r.Context().Value("Token").(oauth2.Token)
		clnt := AuthConfig.Client(r.Context(), &tkn)
		key, err := keys.ParseKey(drx.FindParam(r, "key"))

		if err != nil {
			log.Println("Parse Error", err)
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		data, err := api.FetchArticle(clnt, Endpoints["blog"], key)

		if err != nil {
			log.Println("Fetch Article Error", err)
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		//fac.ChangeTitle(fmt.Sprintf("View Article - %s", result.Title))
		err = mix.Write(w, fact.Create(r, "Article View", "./views/articleview.html", mix.NewDataBag(data)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
