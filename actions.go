package actions

import (
	"fmt"
	"net/http"
)

func getAnswer(url string) {
	resp, err := http.Get(url + "GetUpdates")
	if err != nil {
		fmt.Print(resp)
	}
}
