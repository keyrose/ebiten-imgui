// Package imcolor is a utility package to convert between color.Color and imgui.Vec4.
package imcolor

import (
	imgui "github.com/AllenDang/cimgui-go"
	"image/color"
)

func ToVec4(c color.Color) imgui.ImVec4 {
	r32, g32, b32, a32 := c.RGBA()
	return imgui.ImVec4{
		X: float32(r32) / 0xffff,
		Y: float32(g32) / 0xffff,
		Z: float32(b32) / 0xffff,
		W: float32(a32) / 0xffff,
	}
}

func ToColor(c imgui.ImVec4) color.RGBA64 {
	return color.RGBA64{
		R: uint16(c.X * 0xffff),
		G: uint16(c.Y * 0xffff),
		B: uint16(c.Z * 0xffff),
		A: uint16(c.W * 0xffff),
	}
}
