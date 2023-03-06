package globalhandler

type ResultRouteError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type BasicOutputResponse struct {
	Response BasicResponse     `json:"response"`
	Param    map[string]string `json:"param"`
	Results  interface{}       `json:"results"`
}
type BasicResponse struct {
	Status    bool   `json:"status"`
	Message   string `json:"message"`
	StatusMSG string `json:"status_msg"`
}
