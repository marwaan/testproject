package util

type Response struct  {
	Status string        `json:"status"`
	Body  interface{}    `json:"body"`

}
func ResponseError(msg string) Response  {
         var err Response
	err.Status = "failed";
	err.Body = msg;
	return  err
}

func ResponseSuccess(body interface{}) Response  {
	var suc Response
	suc.Status = "success";
	suc.Body = body;
	return suc
}