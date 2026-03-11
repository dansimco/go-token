package typography

import (
	"go-token/css_util"
	"testing"
)

type typographyTestInputData struct {
	Family Family
	Styles []Style
}

func testSetup() typographyTestInputData {

	helvetica_now := NewFontFamily("Helvetica Now")

	helvetica_now_regular := (&helvetica_now).AddFont()
	helvetica_now_regular.AddLocalSrc("Helvetica Now")
	helvetica_now_regular.AddSrc("/assets/fonts/helvetica_now_regular.woff2")
	helvetica_now_regular.SetWeight("regular")

	helvetica_now_bold := (&helvetica_now).AddFont()
	helvetica_now_bold.AddLocalSrc("Helvetica Now")
	helvetica_now_bold.AddSrc("/assets/fonts/helvetica_now_bold")
	helvetica_now_bold.SetWeightNumber(600)

	(&helvetica_now).AddFallbackFont("arial")

	// styles
	//
	t_body := NewTypeStyle("body")
	t_body.SetFamily(&helvetica_now)
	t_body.SetSize(0.9)
	t_body.SetWeight("regular")
	t_body.SetLineHeight(1.5)

	t_caption := NewTypeStyle("caption")
	t_caption.SetFamily(&helvetica_now)
	t_caption.SetSize(0.875)
	t_caption.SetWeight("regular")
	t_caption.SetLineHeight(1.5)
	t_caption.SetTracking(0.03)

	return typographyTestInputData{
		Family: helvetica_now,
		Styles: []Style{t_body, t_caption},
	}
}

func TestCSSFontFaceGeneration(t *testing.T) {

	family := testSetup().Family

	generated_css := css_util.Format(family.ToCSS())

	expected_css := `

	@font-face {
	  font-family: "Helvetica Now";
	  src:
        local("Helvetica Now"),
        url("/assets/fonts/helvetica_now_regular.woff2");
      font-weight: "regular";
    }

   	@font-face {
	  font-family: "Helvetica Now";
	  src:
        local("Helvetica Now"),
        url("/assets/fonts/helvetica_now_bold.woff2");
      font-weight: 600;
    }

	`

	expected_css = css_util.Format(expected_css)

	if generated_css != expected_css {
		t.Errorf("Expected generated_css to equal \n %s \n got \n %s", expected_css, generated_css)
	}

}

func TestCSSStyleGeneration(t *testing.T) {

	expected_css := `

		.body {
		    font-family: "Helvetica Now", "arial";
		    font-size: 0.9rem;
		    line-height: 1.5rem;
		    font-weight: regular;
		}

		.caption {
		    font-family: "Helvetica Now", "arial";
		    font-size: 0.875rem;
		    line-height: 1.5rem;
			letter-spacing: 0.03rem;
		    font-weight: regular;
		}

	`
	expected_css = css_util.Format(expected_css)

	generated_css := ""
	styles := testSetup().Styles

	for _, style := range styles {
		generated_css += style.ToCSS()
	}

	generated_css = css_util.Format(generated_css)

	println(generated_css)

	if generated_css != expected_css {
		t.Errorf("Expected generated_css to equal \n %s \n got \n %s", expected_css, generated_css)
	}

}
