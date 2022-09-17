package renderer

type Color struct{}

func (c *Color) reset() string {
	return "\u001b[0m"
}

func (c *Color) Red(str string) string {
	return "\u001b[31;1m" + str + c.reset()
}

func (c *Color) Green(str string) string {
	return "\u001b[32;1m" + str + c.reset()
}

func (c *Color) Yellow(str string) string {
	return "\u001b[33;1m" + str + c.reset()
}

func (c *Color) Blue(str string) string {
	return "\u001b[34;1m" + str + c.reset()
}
