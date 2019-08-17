package index

type Alive struct {
	Alive       bool   `json:"alive"`
	Author      string `json:"author"`
	Version     string `json:"version"`
	Environment string `json:"environment"`
}
