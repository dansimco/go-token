package color

type Mode struct {
	Name  string
	Roles []*Role
}

func NewMode(name string) Mode {
	mode := Mode{
		Name: name,
	}
	return mode
}

func (m *Mode) AddRole(name string) *Role {
	role := NewRole(name)
	m.Roles = append(m.Roles, role)
	return role
}

func (m *Mode) ToCSS(prefix string) string {
	css := ""
	varPrefix := "  --"
	if prefix != "" {
		varPrefix += prefix + "-"
	}
	for _, role := range m.Roles {
		css += role.toCSS(varPrefix + m.Name)
	}
	return css
}

func (m *Mode) ToCSSClass(prefix string) string {
	// Build class name
	className := "."
	if prefix != "" {
		className += prefix + "-"
	}
	className += m.Name + " {\n"
	css := className

	for _, role := range m.Roles {
		// Map each state of the role to generic variables
		for _, state := range role.States {
			var genericVar string
			var modeVar string

			// Build variable names with proper prefix handling
			varPrefix := "--"
			if prefix != "" {
				varPrefix += prefix + "-"
			}

			if state.Name == "default" {
				genericVar = "  " + varPrefix + role.Name
				modeVar = "var(" + varPrefix + m.Name + "-" + role.Name + ")"
			} else {
				genericVar = "  " + varPrefix + role.Name + "-" + state.Name
				modeVar = "var(" + varPrefix + m.Name + "-" + role.Name + "-" + state.Name + ")"
			}

			css += genericVar + ": " + modeVar + ";\n"
		}
	}

	css += "}\n"
	return css
}
