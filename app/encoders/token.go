package encoders

import (
	"io"
	"github.com/marwaan/testproject/app/models"
	"io/ioutil"
	"encoding/json"
	"log"
	//"github.com/marwaan/testproject/app/util"
)

func Encodetoken(body io.ReadCloser) (token models.Token) {
	var data,_ = ioutil.ReadAll(body)

	if err := json.Unmarshal(data, &token); err != nil {
		log.Println(err)
		return
	}
	return

}
