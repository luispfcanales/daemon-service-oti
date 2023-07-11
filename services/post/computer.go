package post

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/luispfcanales/daemon-service-oti/model"
)

var urlPOST string = "https://oti.vercel.app/api/upload"

func SendDataAPI(pc *model.RequestComputer) (bool, string) {
	jsonByte, err := json.Marshal(pc)
	if err != nil {
		log.Println("marshal :", err)
		return false, ""
	}

	req, err := http.NewRequest(http.MethodPost, urlPOST, bytes.NewBuffer(jsonByte))
	if err != nil {
		log.Println("error creating request: ", err)
		return false, ""
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("error sending request: ", err)
		return false, ""
	}
	defer resp.Body.Close()

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error reading response: ", err)
		return false, ""
	}

	rApi := model.ResponseApi{}
	json.Unmarshal(responseBody, &rApi)
	log.Println("success: ", rApi)
	return true, rApi.Code
}
