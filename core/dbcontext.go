package core

import (
	"github.com/louisevanderlith/husk"
)

type context struct {
	Content husk.Tabler
}

var ctx context

func CreateContext() {
	ctx = context{
		Content: husk.NewTable(Content{}),
	}

	seed()
}

func seed() {
	err := ctx.Content.Seed("db/contents.seed.json")

	if err != nil {
		panic(err)
	}

	err = ctx.Content.Save()

	if err != nil {
		panic(err)
	}
}

func Shutdown() {
	ctx.Content.Save()
}

func GetContent(key husk.Key) (Content, error) {
	rec, err := ctx.Content.FindByKey(key)

	if err != nil {
		return Content{}, err
	}

	return rec.Data().(Content), nil
}

func GetAllContent(page, size int) (husk.Collection, error) {
	return ctx.Content.Find(page, size, husk.Everything())
}

func GetDisplay(profile string) (husk.Recorder, error) {
	return ctx.Content.FindFirst(byProfile(profile))
}
