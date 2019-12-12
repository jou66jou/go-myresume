package browser

import (
	"os/exec"
	"runtime"
	"strconv"

	log "github.com/sirupsen/logrus"
)

var handler *browserHandler

// type iBrowserHandler interface {
// 	Start(processListener chan bool) error
// 	Stop() error
// 	KillProcess() error
// }

type browserHandler struct {
	cmd            *exec.Cmd
	pathToChromium []string
	pdvEndpoint    string
}

func newBrowserHandler(pathToChromium []string, pdvEndpoint string) *browserHandler {
	b := &browserHandler{}
	b.pathToChromium = pathToChromium
	b.pdvEndpoint = pdvEndpoint

	return b
}

func Init(pdvURL string) {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"./webview"}
	case "windows":
		args = []string{"webview.exe"}
	default:
		args = []string{""}
	}
	handler = newBrowserHandler(args, pdvURL)
}

func StartBrowser(done chan bool) {
	browserProcessListener := make(chan bool)
	defer close(browserProcessListener)
	go handler.Start(browserProcessListener)

	var tryReopenBrowser = true
	for {
		select {
		case <-browserProcessListener:
			if tryReopenBrowser {
				log.Warn("Browser process is stopped. Attempting to restart")
				go handler.Start(browserProcessListener)
			} else {
				log.Warn("Browser process is stopped. Will not attempt to restart")
			}

		case <-done:
			log.Info("Shutting down browser")
			tryReopenBrowser = false
			handler.KillProcess()
			return
		}
	}
}

func (b *browserHandler) Start(processListener chan bool) error {
	// endpoint := fmt.Sprintf("--app=%s", b.pdvEndpoint)

	// b.cmd = exec.Command(b.pathToChromium[0], append(b.pathToChromium[1:], b.pdvEndpoint)...)
	b.cmd = exec.Command(b.pathToChromium[0])
	err := b.cmd.Run()
	log.Println("cmd complete !")
	if !IsClosed(processListener) {
		if err != nil {
			log.WithError(err).Error("Error with the browser process")
			processListener <- false
		} else {
			processListener <- true
		}
	}
	return err
}

func (b *browserHandler) Stop() error {
	err := b.cmd.Process.Release()
	if err != nil {
		log.WithError(err).Fatal("Error shutting down chromium")
	}

	return err
}

func (b *browserHandler) KillProcess() error {
	log.Info("Killing browser process")
	kill := exec.Command("taskkill", "/T", "/F", "/PID", strconv.Itoa(b.cmd.Process.Pid))
	err := kill.Run()
	if err != nil {
		log.WithError(err).Error("Error killing chromium process")
	}
	log.Info("Browser process was killed")

	return err
}

func IsClosed(ch <-chan bool) bool {
	select {
	case <-ch:
		return true
	default:
	}

	return false
}
