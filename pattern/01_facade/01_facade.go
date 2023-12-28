package main

import "fmt"

type EmailSender struct{}

func (s *EmailSender) SendEmail() {
	fmt.Println("Email sent!")
}

type SMSSender struct{}

func (s *SMSSender) SendSMS() {
	fmt.Println("SMS sent!")
}

type PushSender struct{}

func (s *PushSender) SendPushNotification() {
	fmt.Println("Push sent!")
}

type NotificationFacade struct {
	emailSender *EmailSender
	smsSender   *SMSSender
	pushSender  *PushSender
}

func (f *NotificationFacade) SendNotificationFacade() {
	fmt.Println("Sending notifications")
	f.emailSender.SendEmail()
	f.smsSender.SendSMS()
	f.pushSender.SendPushNotification()
}

func main() {
	facade := &NotificationFacade{&EmailSender{}, &SMSSender{}, &PushSender{}}
	facade.SendNotificationFacade()
}
