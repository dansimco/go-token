package color

type UIColor struct {
	Light LABColor
	Dark  LABColor
}

func (c *UIColor) ToCSSLightDark() string {
	return "light-dark(" + c.Light.ToHex() + ", " + c.Dark.ToHex() + ")"
}
