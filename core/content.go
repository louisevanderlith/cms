package core

import (
	"github.com/louisevanderlith/husk/hsk"
	"github.com/louisevanderlith/husk/validation"
)

type Content struct {
	Profile  string
	Language string //en,af, en-US
	Banner   Banner
	SectionA Section
	SectionB Section
	Info     Information
	Colour   Colour
}

func (o Content) Valid() error {
	return validation.Struct(o)
}

func (o Content) Create() (hsk.Key, error) {
	return ctx.Content.Create(o)
}

func (o Content) Update(key hsk.Key) error {
	return ctx.Content.Update(key, o)
}
