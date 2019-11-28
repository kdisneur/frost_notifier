package twilio

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type TwilioHTTPDoer interface {
	Do(*http.Request) (*http.Response, error)
}

type Twilio struct {
	AccountSID string
	AuthToken  string
	Sender     string
	HTTPDoer   TwilioHTTPDoer
}

func New(accountSID string, authToken string, sender string) *Twilio {
	return &Twilio{
		AccountSID: accountSID,
		AuthToken:  authToken,
		Sender:     sender,
		HTTPDoer:   http.DefaultClient,
	}
}

func (t *Twilio) SendMessage(ctx context.Context, recipient string, msg string) error {
	request := buildSendMessageRequest(ctx, t, recipient, msg)

	response, err := t.HTTPDoer.Do(request)
	if err != nil {
		return fmt.Errorf("can't request twilio API: %w", err)
	}

	if response.StatusCode < 200 || response.StatusCode > 299 {
		body, _ := ioutil.ReadAll(response.Body)
		return fmt.Errorf("can't send twilio SMS: %s", body)
	}

	return nil
}

func buildSendMessageRequest(ctx context.Context, t *Twilio, recipient string, msg string) *http.Request {
	form := url.Values{
		"Body": {msg},
		"From": {t.Sender},
		"To":   {recipient},
	}

	body := strings.NewReader(form.Encode())

	sendMessageURL := fmt.Sprintf(
		"https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json",
		t.AccountSID,
	)

	request, _ := http.NewRequestWithContext(ctx, "POST", sendMessageURL, body)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	request.SetBasicAuth(t.AccountSID, t.AuthToken)

	return request
}
