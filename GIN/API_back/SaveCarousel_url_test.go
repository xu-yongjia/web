package api2

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestSaveCarouselURL(t *testing.T) {
	go server()
	time.Sleep(time.Second * 2)

	url := "http://127.0.0.1:3000/api1/v2/admin/saveCarouselUrl"
	data := make(map[string]interface{})
	data["canteen_id"] = 1
	data["carousel_url"] = "http://example.com/carousel_image.jpg"
	bytesData, err := json.Marshal(data)
	handleError(t, err)

	reqest, err := http.NewRequest("POST", url, bytes.NewReader(bytesData))
	handleError(t, err)
	reqest.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(reqest)
	handleError(t, err)
	defer resp.Body.Close()
	
	if resp == nil {
		t.Fatal("Received nil response")
	}

	body, err := ioutil.ReadAll(resp.Body)
	handleError(t, err)

	var str bytes.Buffer
	err = json.Indent(&str, []byte(body), "", "    ")
	handleError(t, err)

	t.Log(str.String())
}
