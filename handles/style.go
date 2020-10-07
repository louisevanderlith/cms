package handles

import (
	"fmt"
	"github.com/louisevanderlith/cms/core"
	"github.com/louisevanderlith/droxolite/drx"
	"github.com/louisevanderlith/droxolite/mix"
	"log"
	"net/http"
	"strings"
)

func ProfileColour(w http.ResponseWriter, r *http.Request) {
	prf := drx.FindParam(r, "profile")

	rec, err := core.GetDisplay(prf)

	if err != nil {
		log.Println("GetDisplay Error", err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	content := rec.GetValue().(core.Content)
	colour := content.Colour.GenerateCSS()

	name := fmt.Sprintf("%s.css", prf)
	err = mix.Write(w, mix.Octet(name, strings.NewReader(colour)))

	if err != nil {
		log.Println("Serve Error", err)
	}
}
