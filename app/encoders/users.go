package encoders

import (
	"io"
	"github.com/marwaan/testproject/app/models"
	"io/ioutil"
	"encoding/json"
	"log"
	//"github.com/marwaan/testproject/app/util"
)

func EncodeSingleUsers(body io.ReadCloser) (user models.User) {
	var data,_ = ioutil.ReadAll(body)

	if err := json.Unmarshal(data, &user); err != nil {
		log.Println(err)
		return
	}
	return

}
