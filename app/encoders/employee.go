package encoders

import (
"log"
"encoding/json"
"io/ioutil"
"github.com/marwaan/testproject/app/models"
"io"
)

func EncodeEmployee(body io.ReadCloser) (employee models.Employee) {
	var data,_ = ioutil.ReadAll(body)
	if err := json.Unmarshal(data, &employee); err != nil {
		log.Println(err)
		return
	}
	return

}