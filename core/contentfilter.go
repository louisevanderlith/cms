package core

import "github.com/louisevanderlith/husk"

type contentFilter func(obj Content) bool

func (f contentFilter) Filter(obj husk.Dataer) bool {
	return f(obj.(Content))
}

func byProfile(profile string) contentFilter {
	return func(obj Content) bool {
		return obj.Profile == profile
	}
}
