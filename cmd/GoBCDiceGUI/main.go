package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	// "github.com/raa0121/GoBCDiceGUI/diceBot"
)

var (
	setting      *ui.Combobox
	settingGroup *ui.Group
)

type areaHandler struct{}

func (areaHandler) Draw(a *ui.Area, p *ui.AreaDrawParams) {
	font := ui.FontDescriptor{
		Family:  "Cica",
		Size:    14,
		Weight:  ui.TextWeightNormal,
		Italic:  ui.TextItalicNormal,
		Stretch: ui.TextStretchNormal,
	}
	tl := ui.DrawNewTextLayout(&ui.DrawTextLayoutParams{
		String:      ui.NewAttributedString(""),
		DefaultFont: &font,
		Width:       p.AreaWidth,
		Align:       ui.DrawTextAlign(setting.Selected()),
	})
	defer tl.Free()
	p.Context.Text(tl, 0, 0)
}

func (areaHandler) MouseEvent(a *ui.Area, me *ui.AreaMouseEvent) {
	// do nothing
}

func (areaHandler) MouseCrossed(a *ui.Area, left bool) {
	// do nothing
}

func (areaHandler) DragBroken(a *ui.Area) {
	// do nothing
}

func (areaHandler) KeyEvent(a *ui.Area, ke *ui.AreaKeyEvent) (handled bool) {
	// reject all keys
	return false
}

func setupUI() {
	mainwin := ui.NewWindow("GoBCDice", 640, 480, true)
	mainwin.SetMargined(true)
	mainwin.OnClosing(func(*ui.Window) bool {
		mainwin.Destroy()
		ui.Quit()
		return false
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	mainwin.SetChild(hbox)

	area := ui.NewArea(areaHandler{})

	form := ui.NewForm()
	form.SetPadded(true)
	// TODO on OS X if this is set to 1 then the window can't resize; does the form not have the concept of stretchy trailing space?
	hbox.Append(form, true)

	setting = ui.NewCombobox()
	// note that the items match with the values of the uiDrawTextAlign values
	setting.Append("Left")
	setting.Append("Center")
	setting.Append("Right")
	setting.SetSelected(0) // start with left alignment
	setting.OnSelected(func(*ui.Combobox) {
		area.QueueRedrawAll()
	})
	form.Append("設定", setting, false)
	saveButton := ui.NewButton("この設定を保存")
	deleteButton := ui.NewButton("この設定を削除")
	form.Append("", saveButton, false)
	form.Append("", deleteButton, false)

	hbox.Append(area, true)

	mainwin.Show()
}

func main() {
	ui.Main(setupUI)
}
