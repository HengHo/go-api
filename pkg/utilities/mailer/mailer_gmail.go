package mailer

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

// SendEmail sends an email using Gmail API
func SendEmail(to, subject, body string) error {
    // Load credentials.json
    ctx := context.Background()
    b, err := os.ReadFile("client_secret_708653700731-hbpkkk380fjn2ald6d9ah3ukva4rq3in.apps.googleusercontent.com.json")
    if err != nil {
        return fmt.Errorf("unable to read client secret file: %v", err)
    }

    // Create a config from credentials
    config, err := google.ConfigFromJSON(b, gmail.GmailSendScope)
    if err != nil {
        return fmt.Errorf("unable to parse client secret file to config: %v", err)
    }

    // Get a token
    client := getClient(config)
    srv, err := gmail.NewService(ctx, option.WithHTTPClient(client))
    if err != nil {
        return fmt.Errorf("unable to retrieve Gmail client: %v", err)
    }

    // Create the email message
    var message gmail.Message
    email := fmt.Sprintf("From: 'me'\r\nTo: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body)
    message.Raw = base64.URLEncoding.EncodeToString([]byte(email))

    // Send the email
    _, err = srv.Users.Messages.Send("me", &message).Do()
    if err != nil {
        return fmt.Errorf("unable to send email: %v", err)
    }

    log.Println("Email sent successfully!")
    return nil
}

// getClient retrieves a token, saves it, and returns the generated client
func getClient(config *oauth2.Config) *http.Client {
    // Define the token file path
    const tokenFile = "token.json"

    // Load token from file if it exists
    token, err := tokenFromFile(tokenFile)
    if err != nil {
        // If token doesn't exist, get a new one from the web
        token = getTokenFromWeb(config)
        saveToken(tokenFile, token)
    }

    // Create an HTTP client using the token
    return config.Client(context.Background(), token)
}

// tokenFromFile reads a token from a file
func tokenFromFile(file string) (*oauth2.Token, error) {
    f, err := os.Open(file)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    var token oauth2.Token
    err = json.NewDecoder(f).Decode(&token)
    return &token, err
}

// saveToken saves a token to a file
func saveToken(file string, token *oauth2.Token) {
    f, err := os.Create(file)
    if err != nil {
        log.Fatalf("Unable to create token file: %v", err)
    }
    defer f.Close()

    err = json.NewEncoder(f).Encode(token)
    if err != nil {
        log.Fatalf("Unable to save token to file: %v", err)
    }
    log.Printf("Token saved to %s\n", file)
}

// getTokenFromWeb requests a token from the web and returns it
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
    authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
    fmt.Printf("Go to the following link in your browser then type the authorization code: \n%v\n", authURL)

    var authCode string
    if _, err := fmt.Scan(&authCode); err != nil {
        log.Fatalf("Unable to read authorization code: %v", err)
    }

    token, err := config.Exchange(context.Background(), authCode)
    if err != nil {
        log.Fatalf("Unable to retrieve token from web: %v", err)
    }
    return token
}