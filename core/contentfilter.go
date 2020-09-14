package core

import (
	"github.com/louisevanderlith/husk/hsk"
)

type contentFilter func(obj Content) bool

func (f contentFilter) Filter(obj hsk.Record) bool {
	return f(obj.GetValue().(Content))
}

func byProfile(profile string) contentFilter {
	return func(obj Content) bool {
		return obj.Profile == profile
	}
}
