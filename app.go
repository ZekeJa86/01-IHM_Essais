// app.go
package main

import (
	"context"
	"fmt"
	"net/smtp"
	"os"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Structure pour le formulaire de contact
type ContactForm struct {
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	Company       string `json:"company"`
	Email         string `json:"email"`
	Country       string `json:"country"`
	PhoneNumber   string `json:"phoneNumber"`
	Message       string `json:"message"`
	AgreeToPolicy bool   `json:"agreeToPolicy"`
}

// Structure pour les réponses
type EmailResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// Configuration SMTP
type SMTPConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	From     string
	To       string
}

// SendEmail - Méthode exportée pour Wails
// Cette méthode sera appelable depuis votre frontend Vue.js
func (a *App) SendEmail(form ContactForm) EmailResponse {
	// Validation
	if form.FirstName == "" || form.LastName == "" || form.Email == "" || form.Message == "" {
		return EmailResponse{
			Success: false,
			Message: "Champs obligatoires manquants",
		}
	}

	if !form.AgreeToPolicy {
		return EmailResponse{
			Success: false,
			Message: "Vous devez accepter la politique de confidentialité",
		}
	}

	// Configuration SMTP depuis les variables d'environnement
	config := SMTPConfig{
		Host:     os.Getenv("SMTP_HOST"),
		Port:     os.Getenv("SMTP_PORT"),
		User:     os.Getenv("EMAIL_USER"),
		Password: os.Getenv("EMAIL_PASS"),
		To:       os.Getenv("EMAIL_TO"),
	}
	config.From = config.User

	// Validation de la configuration
	if config.User == "" || config.Password == "" || config.To == "" {
		return EmailResponse{
			Success: false,
			Message: "⚠️ Configuration email manquante. Vérifiez votre fichier .env",
		}
	}

	if config.Host == "" {
		config.Host = "smtp.gmail.com"
	}
	if config.Port == "" {
		config.Port = "587"
	}

	// Envoi de l'email
	err := a.sendEmailSMTP(config, form)
	if err != nil {
		return EmailResponse{
			Success: false,
			Message: fmt.Sprintf("Erreur lors de l'envoi: %v", err),
		}
	}

	return EmailResponse{
		Success: true,
		Message: "Email envoyé avec succès",
	}
}

// Fonction privée pour l'envoi SMTP
func (a *App) sendEmailSMTP(config SMTPConfig, form ContactForm) error {
	// Authentification SMTP
	auth := smtp.PlainAuth("", config.User, config.Password, config.Host)

	// Construction du corps de l'email
	company := form.Company
	if company == "" {
		company = "Non renseignée"
	}
	phone := form.PhoneNumber
	if phone == "" {
		phone = "Non renseigné"
	}

	subject := fmt.Sprintf("Contact from %s %s", form.FirstName, form.LastName)
	
	body := fmt.Sprintf(`<html>
<head><meta charset="UTF-8"></head>
<body style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto;">
	<h2 style="color: #4f46e5;">🎯 Nouveau message de contact - Go/No-Go Zone</h2>
	
	<div style="background-color: #f3f4f6; padding: 20px; border-radius: 8px; margin: 20px 0;">
		<p style="margin: 10px 0;"><strong>👤 Nom:</strong> %s %s</p>
		<p style="margin: 10px 0;"><strong>📧 Email:</strong> <a href="mailto:%s">%s</a></p>
		<p style="margin: 10px 0;"><strong>🏢 Entreprise:</strong> %s</p>
		<p style="margin: 10px 0;"><strong>📱 Téléphone:</strong> %s</p>
		<p style="margin: 10px 0;"><strong>🌍 Pays:</strong> %s</p>
	</div>
	
	<div style="background-color: #ffffff; padding: 20px; border: 1px solid #e5e7eb; border-radius: 8px;">
		<h3 style="color: #374151; margin-top: 0;">💬 Message:</h3>
		<p style="color: #4b5563; line-height: 1.6; white-space: pre-wrap;">%s</p>
	</div>
	
	<div style="margin-top: 20px; padding: 15px; background-color: #eff6ff; border-left: 4px solid #4f46e5; border-radius: 4px;">
		<p style="margin: 0; color: #1e40af; font-size: 14px;">
			📧 Pour répondre directement, cliquez sur "Répondre" ou écrivez à: %s
		</p>
	</div>
	
	<div style="margin-top: 20px; text-align: center; color: #9ca3af; font-size: 12px;">
		<p>Ce message a été envoyé depuis Go/No-Go Zone</p>
	</div>
</body>
</html>`, 
		form.FirstName, form.LastName, 
		form.Email, form.Email, 
		company, phone, form.Country, 
		form.Message, 
		form.Email)

	// Construction du message avec headers
	message := []byte(
		"From: " + config.From + "\r\n" +
		"To: " + config.To + "\r\n" +
		"Reply-To: " + form.Email + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"\r\n" +
		body,
	)

	// Envoi via SMTP
	addr := config.Host + ":" + config.Port
	err := smtp.SendMail(addr, auth, config.From, []string{config.To}, message)
	if err != nil {
		return fmt.Errorf("échec de l'envoi SMTP: %v", err)
	}

	return nil
}

// Greet returns a greeting for the given name
// (Fonction d'exemple Wails - vous pouvez la garder ou la supprimer)
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}