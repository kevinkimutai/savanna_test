package ports

type SMSPort interface{
	SendSMS(msg string, phoneNumbers []string) (string, error)
}
