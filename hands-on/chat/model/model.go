package model

type Chat struct {
	ID        string `json:"id"`
	From      string `json:"from"`
	To        string `json:"to"`
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

type Contacts struct {
	Username string `json:"username"`
	Activity int64  `json:"activity"`
}
