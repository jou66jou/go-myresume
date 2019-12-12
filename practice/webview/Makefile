build.webview:
	CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o webview.exe
build.startWebView:
	GOOS=windows GOARCH=amd64 go build -o startWebView.exe