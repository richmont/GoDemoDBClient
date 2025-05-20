package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type MainWindow struct {
	app           fyne.App
	mainWindow    fyne.Window
	mainContainer *fyne.Container
	selectedSGBD  string
}

func NewMainWindow(app fyne.App) MainWindow {
	a := app
	mainWindow := a.NewWindow("GoDemoDBClient")
	cont := container.NewVBox(widget.NewLabel("Hello World!"))
	mainWindow.SetContent(cont)

	mw := MainWindow{app: a, mainWindow: mainWindow, mainContainer: cont}
	mw.renderBaseContent()
	return mw
	/*w2 := a.NewWindow("Larger")
	w2.SetContent(widget.NewLabel("More content"))
	w2.Resize(fyne.NewSize(100, 100))
	w2.Show()

	//a.Run()*/
}

func (mw *MainWindow) Run() {
	mw.app.Run()
}
func (mw *MainWindow) ShowMainWindow() {
	mw.mainWindow.Show()
}

func (mw *MainWindow) add(content fyne.CanvasObject) {
	//mw.mainWindow.SetContent(content)
	mw.mainContainer.Add(content)
	mw.mainContainer.Refresh()
}

func (mw *MainWindow) renderBaseContent() {
	mw.add(widget.NewLabel("Select the SGBD"))
	sgbdList := []string{"MariaDB"}

	sgbdSelect := widget.NewSelect(sgbdList, mw.onSelectedSGBD)
	mw.add(sgbdSelect)

}

func (mw *MainWindow) onSelectedSGBD(selected string) {
	mw.selectedSGBD = selected

}
