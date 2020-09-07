package core

import (
	"github.com/louisevanderlith/husk/keys"
)

type Section struct {
	Heading  string  `hsk:"size(50)"`
	Text     string  `hsk:"size(512)"`
	ImageUrl string  `hsk:"null"`
	ImageKey keys.TimeKey `hsk:"null"`
}
