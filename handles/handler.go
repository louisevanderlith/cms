package handles

import (
	"github.com/coreos/go-oidc"
	"github.com/gorilla/mux"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/menu"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/droxolite/open"
	folio "github.com/louisevanderlith/folio/api"
	"github.com/louisevanderlith/theme/api"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"log"
	"net/http"
)

var (
	AuthConfig *oauth2.Config
	credConfig *clientcredentials.Config
	Endpoints  map[string]string
)

func FullMenu() *menu.Menu {
	m := menu.NewMenu()

	m.AddItem(menu.NewItem("b", "/articles", "Blog", nil))
	m.AddItem(menu.NewItem("a", "/comms", "Messages", nil))
	m.AddItem(menu.NewItem("e", "/content", "Content Management", nil))
	m.AddItem(menu.NewItem("f", "/heroes", "Heroes", nil))
	m.AddItem(menu.NewItem("c", "/uploads", "Uploads & Media", nil))

	return m
}

func SetupRoutes(host, clientId, clientSecret string, endpoints map[string]string) http.Handler {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, endpoints["issuer"])

	if err != nil {
		panic(err)
	}

	Endpoints = endpoints

	AuthConfig = &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  host + "/callback",
		Scopes:       []string{oidc.ScopeOpenID, "upload-artifact", "folio", "blog-view", "blog-save"},
	}

	credConfig = &clientcredentials.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		TokenURL:     provider.Endpoint().TokenURL,
		Scopes:       []string{oidc.ScopeOpenID, "theme", "upload-artifact", "folio"},
	}

	err = api.UpdateTemplate(credConfig.Client(ctx), endpoints["theme"])

	if err != nil {
		panic(err)
	}

	tmpl, err := drx.LoadTemplate("./views")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	distPath := http.FileSystem(http.Dir("dist/"))
	fs := http.FileServer(distPath)
	r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", fs))

	lock := open.NewUILock(provider, AuthConfig)

	r.HandleFunc("/login", lock.Login).Methods(http.MethodGet)
	r.HandleFunc("/callback", lock.Callback).Methods(http.MethodGet)

	r.Handle("/", lock.Middleware(Index(tmpl))).Methods(http.MethodGet)

	r.Handle("/content", lock.Middleware(GetAllContent(tmpl))).Methods(http.MethodGet)
	r.Handle("/content/{pagesize:[A-Z][0-9]+}", lock.Middleware(SearchContent(tmpl))).Methods(http.MethodGet)
	r.Handle("/content/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", lock.Middleware(SearchContent(tmpl))).Methods(http.MethodGet)
	r.Handle("/content/{key:[0-9]+\\x60[0-9]+}", lock.Middleware(ViewContent(tmpl))).Methods(http.MethodGet)

	r.Handle("/articles", lock.Middleware(GetArticles(tmpl))).Methods(http.MethodGet)
	r.Handle("/articles/{pagesize:[A-Z][0-9]+}", lock.Middleware(SearchArticles(tmpl))).Methods(http.MethodGet)
	r.Handle("/articles/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", lock.Middleware(SearchArticles(tmpl))).Methods(http.MethodGet)
	r.Handle("/articles/{key:[0-9]+\\x60[0-9]+}", lock.Middleware(ViewArticle(tmpl))).Methods(http.MethodGet)

	r.Handle("/comms", lock.Middleware(GetMessages(tmpl))).Methods(http.MethodGet)
	r.Handle("/comms/{pagesize:[A-Z][0-9]+}", lock.Middleware(SearchMessages(tmpl))).Methods(http.MethodGet)
	r.Handle("/comms/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", lock.Middleware(SearchMessages(tmpl))).Methods(http.MethodGet)
	r.Handle("/comms/{key:[0-9]+\\x60[0-9]+}", lock.Middleware(ViewMessage(tmpl))).Methods(http.MethodGet)

	r.Handle("/uploads", lock.Middleware(GetUploads(tmpl))).Methods(http.MethodGet)
	r.Handle("/uploads/{pagesize:[A-Z][0-9]+}", lock.Middleware(SearchUploads(tmpl))).Methods(http.MethodGet)
	r.Handle("/uploads/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", lock.Middleware(SearchUploads(tmpl))).Methods(http.MethodGet)
	r.Handle("/uploads/{key:[0-9]+\\x60[0-9]+}", lock.Middleware(ViewUpload(tmpl))).Methods(http.MethodGet)

	return r
}

func ThemeContentMod() mix.ModFunc {
	return func(f mix.MixerFactory, r *http.Request) {
		clnt := credConfig.Client(r.Context())

		content, err := folio.FetchDisplay(clnt, Endpoints["folio"])

		if err != nil {
			log.Println("Fetch Profile Error", err)
			panic(err)
			return
		}

		f.SetValue("Folio", content)
	}
}
