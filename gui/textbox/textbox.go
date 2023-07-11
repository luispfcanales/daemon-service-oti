package textbox

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

const (
	WIDTH_WIN         float32 = 515
	HEIGHT_WIN        float32 = 500
	HEIGHT_ENTRY      float32 = 35
	WIDTH_ENTRY       float32 = 360
	LEFT_WIDTH_ENTRY  float32 = 115
	RIGHT_WIDTH_ENTRY float32 = 107
)

func textEntry(sizeWidth float32, str binding.String) *widget.Entry {
	e := widget.NewEntry()
	e.Resize(fyne.NewSize(sizeWidth, HEIGHT_ENTRY))
	e.Bind(str)

	return e
}

func NewStrPatrimonialCode() (binding.String, *widget.Entry) {
	str := binding.NewString()
	entry := textEntry(WIDTH_ENTRY, str)
	entry.Move(fyne.NewPos(30, 0))
	return str, entry
}

func NewStrBios() (binding.String, *widget.Entry) {
	str := binding.NewString()
	entry := textEntry(WIDTH_ENTRY, str)
	entry.Move(fyne.NewPos(62, 0))
	return str, entry
}
func NewStrHostname() (binding.String, *widget.Entry) {
	str := binding.NewString()
	entry := textEntry(WIDTH_ENTRY, str)
	return str, entry
}
func NewStrFacturer() (binding.String, *widget.Entry) {
	str := binding.NewString()
	entry := textEntry(WIDTH_ENTRY, str)
	entry.Move(fyne.NewPos(32, 0))
	return str, entry
}
func NewStrModel() (binding.String, *widget.Entry) {
	str := binding.NewString()
	entry := textEntry(WIDTH_ENTRY, str)
	entry.Move(fyne.NewPos(50, 0))
	return str, entry
}
func NewStrArchitecture() (binding.String, *widget.Entry) {
	str := binding.NewString()
	entry := textEntry(WIDTH_ENTRY, str)
	entry.Move(fyne.NewPos(20, 0))
	return str, entry
}
func NewStrRam() (binding.String, *widget.Entry) {
	str := binding.NewString()
	entry := textEntry(LEFT_WIDTH_ENTRY, str)
	entry.Move(fyne.NewPos(69, 0))
	return str, entry
}
func NewStrProcessor() (binding.String, *widget.Entry) {
	str := binding.NewString()
	entry := textEntry(WIDTH_ENTRY, str)
	entry.Move(fyne.NewPos(27, 0))
	return str, entry
}
func NewStrCore() (binding.String, *widget.Entry) {
	str := binding.NewString()
	entry := textEntry(LEFT_WIDTH_ENTRY, str)
	entry.Move(fyne.NewPos(47, 0))
	return str, entry
}
func NewStrLogic() (binding.String, *widget.Entry) {
	str := binding.NewString()
	entry := textEntry(RIGHT_WIDTH_ENTRY, str)
	return str, entry
}
func NewStrDisktype() (binding.String, *widget.Entry) {
	str := binding.NewString()
	entry := textEntry(LEFT_WIDTH_ENTRY, str)
	entry.Move(fyne.NewPos(63, 0))
	return str, entry
}
func NewStrDisksize() (binding.String, *widget.Entry) {
	str := binding.NewString()
	entry := textEntry(RIGHT_WIDTH_ENTRY, str)
	entry.Move(fyne.NewPos(10, 0))
	return str, entry
}
