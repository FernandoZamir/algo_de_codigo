package main

import "fmt"

/*
Reto: Administrar diferentes tipos de notificaciones con un proceso polimorfico
Notificaciones: SMS, Email

*/

// Interface principal
type InterfaceNotificacionfactory interface {
	SendNotification()
	GetSender() ISender
}

// Propiedades de las notficaciones
type ISender interface {
	GetSenderMethod() string
	GetSenderChanel() string
}

// Estructura del mensaje SMS
type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Enviando notificación via SMS")
}

func (SMSNotification) GetSender() ISender {
	// Todas las propiedades
	return SMSNotificationSender{}
}

// Propiedades del envio del mensaje
type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChanel() string {
	return "Twilo"
}

// Email
type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Enviando notificación via Email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChanel() string {
	return "SES Amazon"
}

func getNotificationFactory(Notification string) (InterfaceNotificacionfactory, error) {
	if Notification == "SMS" {
		return &SMSNotification{}, nil
	}

	if Notification == "Email" {
		return &EmailNotification{}, nil
	}

	return nil, fmt.Errorf("Falta el tipo de notificación a enviar")
}

// Tipo de notificación
func SendNotification(f InterfaceNotificacionfactory) {
	f.SendNotification()
}

// Metodo de la notificacion
func getMethod(f InterfaceNotificacionfactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")

	SendNotification(smsFactory)
	SendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)

}
