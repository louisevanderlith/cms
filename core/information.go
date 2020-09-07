package core

type Information struct {
	Heading string `hsk:"size(50)"`
	Text    string `hsk:"size(512)"`
	Blocks  []SimpleBlock
}

type SimpleBlock struct {
	Icon string
	Text string
}
