package gui

import (
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/luispfcanales/daemon-service-oti/entity"
)

const (
	WIDTH_WIN         float32 = 515
	HEIGHT_WIN        float32 = 500
	HEIGHT_ENTRY      float32 = 35
	WIDTH_ENTRY       float32 = 360
	LEFT_WIDTH_ENTRY  float32 = 115
	RIGHT_WIDTH_ENTRY float32 = 107
)

func Run() {
	a := app.New()
	w := a.NewWindow("OTI soporte")
	w.Resize(fyne.NewSize(WIDTH_WIN, HEIGHT_WIN))
	w.CenterOnScreen()

	//txt bindings
	strBIOS := binding.NewString()
	entryBIOS := textEntry(WIDTH_ENTRY, strBIOS)
	entryBIOS.Move(fyne.NewPos(62, 0))

	strHOSTNAME := binding.NewString()
	entryHOSTNAME := textEntry(WIDTH_ENTRY, strHOSTNAME)
	strFACTURER := binding.NewString()
	entryFACTURER := textEntry(WIDTH_ENTRY, strFACTURER)
	entryFACTURER.Move(fyne.NewPos(32, 0))
	strMODEL := binding.NewString()
	entryMODEL := textEntry(WIDTH_ENTRY, strMODEL)
	entryMODEL.Move(fyne.NewPos(50, 0))
	strARCHITECTURE := binding.NewString()
	entryARCHITECTURE := textEntry(WIDTH_ENTRY, strARCHITECTURE)
	entryARCHITECTURE.Move(fyne.NewPos(20, 0))
	strRAM := binding.NewString()
	entryRAM := textEntry(LEFT_WIDTH_ENTRY, strRAM)
	entryRAM.Move(fyne.NewPos(69, 0))

	strPROCESSOR := binding.NewString()
	entryPROCESSOR := textEntry(WIDTH_ENTRY, strPROCESSOR)
	entryPROCESSOR.Move(fyne.NewPos(27, 0))
	strCORE := binding.NewString()
	entryCORE := textEntry(LEFT_WIDTH_ENTRY, strCORE)
	entryCORE.Move(fyne.NewPos(47, 0))
	strLOGIC_CORE := binding.NewString()
	entryLOGIC_CORE := textEntry(RIGHT_WIDTH_ENTRY, strLOGIC_CORE)

	strDISK_TYPE := binding.NewString()
	entryDISK_TYPE := textEntry(LEFT_WIDTH_ENTRY, strDISK_TYPE)
	entryDISK_TYPE.Move(fyne.NewPos(63, 0))
	strDISK_SIZE := binding.NewString()
	entryDISK_SIZE := textEntry(RIGHT_WIDTH_ENTRY, strDISK_SIZE)
	entryDISK_SIZE.Move(fyne.NewPos(10, 0))

	//buttons to load and send information
	btnSuccess := widget.NewButtonWithIcon("Enviar Info.", theme.ConfirmIcon(), nil)
	btnSuccess.Importance = 1
	btnSuccess.Disable()

	btnLoad := widget.NewButtonWithIcon("Cargar Info.", theme.MediaReplayIcon(), func() {
		c := entity.NewCommand()
		compSystem := entity.NewComputerSystem(c)
		cpuSys := entity.NewCPUSystem(c)
		disk := entity.NewPhysicalDisk(c)
		entity.NewSystemDescriptor().Run(compSystem, cpuSys, disk)

		strFACTURER.Set(string(compSystem.Manufacturer))
		strMODEL.Set(string(compSystem.Model))
		strARCHITECTURE.Set(string(compSystem.System))
		strRAM.Set(fmt.Sprintf("%v", compSystem.TotalPhysicalMemory))

		strPROCESSOR.Set(string(cpuSys.Name))
		strCORE.Set(string(cpuSys.Cores))
		strLOGIC_CORE.Set(string(cpuSys.LogicalProcessors))

		strDISK_TYPE.Set(string(disk.MediaType))
		strDISK_SIZE.Set(string(disk.Size))

		//data, _ := strBios.Get()
		btnSuccess.Enable()
	})
	btnLoad.Importance = 2

	//header GUI
	lblHeader := canvas.NewText("Mi Equipo", color.RGBA{R: 255, G: 23, B: 122, A: 255})
	lblHeader.TextSize = 20
	lblHeader.TextStyle.Bold = true

	//content to render GUI
	mainContainer := container.NewVBox(
		container.NewCenter(
			lblHeader,
		),
		container.NewVBox(
			rowField("Serial", entryBIOS),
			rowField("Nombre equipo", entryHOSTNAME),
			rowField("Fabricante", entryFACTURER),
			rowField("Modelo", entryMODEL),
			rowField("Arquitectura", entryARCHITECTURE),
			rowField("Procesador", entryPROCESSOR),
			container.NewGridWithColumns(
				2,
				rowField("Nucleos", entryCORE),
				rowField("Nucleos logicos", entryLOGIC_CORE),
			),
			container.NewGridWithColumns(
				2,
				rowField("Disco", entryDISK_TYPE),
				rowField("Tamano Disco", entryDISK_SIZE),
			),
			rowField("Ram", entryRAM),
		),
		container.NewCenter(
			container.NewHBox(
				btnLoad,
				btnSuccess,
			),
		),
	)
	w.SetContent(mainContainer)
	w.SetFixedSize(true)
	w.ShowAndRun()
}

func rowField(lbl string, entry *widget.Entry) *fyne.Container {
	return container.NewHBox(
		widget.NewLabel(lbl),
		container.NewWithoutLayout(
			entry,
		),
	)
}
func textEntry(sizeWidth float32, str binding.String) *widget.Entry {
	e := widget.NewEntry()
	e.Resize(fyne.NewSize(sizeWidth, HEIGHT_ENTRY))
	e.Bind(str)

	return e
}
