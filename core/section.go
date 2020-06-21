package core

import "github.com/louisevanderlith/husk"

type Section struct {
	Heading string  `hsk:"size(50)"`
	Text string `hsk:"size(512)"`
	ImageKey husk.Key `hsk:"null"`
}
