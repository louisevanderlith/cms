package handles

import (
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/kong/tokens"
	"html/template"
	"log"
	"net/http"
)

func Index(tmpl *template.Template) http.HandlerFunc {
	pge := mix.PreparePage("Index", tmpl, "./views/index.html")
	pge.AddMenu(FullMenu())
	return func(w http.ResponseWriter, r *http.Request) {
		uclaims := r.Context().Value("userclaims")
		log.Println(uclaims.(tokens.UserIdentity).GetDisplayName())
		err := mix.Write(w, pge.Create(r, nil))

		if err != nil {
			log.Println("Serve Error", err)
		}
	}
}
