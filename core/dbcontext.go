package core

import (
	"encoding/json"
	"github.com/louisevanderlith/husk"
	"github.com/louisevanderlith/husk/collections"
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/op"
	"github.com/louisevanderlith/husk/records"
	"os"
	"reflect"
)

type context struct {
	Content husk.Table
}

var ctx context

func CreateContext() {
	ctx = context{
		Content: husk.NewTable(Content{}),
	}

	seed()
}

func seed() {
	contents, err := contentSeeds()

	if err != nil {
		panic(err)
	}

	err = ctx.Content.Seed(contents)

	if err != nil {
		panic(err)
	}
}

func contentSeeds() (collections.Enumerable, error) {
	f, err := os.Open("db/contents.seed.json")

	if err != nil {
		return nil, err
	}

	var items []Content
	dec := json.NewDecoder(f)
	err = dec.Decode(&items)

	if err != nil {
		return nil, err
	}

	return collections.ReadOnlyList(reflect.ValueOf(items)), nil
}

func Shutdown() {
	ctx.Content.Save()
}

func GetContent(key hsk.Key) (Content, error) {
	rec, err := ctx.Content.FindByKey(key)

	if err != nil {
		return Content{}, err
	}

	return rec.Data().(Content), nil
}

func GetAllContent(page, size int) (records.Page, error) {
	return ctx.Content.Find(page, size, op.Everything())
}

func GetDisplay(profile string) (hsk.Record, error) {
	return ctx.Content.FindFirst(byProfile(profile))
}
