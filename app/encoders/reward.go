package encoders

import (
	"log"
	"encoding/json"
	"io/ioutil"
	"github.com/marwaan/testproject/app/models"
	"io"
)

func EncodeReward(body io.ReadCloser) (reward models.Reward) {
	var data,_ = ioutil.ReadAll(body)
	if err := json.Unmarshal(data, &reward); err != nil {
		log.Println(err)
		return
	}
	return

}