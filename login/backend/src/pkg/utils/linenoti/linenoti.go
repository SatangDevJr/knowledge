package linenoti

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func LineNoti(message string) int {
	client := resty.New()
	var result map[string]string
	LineMessage := fmt.Sprintf(`{
		"message": "%[1]v"
	 }`, message)
	json.Unmarshal([]byte(LineMessage), &result)

	resp, err := client.R().
		SetHeader("Authorization", "Bearer rhYu1UQcLqlC2Yew6pn5nS05efO1kr2rfL2kXnGvudE").
		SetFormData(result).Post("https://notify-api.line.me/api/notify")
	if err != nil {
		fmt.Println("ERROR LINE Notify API:", err)
	}

	return resp.StatusCode()
}
