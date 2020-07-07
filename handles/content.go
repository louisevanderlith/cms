package handles

import (
	"github.com/louisevanderlith/cms/core"
	"github.com/louisevanderlith/droxolite/context"
	"github.com/louisevanderlith/droxolite/mix"
	"github.com/louisevanderlith/husk"
	"log"
	"net/http"
)

func GetContent(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	results, err := core.GetAllContent(1, 10)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(results))

	if err != nil {
		log.Println(err)
	}
}

func DisplayContent(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	ti := ctx.GetTokenInfo()
	prf := ti.GetProfile()

	log.Println("CMS:", prf)
	rec, err := core.GetDisplay(prf)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(rec.Data()))

	if err != nil {
		log.Println(err)
	}
}

func ViewContent(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	k := ctx.FindParam("key")
	key, err := husk.ParseKey(k)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := core.GetContent(key)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(rec))

	if err != nil {
		log.Println(err)
	}
}

func SearchContent(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	page, size := ctx.GetPageData()
	results, err := core.GetAllContent(page, size)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(results))

	if err != nil {
		log.Println(err)
	}
}

func CreateContent(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	var obj core.Content
	err := ctx.Body(&obj)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	rec, err := obj.Create()

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(rec))

	if err != nil {
		log.Println(err)
	}
}

func UpdateContent(w http.ResponseWriter, r *http.Request) {
	ctx := context.New(w, r)
	key, err := husk.ParseKey(ctx.FindParam("key"))

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	body := &core.Content{}
	err = ctx.Body(body)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusBadRequest)
		return
	}

	err = body.Update(key)

	if err != nil {
		log.Println(err)
		http.Error(w, "", http.StatusNotFound)
		return
	}

	err = ctx.Serve(http.StatusOK, mix.JSON(nil))

	if err != nil {
		log.Println(err)
	}
}
