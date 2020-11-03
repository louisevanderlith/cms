package handles

import (
	"github.com/coreos/go-oidc"
	"github.com/louisevanderlith/droxolite/mix"
	"html/template"
	"log"
	"net/http"
)

func Index(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Index", tmpl, "./views/index.html")
	pge.AddMenu(FullMenu())
	pge.AddModifier(mix.EndpointMod(Endpoints))
	pge.AddModifier(mix.IdentityMod(CredConfig.ClientID))
	pge.AddModifier(ThemeContentMod())
	return func(w http.ResponseWriter, r *http.Request) {
		tknVal := r.Context().Value("IDToken")
		if tknVal == nil {
			http.Error(w, "", http.StatusUnauthorized)
			return
		}

		idToken := tknVal.(*oidc.IDToken)
		claims := make(map[string]interface{})
		err := idToken.Claims(&claims)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = mix.Write(w, pge.Create(r, claims))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
