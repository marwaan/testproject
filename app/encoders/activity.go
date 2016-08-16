package encoders

import (
	"log"
	"encoding/json"
	"io/ioutil"
	"github.com/marwaan/testproject/app/models"
	"io"
)

func EncodeActivity(body io.ReadCloser) (activity models.Activity) {
	var data,_ = ioutil.ReadAll(body)
	if err := json.Unmarshal(data, &activity); err != nil {
		log.Println(err)
		return
	}
	return

}