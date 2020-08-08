package screen

import (
	"encoding/json"
	"fmt"
	"net/http"

	"fyne.io/fyne"
	"fyne.io/fyne/widget"
	"github.com/gangjun06/gChat/server/ws"
)

func HomeScreen(a fyne.App) fyne.CanvasObject {

	inputIP := widget.NewEntry()
	inputIP.Disable()

	info := widget.NewLabel("")
	serverOn := false

	// Get my IP
	go func(entry *widget.Entry) {
		var ipInfo map[string]interface{}
		resp, _ := http.Get("https://api.myip.com")
		json.NewDecoder(resp.Body).Decode(&ipInfo)
		ip := fmt.Sprintf("%v", ipInfo["ip"])
		resp.Body.Close()
		entry.SetText(ip)
	}(inputIP)

	inputPort := widget.NewEntry()
	inputPort.SetText("8080")
	inputPort.SetPlaceHolder("Enter Chatapp Port")

	formServe := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "YourIP", Widget: inputIP},
			{Text: "Port", Widget: inputPort},
		},
		OnSubmit: func() {
			if serverOn {
				return
			}
			serverOn = true
			info.SetText("Server is started. Check the log.")
			go ws.Serve(inputPort.Text)
		},
		CancelText: "",
		SubmitText: "Start Server",
	}

	return widget.NewVBox(
		info,
		formServe,
	)
}
