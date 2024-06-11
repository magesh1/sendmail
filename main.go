package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/veyselaksin/gomailer/pkg/mailer"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env file: %v", err)
	}
}

func main() {

	sendMail()

}

func sendMail() {
	auth := mailer.Authentication{
		Username: os.Getenv("SENDERMAIL"),
		Password: os.Getenv("PASSWORD"),
		Host:     "smtp.gmail.com",
		Port:     "587",
	}

	sender := mailer.NewPlainAuth(&auth)

	// 	bodyHTML := `
	//     <html>
	//         <body>
	//             <h1>Hello World</h1>
	//         </body>
	//     </html>
	// `
	subject := "go test mail"

	// message := mailer.NewMessage(subject, bodyHTML)
	message := mailer.NewMessage(subject, "simple test mail sent from go")
	// message.SetTo([]string{"blabla@hotmail.com", "blabla@gmail.com", "blabla@company.com"})
	message.SetTo([]string{os.Getenv("TOMAIL")})
	// message.SetAttachFiles("./src/file")

	if err := sender.SendMail(message); err != nil {
		log.Fatalf("error sending mail: %v", err)
	}

	fmt.Println("mail sent.....")
}
