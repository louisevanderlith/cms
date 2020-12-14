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
		Scopes:       []string{oidc.ScopeOpenID, oidc.ScopeOfflineAccess, "upload-artifact", "folio", "blog-view", "blog-save"},
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

	lock := open.NewHybridLock(provider, credConfig, AuthConfig)

	r.HandleFunc("/login", lock.Login).Methods(http.MethodGet)
	r.HandleFunc("/callback", lock.Callback).Methods(http.MethodGet)
	r.HandleFunc("/logout", lock.Logout).Methods(http.MethodGet)
	r.HandleFunc("/refresh", lock.Refresh).Methods(http.MethodGet)

	fact := mix.NewPageFactory(tmpl)
	fact.AddMenu(FullMenu())
	fact.AddModifier(mix.EndpointMod(Endpoints))
	fact.AddModifier(mix.IdentityMod(AuthConfig.ClientID))
	fact.AddModifier(ThemeContentMod())

	r.Handle("/", lock.Protect(lock.Lock(Index(fact)))).Methods(http.MethodGet)

	rcontent := r.PathPrefix("/content").Subrouter()
	rcontent.Handle("", GetAllContent(fact)).Methods(http.MethodGet)
	rcontent.Handle("/{pagesize:[A-Z][0-9]+}", SearchContent(fact)).Methods(http.MethodGet)
	rcontent.Handle("/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", SearchContent(fact)).Methods(http.MethodGet)
	rcontent.Handle("/{key:[0-9]+\\x60[0-9]+}", ViewContent(fact)).Methods(http.MethodGet)
	rcontent.Use(lock.Protect)
	rcontent.Use(lock.Lock)

	rarticle := r.PathPrefix("/articles").Subrouter()
	rarticle.Handle("", GetArticles(fact)).Methods(http.MethodGet)
	rarticle.Handle("/{pagesize:[A-Z][0-9]+}", SearchArticles(fact)).Methods(http.MethodGet)
	rarticle.Handle("/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", SearchArticles(fact)).Methods(http.MethodGet)
	rarticle.Handle("/{key:[0-9]+\\x60[0-9]+}", ViewArticle(fact)).Methods(http.MethodGet)
	rarticle.Use(lock.Protect)
	rarticle.Use(lock.Lock)

	rcomms := r.PathPrefix("/comms").Subrouter()
	rcomms.Handle("", GetMessages(fact)).Methods(http.MethodGet)
	rcomms.Handle("/{pagesize:[A-Z][0-9]+}", SearchMessages(fact)).Methods(http.MethodGet)
	rcomms.Handle("/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", SearchMessages(fact)).Methods(http.MethodGet)
	rcomms.Handle("/{key:[0-9]+\\x60[0-9]+}", ViewMessage(fact)).Methods(http.MethodGet)
	rcomms.Use(lock.Protect)
	rcomms.Use(lock.Lock)

	rupload := r.PathPrefix("/uploads").Subrouter()
	rupload.Handle("", GetUploads(fact)).Methods(http.MethodGet)
	rupload.Handle("/{pagesize:[A-Z][0-9]+}", SearchUploads(fact)).Methods(http.MethodGet)
	rupload.Handle("/{pagesize:[A-Z][0-9]+}/{hash:[a-zA-Z0-9]+={0,2}}", SearchUploads(fact)).Methods(http.MethodGet)
	rupload.Handle("/{key:[0-9]+\\x60[0-9]+}", ViewUpload(fact)).Methods(http.MethodGet)
	rupload.Use(lock.Protect)
	rupload.Use(lock.Lock)

	return r
}

func ThemeContentMod() mix.ModFunc {
	return func(b mix.Bag, r *http.Request) {
		clnt := credConfig.Client(r.Context())

		content, err := folio.FetchDisplay(clnt, Endpoints["folio"])

		if err != nil {
			log.Println("Fetch Profile Error", err)
			panic(err)
			return
		}

		b.SetValue("Folio", content)
	}
}
