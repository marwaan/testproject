package encoders

import (
	"log"
	"encoding/json"
	"io/ioutil"
	"github.com/marwaan/testproject/app/models"
	"io"
)

func EncodePayoff(body io.ReadCloser) (payoff models.Payoff) {
	var data,_ = ioutil.ReadAll(body)
	if err := json.Unmarshal(data, &payoff); err != nil {
		log.Println(err)
		return
	}
	return

}