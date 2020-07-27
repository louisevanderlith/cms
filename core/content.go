package core

import "github.com/louisevanderlith/husk"

type Content struct {
	Profile  string
	Language string //en,af, en-US
	Banner   Banner
	SectionA Section
	SectionB Section
	Info     Information
}

func (o Content) Valid() error {
	return husk.ValidateStruct(o)
}

func (o Content) Create() (husk.Recorder, error) {
	defer ctx.Content.Save()
	return ctx.Content.Create(o)
}

func (o Content) Update(key husk.Key) error {
	obj, err := ctx.Content.FindByKey(key)

	if err != nil {
		return err
	}

	err = obj.Set(o)

	if err != nil {
		return err
	}

	err = ctx.Content.Update(obj)

	if err != nil {
		return err
	}

	return ctx.Content.Save()
}
