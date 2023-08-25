package gui

import (
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	"github.com/luispfcanales/daemon-service-oti/entity"
	"github.com/luispfcanales/daemon-service-oti/gui/textbox"
	"github.com/luispfcanales/daemon-service-oti/model"
	"github.com/luispfcanales/daemon-service-oti/services/post"
	"github.com/luispfcanales/daemon-service-oti/services/stream"
)

const (
	WIDTH_WIN         float32 = 515
	HEIGHT_WIN        float32 = 500
	HEIGHT_ENTRY      float32 = 35
	WIDTH_ENTRY       float32 = 360
	LEFT_WIDTH_ENTRY  float32 = 115
	RIGHT_WIDTH_ENTRY float32 = 107
)

func Run(STREAM_SRV *stream.StreamSrv) {
	a := app.New()
	w := a.NewWindow("OTI soporte")
	w.Resize(fyne.NewSize(WIDTH_WIN, HEIGHT_WIN))
	w.CenterOnScreen()

	//txt bindings
	strPatrimonialCode, entryPatrimonialCode := textbox.NewStrPatrimonialCode()
	strBIOS, entryBIOS := textbox.NewStrBios()
	strHOSTNAME, entryHOSTNAME := textbox.NewStrHostname()
	strFACTURER, entryFACTURER := textbox.NewStrFacturer()
	strMODEL, entryMODEL := textbox.NewStrModel()
	strARCHITECTURE, entryARCHITECTURE := textbox.NewStrArchitecture()
	strRAM, entryRAM := textbox.NewStrRam()
	strPROCESSOR, entryPROCESSOR := textbox.NewStrProcessor()
	strCORE, entryCORE := textbox.NewStrCore()
	strLOGIC_CORE, entryLOGIC_CORE := textbox.NewStrLogic()
	strDISK_TYPE, entryDISK_TYPE := textbox.NewStrDisktype()
	strDISK_SIZE, entryDISK_SIZE := textbox.NewStrDisksize()

	//buttons to load and send information
	btnSendInfoApi := widget.NewButtonWithIcon("Enviar Info.", theme.ConfirmIcon(), nil)
	btnSendInfoApi.Importance = 1
	btnSendInfoApi.Disable()
	btnSendInfoApi.OnTapped = func() {
		var requestPC model.RequestComputer

		requestPC.PatrimonialCode, _ = strPatrimonialCode.Get()
		requestPC.Serial, _ = strBIOS.Get()
		requestPC.Name, _ = strHOSTNAME.Get()
		requestPC.Maker, _ = strFACTURER.Get()
		requestPC.Model, _ = strMODEL.Get()
		requestPC.Architecture, _ = strARCHITECTURE.Get()
		requestPC.Ram, _ = strRAM.Get()
		requestPC.Processor, _ = strPROCESSOR.Get()
		requestPC.Core, _ = strCORE.Get()
		requestPC.LogicalCore, _ = strLOGIC_CORE.Get()
		requestPC.SizeDisk, _ = strDISK_SIZE.Get()
		requestPC.Disk, _ = strDISK_TYPE.Get()

		log.Println(requestPC)

		btnSendInfoApi.Disable()

		data, _ := strPatrimonialCode.Get()
		if data == "" {
			m := ModalGeneral("Error", "Ingrese Codigo de Patrimonio", w)
			m.Show()
			btnSendInfoApi.Enable()
			return
		}
		ok, value := post.SendDataAPI(&requestPC)
		if !ok {
			m := ModalGeneral("Error", "Internet not CONNECTED!", w)
			m.Show()
			return
		}

		qrModal := dialog.NewCustom("QR CODE", "Cerrar Ventana", GetCanvasQRCode(value), w)
		qrModal.Show()
	}

	btnLoad := widget.NewButtonWithIcon("Cargar Info.", theme.MediaReplayIcon(), nil)
	btnLoad.OnTapped = func() {
		btnLoad.Disable()
		c := entity.NewCommand()

		compSystem := entity.NewComputerSystem(c)
		cpuSys := entity.NewCPUSystem(c)
		disk := entity.NewPhysicalDisk(c)
		bios := entity.NewBios(c)

		entity.NewSystemDescriptor().Run(compSystem, cpuSys, disk, bios)

		strBIOS.Set(string(bios.SerialNumber))

		strHOSTNAME.Set(string(compSystem.Hostname))
		strFACTURER.Set(string(compSystem.Manufacturer))
		strMODEL.Set(string(compSystem.Model))
		strARCHITECTURE.Set(string(compSystem.System))
		strRAM.Set(fmt.Sprintf("%v", compSystem.TotalPhysicalMemory))

		strPROCESSOR.Set(string(cpuSys.Name))
		strCORE.Set(string(cpuSys.Cores))
		strLOGIC_CORE.Set(string(cpuSys.LogicalProcessors))

		strDISK_TYPE.Set(string(disk.MediaType))
		strDISK_SIZE.Set(string(disk.Size))

		btnSendInfoApi.Enable()

		//a.SendNotification()
	}
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
			rowField("Patrimonio", entryPatrimonialCode),
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
				btnSendInfoApi,
			),
		),
	)
	w.SetContent(mainContainer)
	w.SetFixedSize(true)

	icon, err := fyne.LoadResourceFromPath("logo-unamad.png")
	if err != nil {
		panic("no se cargo el icono")
	}
	w.SetIcon(icon)

	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu("Myapp",
			fyne.NewMenuItem("Show", func() {
				w.Show()
			}))
		desk.SetSystemTrayMenu(m)
		desk.SetSystemTrayIcon(icon)
	}

	w.SetCloseIntercept(func() {
		w.Hide()
	})
	//w.ShowAndRun()
	a.Run()
}

// ModalGeneral return modal information
func ModalGeneral(title, info string, parent fyne.Window) dialog.Dialog {
	render := dialog.NewInformation(title, info, parent)
	return render
}

func rowField(lbl string, entry *widget.Entry) *fyne.Container {
	return container.NewHBox(
		widget.NewLabel(lbl),
		container.NewWithoutLayout(
			entry,
		),
	)
}
