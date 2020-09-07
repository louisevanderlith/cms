package core

import (
	"github.com/louisevanderlith/husk/keys"
)

type Banner struct {
	Background keys.TimeKey
	Image      keys.TimeKey `hsk:"null"`
	Heading    string
	Subtitle   string
}
