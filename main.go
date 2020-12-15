package main

import (
	"flag"
	"net/http"
	"time"

	"github.com/louisevanderlith/cms/handles"
)

func main() {
	host := flag.String("host", "http://127.0.0.1:8107", "This application's URL")
	clientId := flag.String("client", "cms", "Client ID which will be used to verify this instance")
	clientSecrt := flag.String("secret", "secret", "Client Secret which will be used to authenticate this instance")
	issuer := flag.String("issuer", "http://127.0.0.1:8080/auth/realms/mango", "OIDC Provider's URL")
	theme := flag.String("theme", "http://127.0.0.1:8093", "Theme URL")
	folio := flag.String("folio", "http://127.0.0.1:8090", "Folio URL")
	artifact := flag.String("artifact", "http://127.0.0.1:8082", "Artifact URL")
	blog := flag.String("blog", "http://127.0.0.1:8102", "Blog URL")
	comms := flag.String("comms", "http://127.0.0.1:8085", "Comms URL")
	comment := flag.String("comment", "http://127.0.0.1:8084", "Comment URL")
	flag.Parse()

	ends := map[string]string{
		"issuer":   *issuer,
		"theme":    *theme,
		"folio":    *folio,
		"artifact": *artifact,
		"blog":     *blog,
		"comms":    *comms,
		"comment":  *comment,
	}

	srvr := &http.Server{
		ReadTimeout:  time.Second * 15,
		WriteTimeout: time.Second * 15,
		Addr:         ":8107",
		Handler:      handles.SetupRoutes(*host, *clientId, *clientSecrt, ends),
	}

	err := srvr.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
