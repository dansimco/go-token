# Go Design Tokens (WIP!)

This module is designed to streamline the creation of design systems. Docs to come but if you're here early an example implementation can be seen at [https://github.com/dansimco/go-design-tokens/blob/main/theme/theme_test.go](https://github.com/dansimco/go-design-tokens/blob/main/theme/theme_test.go)

## Features

### Color Ramps

![](doc/img/ramp_extended.png)

Color Ramps are the original motivation for this package. Rather than creating a rigid list of swatches (`orange-100`, `orange-200`) etc. One can simply create a ramp which blends from one color to another and pick anywhere on that gradient. 

```
intlOrange := color.NewRamp()
intlOrange.AddKey("#FAF9FF", 0)
intlOrange.AddKey("#FF4F01", 0.406)
intlOrange.AddKey("#010000", 1)

println(intlOrange.At(0.4).ToHex()) // returns #D34409
```

### Color Tokens

Color tokens are structured in a three-tier system:

**mode -> role -> state**

As personal systems and naming preferences differ, no mode/role/state names are prescribed. 

#### Mode

Modes are akin to mini-themes, and can be presentational, semantic, whatever you like.

The term "mode" is in line with the [mise en mode](https://mode.place) design system practice. 

```
action := theme.AddColorMode("action")
```

#### Role

A role is what the color is used for. Is it text content? background fill? border? etc.

**Context** sets whether a color role is foreground or background. This allows automatic code generation (color vs background-color), and color-contrast accessibility checks in future. 

```
action_content := action.AddRole("content")
action_content.SetContext("foreground")
```

#### State

State is open ended, but generally intented for ui interaction states. 

```
action_content.AddState("default", color.UIColor{Light: azimuth.At(0.4), Dark: azimuth.At(0.4375)})
action_content.AddState("hover", color.UIColor{Light: azimuth.At(0.37), Dark: azimuth.At(0.35)})
action_content.AddState("pressed", color.UIColor{Light: azimuth.At(0.7), Dark: azimuth.At(0.3)})
action_content.AddState("focus", color.UIColor{Light: azimuth.At(0.6), Dark: azimuth.At(0.35)})
```


**default** is a special state as it will not add "default" to the end of the token name, for code brevity. Eg. `action -> content -> default` would resolve to `var(--action-content)` in css where as `action -> content -> hover` would resolve to `var(--action-content-hover)`

Light and dark modes are built in to UIColor
