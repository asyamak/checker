package entity

type StringRequest struct {
	Text string `json:"text"`
}

type EmailRequest struct {
	Emails []string `json:"emails"`
}

type IinRequest struct {
	IIN uint64 `json:"iin"`
}
