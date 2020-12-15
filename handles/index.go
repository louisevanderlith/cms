package handles

import (
	"github.com/coreos/go-oidc"
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"
)

func Index(fact mix.MixerFactory) http.HandlerFunc {
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

		err = mix.Write(w, fact.Create(r, "Index", "./views/index.html", mix.NewDataBag(claims)))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
