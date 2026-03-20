package typography

import (
	"fmt"

	"github.com/dansimco/go-design-tokens/css_util"
)

type Style struct {
	Name              string
	Family            *Family
	Size              float32
	LineHeight        float32
	Tracking          float32
	Weight            string
	WeightNumber      int
	Style             string
	cssClass          string
	UseNumberedWeight bool
}

func NewTypeStyle(name string) Style {
	s := Style{
		Name: name,
	}
	return s
}

func (s *Style) SetCSSClass(className string) {
	s.cssClass = className
}

func (s *Style) SetFamily(f *Family) {
	s.Family = f
}

func (s *Style) SetSize(size float32) {
	s.Size = size
}

func (s *Style) SetLineHeight(lineHeight float32) {
	s.LineHeight = lineHeight
}

func (s *Style) SetTracking(tracking float32) {
	s.Tracking = tracking
}

func (s *Style) SetWeight(weight string) {
	s.Weight = weight
	s.UseNumberedWeight = false
}

func (s *Style) SetWeightNumber(weight int) {
	s.WeightNumber = weight
	s.UseNumberedWeight = true
}

func (s *Style) SetStyle(style string) {
	s.Style = style
}

func (s *Style) ToCSS() string {

	className := s.Name
	if s.cssClass != "" {
		className = s.cssClass
	}

	css := "." + className + " {\n"

	// Add font-family if specified
	if s.Family != nil && s.Family.Name != "" {
		css += "  font-family: \"" + s.Family.Name + "\""
		for _, fallback := range s.Family.FallbackFonts {
			css += ", \"" + fallback + "\""
		}
		css += ";\n"
	}

	// Add font-size if specified
	if s.Size > 0 {
		css += fmt.Sprintf("  font-size: %grem;\n", s.Size)
	}

	// Add line-height if specified
	if s.LineHeight > 0 {
		css += fmt.Sprintf("  line-height: %grem;\n", s.LineHeight)
	}

	// Add letter-spacing (tracking) if specified
	if s.Tracking != 0 {
		css += fmt.Sprintf("  letter-spacing: %grem;\n", s.Tracking)
	}

	// Add font-weight if specified

	if s.UseNumberedWeight && s.WeightNumber != 0 {
		css += fmt.Sprintf("  font-weight: %d;\n", s.WeightNumber)
	}

	if !s.UseNumberedWeight && s.Weight != "" {
		css += `  font-weight: ` + s.Weight + ";\n"
	}

	// Add font-style if specified
	if s.Style != "" {
		css += "  font-style: " + s.Style + ";\n"
	}

	css += "}"

	css = css_util.Format(css)
	return css
}

func (s *Style) ToCSSVars(prefix string) string {
	varPrefix := "--"
	if prefix != "" {
		varPrefix += prefix + "-"
	}
	varPrefix += s.Name

	css := ""

	// Add font-family if specified
	if s.Family != nil && s.Family.Name != "" {
		css += varPrefix + "-font-family: \"" + s.Family.Name + "\""
		for _, fallback := range s.Family.FallbackFonts {
			css += ", \"" + fallback + "\""
		}
		css += ";\n"
	}

	// Add font-size if specified
	if s.Size > 0 {
		css += fmt.Sprintf("%s-font-size: %grem;\n", varPrefix, s.Size)
	}

	// Add line-height if specified
	if s.LineHeight > 0 {
		css += fmt.Sprintf("%s-line-height: %grem;\n", varPrefix, s.LineHeight)
	}

	// Add letter-spacing (tracking) if specified
	if s.Tracking != 0 {
		css += fmt.Sprintf("%s-letter-spacing: %grem;\n", varPrefix, s.Tracking)
	}

	// Add font-weight if specified
	if s.UseNumberedWeight && s.WeightNumber != 0 {
		css += fmt.Sprintf("%s-font-weight: %d;\n", varPrefix, s.WeightNumber)
	}

	if !s.UseNumberedWeight && s.Weight != "" {
		css += varPrefix + "-font-weight: " + s.Weight + ";\n"
	}

	// Add font-style if specified
	if s.Style != "" {
		css += varPrefix + "-font-style: " + s.Style + ";\n"
	}

	return css
}
