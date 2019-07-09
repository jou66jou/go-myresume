package debug

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

// PrintRequestBody(*http.Request) 將request的body內容print出來
func PrintRequestBody(r *http.Request) {

	requestDump, e := httputil.DumpRequest(r, true)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(requestDump))
}

func PrintRespBody(r *http.Response) {

	requestDump, e := httputil.DumpResponse(r, true)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(string(requestDump))
}
