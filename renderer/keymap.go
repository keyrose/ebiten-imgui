package renderer

import (
	"github.com/AllenDang/cimgui-go"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var keys = map[int]int{

	cimgui.ImGuiKey_Tab:        int(ebiten.KeyTab),
	cimgui.ImGuiKey_LeftArrow:  int(ebiten.KeyLeft),
	cimgui.ImGuiKey_RightArrow: int(ebiten.KeyRight),
	cimgui.ImGuiKey_UpArrow:    int(ebiten.KeyUp),
	cimgui.ImGuiKey_DownArrow:  int(ebiten.KeyDown),
	cimgui.ImGuiKey_PageUp:     int(ebiten.KeyPageUp),
	cimgui.ImGuiKey_PageDown:   int(ebiten.KeyPageDown),
	cimgui.ImGuiKey_Home:       int(ebiten.KeyHome),
	cimgui.ImGuiKey_End:        int(ebiten.KeyEnd),
	cimgui.ImGuiKey_Insert:     int(ebiten.KeyInsert),
	cimgui.ImGuiKey_Delete:     int(ebiten.KeyDelete),
	cimgui.ImGuiKey_Backspace:  int(ebiten.KeyBackspace),
	cimgui.ImGuiKey_Space:      int(ebiten.KeySpace),
	cimgui.ImGuiKey_Enter:      int(ebiten.KeyEnter),
	cimgui.ImGuiKey_Escape:     int(ebiten.KeyEscape),
	cimgui.ImGuiKey_A:          int(ebiten.KeyA),
	cimgui.ImGuiKey_C:          int(ebiten.KeyC),
	cimgui.ImGuiKey_V:          int(ebiten.KeyV),
	cimgui.ImGuiKey_X:          int(ebiten.KeyX),
	cimgui.ImGuiKey_Y:          int(ebiten.KeyY),
	cimgui.ImGuiKey_Z:          int(ebiten.KeyZ),
}

func sendInput(io *cimgui.ImGuiIO, inputChars []rune) []rune {

	// Ebiten hides the LeftAlt RightAlt implementation (inside the uiDriver()), so
	// here only the left alt is sent
	if ebiten.IsKeyPressed(ebiten.KeyAlt) {
		io.SetKeyAlt(true)

	} else {
		io.SetKeyAlt(false)
	}
	if ebiten.IsKeyPressed(ebiten.KeyShift) {
		io.SetKeyShift(true)
	} else {
		io.SetKeyShift(false)
	}
	if ebiten.IsKeyPressed(ebiten.KeyControl) {
		io.SetKeyCtrl(true)
	} else {
		io.SetKeyCtrl(false)
	}

	// if ebiten.IsKeyPressed(ebiten.KeyLeftSuper)

	// TODO: get KeySuper somehow (GLFW: KeyLeftSuper    = Key(343), R: 347)
	inputChars = ebiten.AppendInputChars(inputChars)
	if len(inputChars) > 0 {
		io.AddInputCharactersUTF8(string(inputChars))
		inputChars = inputChars[:0]
	}
	for _, iv := range keys {
		if inpututil.IsKeyJustPressed(ebiten.Key(iv)) {

			io.AddKeyEvent(cimgui.ImGuiKey(iv), true)
			// io.KeyPress(iv)
		}
		if inpututil.IsKeyJustReleased(ebiten.Key(iv)) {
			io.AddKeyEvent(cimgui.ImGuiKey(iv), false)
			//io.KeyRelease(iv)
		}
	}
	return inputChars
}
