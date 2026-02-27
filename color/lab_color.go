package color

import (
	"math"
	"strconv"
	"strings"
)

type LABColor struct {
	L float64
	A float64
	B float64
}

func NewLAB() LABColor {
	return LABColor{}
}

func (c *LABColor) ToHex() string {
	rgb := c.ToRGB()

	// ToRGB returns values in 0-255 range, clamp and convert to int
	rInt := int(clamp(rgb[0], 0, 255) + 0.5)
	gInt := int(clamp(rgb[1], 0, 255) + 0.5)
	bInt := int(clamp(rgb[2], 0, 255) + 0.5)

	// Format as hex string with padding
	return strings.ToUpper("#" +
		padHex(strconv.FormatInt(int64(rInt), 16)) +
		padHex(strconv.FormatInt(int64(gInt), 16)) +
		padHex(strconv.FormatInt(int64(bInt), 16)))
}

func (c *LABColor) ToRGB() []float64 {
	// Convert LAB to XYZ
	x, y, z := labToXyz(c.L, c.A, c.B)

	// Convert XYZ to linear RGB
	r := x*3.2404542 + y*-1.5371385 + z*-0.4985314
	g := x*-0.9692660 + y*1.8760108 + z*0.0415560
	b := x*0.0556434 + y*-0.2040259 + z*1.0572252

	// Apply gamma correction (linear to sRGB)
	r = inverseGammaCorrect(r)
	g = inverseGammaCorrect(g)
	b = inverseGammaCorrect(b)

	// Clamp to 0-1 range and convert to 0-255
	r = clamp(r, 0, 1) * 255
	g = clamp(g, 0, 1) * 255
	b = clamp(b, 0, 1) * 255

	return []float64{r, g, b}
}

func (c *LABColor) ToOKLCH() []float64 {
	// Convert LAB to XYZ first
	x, y, z := labToXyz(c.L, c.A, c.B)

	// Convert XYZ to linear RGB
	r := x*3.2404542 + y*-1.5371385 + z*-0.4985314
	g := x*-0.9692660 + y*1.8760108 + z*0.0415560
	b := x*0.0556434 + y*-0.2040259 + z*1.0572252

	// Convert linear RGB to OKLAB
	l := 0.4122214708*r + 0.5363325363*g + 0.0514459929*b
	m := 0.2119034982*r + 0.6806995451*g + 0.1073969566*b
	s := 0.0883024619*r + 0.2817188376*g + 0.6299787005*b

	l = math.Pow(l, 1.0/3.0)
	m = math.Pow(m, 1.0/3.0)
	s = math.Pow(s, 1.0/3.0)

	okL := 0.2104542553*l + 0.7936177850*m - 0.0040720468*s
	okA := 1.9779984951*l - 2.4285922050*m + 0.4505937099*s
	okB := 0.0259040371*l + 0.7827717662*m - 0.8086757660*s

	// Convert OKLAB to OKLCH
	okC := math.Sqrt(okA*okA + okB*okB)
	okH := math.Atan2(okB, okA) * 180.0 / math.Pi

	// Normalize hue to 0-360
	if okH < 0 {
		okH += 360
	}

	return []float64{okL, okC, okH}
}

func FromHex(hex string) LABColor {
	// Remove # if present
	hex = strings.TrimPrefix(hex, "#")

	// Parse hex values
	r, _ := strconv.ParseInt(hex[0:2], 16, 0)
	g, _ := strconv.ParseInt(hex[2:4], 16, 0)
	b, _ := strconv.ParseInt(hex[4:6], 16, 0)

	// Convert to 0-1 range and call RGB conversion
	return FromRGB(float64(r)/255.0, float64(g)/255.0, float64(b)/255.0)
}

func FromHSL(h float64, s float64, l float64) LABColor {
	// Convert HSL to RGB first
	s = s / 100.0
	l = l / 100.0

	c := (1 - math.Abs(2*l-1)) * s
	x := c * (1 - math.Abs(math.Mod(h/60.0, 2)-1))
	m := l - c/2

	var r, g, b float64

	switch {
	case h < 60:
		r, g, b = c, x, 0
	case h < 120:
		r, g, b = x, c, 0
	case h < 180:
		r, g, b = 0, c, x
	case h < 240:
		r, g, b = 0, x, c
	case h < 300:
		r, g, b = x, 0, c
	default:
		r, g, b = c, 0, x
	}

	return FromRGB(r+m, g+m, b+m)
}

func FromLCH(l float64, c float64, h float64) LABColor {
	// Convert LCH to LAB
	// LCH uses cylindrical coordinates, LAB uses rectangular
	hRad := h * math.Pi / 180.0

	color := LABColor{
		L: l,
		A: c * math.Cos(hRad),
		B: c * math.Sin(hRad),
	}

	return color
}

func FromRGB(r float64, g float64, b float64) LABColor {
	// Convert RGB to XYZ
	r = gammaCorrect(r)
	g = gammaCorrect(g)
	b = gammaCorrect(b)

	// RGB to XYZ conversion matrix (sRGB D65)
	x := r*0.4124564 + g*0.3575761 + b*0.1804375
	y := r*0.2126729 + g*0.7151522 + b*0.0721750
	z := r*0.0193339 + g*0.1191920 + b*0.9503041

	// Convert XYZ to LAB
	return xyzToLab(x, y, z)
}

func gammaCorrect(c float64) float64 {
	if c <= 0.04045 {
		return c / 12.92
	}
	return math.Pow((c+0.055)/1.055, 2.4)
}

func xyzToLab(x, y, z float64) LABColor {
	// Reference white D65
	const xn float64 = 0.95047
	const yn float64 = 1.00000
	const zn float64 = 1.08883

	x = x / xn
	y = y / yn
	z = z / zn

	x = labF(x)
	y = labF(y)
	z = labF(z)

	l := 116*y - 16
	a := 500 * (x - y)
	b := 200 * (y - z)

	return LABColor{L: l, A: a, B: b}
}

// Helper function for LAB conversion
func labF(t float64) float64 {
	const delta float64 = 6.0 / 29.0
	if t > delta*delta*delta {
		return math.Pow(t, 1.0/3.0)
	}
	return t/(3*delta*delta) + 4.0/29.0
}

func labFInverse(t float64) float64 {
	const delta float64 = 6.0 / 29.0
	if t > delta {
		return t * t * t
	}
	return 3 * delta * delta * (t - 4.0/29.0)
}

func labToXyz(l, a, b float64) (float64, float64, float64) {
	// Reference white D65
	const xn float64 = 0.95047
	const yn float64 = 1.00000
	const zn float64 = 1.08883

	fy := (l + 16) / 116
	fx := a/500 + fy
	fz := fy - b/200

	x := xn * labFInverse(fx)
	y := yn * labFInverse(fy)
	z := zn * labFInverse(fz)

	return x, y, z
}

// Helper function for inverse gamma correction (linear to sRGB)
func inverseGammaCorrect(c float64) float64 {
	if c <= 0.0031308 {
		return c * 12.92
	}
	return 1.055*math.Pow(c, 1.0/2.4) - 0.055
}

// Helper function to clamp values
func clamp(value, min, max float64) float64 {
	if value < min {
		return min
	}
	if value > max {
		return max
	}
	return value
}

// Helper function to pad hex values with leading zeros
func padHex(hex string) string {
	if len(hex) == 1 {
		return "0" + hex
	}
	return hex
}
