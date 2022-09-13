package renderer

type Color struct{}

func (c *Color) Red(str string) string {
	return "\u001b[31;1m" + str
}

func (c *Color) Green(str string) string {
	return "\u001b[32;1m" + str
}

func (c *Color) Yellow(str string) string {
	return "\u001b[33;1m" + str
}

func (c *Color) Blue(str string) string {
	return "\u001b[34;1m" + str
}
