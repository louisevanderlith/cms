package core

import "github.com/louisevanderlith/husk"

type Banner struct {
	Background husk.Key
	Image husk.Key `hsk:"null"`
	Heading string
	Subtitle string
}