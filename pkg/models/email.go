package models

type Mail struct {
	Domain      string
	Templates   string
	Host        string
	Port        string
	Username    string
	Password    string
	Encryption  string
	FromAddress string
	FromName    string
	ApiUrl      string
	ApiKey      string

	Jobs    chan Message
	Results chan Message
}

type Message struct {
	From        string
	FromName    string
	Subject     string
	Template    string
	Attachments []string
	Data        interface{}
}

type Results struct {
	Success bool
	Error   error
}
