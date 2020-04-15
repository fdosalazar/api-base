package apibasebundle

type Api struct {
	Status  int         `json:"status"`
	Message Message     `json:"message"`
	Time    float64     `json:"time"`
	Error   string      `json:"error"`
	Data    interface{} `json:"data"`
}

type Message struct {
	Description string  `json:"description"`
	Result      float64 `json:"result"`
}
