build.webview:
	CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o webview.exe
build.startWebView:
	CGO_ENABLED=1 CC="x86_64-w64-mingw32-gcc" GOOS=windows GOARCH=amd64 go build -o startWebView.exe
build.webview.32:
	CGO_ENABLED=1 CC="i686-w64-mingw32-gcc"   GOOS=windows GOARCH="386" go build -ldflags -H=windowsgui -o webview.exe