run:
	go run .

build: desktop open

movile:
	fyne-cross android -icon="logo-unamad.png" -name="OTI" -app-id="version.xyz"

desktop:
	fyne-cross windows -arch=amd64 -icon="logo-unamad.png" -app-version="1.5.0" -name="OTI" -app-id="v.1.5.0"

open:
	explorer.exe .

test: 
	go test -v ./entity/...
