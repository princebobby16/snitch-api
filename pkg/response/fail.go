package response

type FailResponse struct {
	Status string `json:"status"`
	Data struct{} `json:"data"`
}
