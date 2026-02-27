package color

type Role struct {
	Name     string
	Default  UIColor
	Hover    UIColor
	Pressed  UIColor
	Focus    UIColor
	Disabled UIColor
}

func (r *Role) toCSS(prepend string) string {
	css := ""
	if prepend == "" {
		prepend = "--"
	} else {
		prepend = prepend + "-"
	}
	css += prepend + r.Name + ": " + r.Default.ToCSSLightDark() + ";\n"
	css += prepend + r.Name + "-hover: " + r.Hover.ToCSSLightDark() + ";\n"
	css += prepend + r.Name + "-pressed: " + r.Pressed.ToCSSLightDark() + ";\n"
	css += prepend + r.Name + "-focus: " + r.Focus.ToCSSLightDark() + ";\n"
	css += prepend + r.Name + "-disabled: " + r.Disabled.ToCSSLightDark() + ";\n"
	return css
}
