# Go Design Tokens (WIP!)

This module is designed to streamline the creation of design systems. Docs to come but if you're here early an example implementation can be seen at [https://github.com/dansimco/go-design-tokens/blob/main/theme/theme_test.go](https://github.com/dansimco/go-design-tokens/blob/main/theme/theme_test.go)

## Features

### Color Utilities

#### Color Ramps

![](doc/img/ramp_extended.png)

Color Ramps are the original motivation for this package. Rather than creating a rigid list of swatches (`orange-100`, `orange-200`) etc. One can simply create a ramp which blends from one color to another and pick anywhere on that gradient. 

```
intlOrange := color.NewRamp()
intlOrange.AddKey("#FAF9FF", 0)
intlOrange.AddKey("#FF4F01", 0.406)
intlOrange.AddKey("#010000", 1)

println(intlOrange.At(0.4).ToHex()) // returns #D34409

```
