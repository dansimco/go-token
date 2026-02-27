package color

type Mode struct {
	Name  string
	Roles []Role
}

func NewMode(name string) Mode {
	mode := Mode{
		Name: name,
	}
	return mode
}

func (m *Mode) AddRole(role Role) {
	m.Roles = append(m.Roles, role)
}

func (m *Mode) ToCSS() string {
	css := ""
	for _, role := range m.Roles {
		css += role.toCSS("--" + m.Name)
	}
	return css
}
