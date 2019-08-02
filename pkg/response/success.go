package response

type SuccessResponse struct {
	Status string `json:"status"`
	Data struct{} `json:"data"`
	Links []link `json:"links"`
}

type link struct {
	Rel string `json:"rel"`
	Href string `json:"href"`
	Action string `json:"action"`
}
