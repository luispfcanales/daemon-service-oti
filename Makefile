build: fyne open

fyne:
	fyne-cross windows -arch=amd64 -icon="logo-unamad.png" -app-version="1.5.0" -name="OTI" -app-id="v.1.5.0"

open:
	explorer.exe .
