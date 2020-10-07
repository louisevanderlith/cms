package core

import (
	"fmt"
	"strings"
)

type Colour struct {
	Primary    RGB
	Secondary  RGB
	Tertiary   RGB
	Shadows    RGB
	Accent     RGB
	Background RGB
}

func (c Colour) GenerateCSS() string {
	primaryrgb := fmt.Sprintf("%v, %v, %v", c.Primary.Red, c.Primary.Green, c.Primary.Blue)
	secondaryrgb := fmt.Sprintf("%v, %v, %v", c.Secondary.Red, c.Secondary.Green, c.Secondary.Blue)
	tertiaryrgb := fmt.Sprintf("%v, %v, %v", c.Tertiary.Red, c.Tertiary.Green, c.Tertiary.Blue)
	shadowsrgb := fmt.Sprintf("%v, %v, %v", c.Shadows.Red, c.Shadows.Green, c.Shadows.Blue)
	accentrgb := fmt.Sprintf("%v, %v, %v", c.Accent.Red, c.Accent.Green, c.Accent.Blue)
	backgroundrgb := fmt.Sprintf("%v, %v, %v", c.Background.Red, c.Background.Green, c.Background.Blue)

	groups := map[string]string{
		"colour":     "color: rgb(%s);",
		"background": "background: rgb(%s);",
		"bgcolour":   "background-color: rgb(%s);",
		"border":     "border-bottom: 2px solid rgb(%s);",
		"box":        "box-shadow: 0 2px 25px 0 rgba(%s, 0.1);",
	}

	final := strings.Builder{}
	for group, style := range groups {
		primstyle := fmt.Sprintf(style, primaryrgb)
		secstyle := fmt.Sprintf(style, secondaryrgb)
		terstyle := fmt.Sprintf(style, tertiaryrgb)
		shadowstyle := fmt.Sprintf(style, shadowsrgb)
		accentstyle := fmt.Sprintf(style, accentrgb)
		backstyle := fmt.Sprintf(style, backgroundrgb)
		final.WriteString(fmt.Sprintf(".%sPrimary{%s}", group, primstyle))
		final.WriteString(fmt.Sprintf(".%sSecondary{%s}", group, secstyle))
		final.WriteString(fmt.Sprintf(".%sTertiary{%s}", group, terstyle))
		final.WriteString(fmt.Sprintf(".%sShadows{%s}", group, shadowstyle))
		final.WriteString(fmt.Sprintf(".%sAccent{%s}", group, accentstyle))
		final.WriteString(fmt.Sprintf(".%sBackground{%s}", group, backstyle))
	}

	return final.String()
}

type RGB struct {
	Red   int
	Green int
	Blue  int
}
