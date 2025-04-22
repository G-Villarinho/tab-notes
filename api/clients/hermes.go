package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/g-villarinho/tab-notes-api/configs"
	"github.com/g-villarinho/tab-notes-api/models"
)

type EmailClient interface {
	SendEmail(ctx context.Context, email *models.Email) error
}

type HemesMailerClient struct {
	apiKey string
	apiURL string
}

func NewHermesMailerClient() *HemesMailerClient {
	return &HemesMailerClient{
		apiKey: configs.Env.Hermes.APIKey,
		apiURL: configs.Env.Hermes.APIURL,
	}
}

func (h *HemesMailerClient) SendEmail(ctx context.Context, notification *models.Email) error {
	jsonBody, err := json.Marshal(notification)
	if err != nil {
		return fmt.Errorf("marshal notification: %w", err)
	}

	url := fmt.Sprintf("%s/email/send", h.apiURL)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(jsonBody))
	if err != nil {
		return fmt.Errorf("build request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-API-Key", h.apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		return fmt.Errorf("email not accepted: status %d", resp.StatusCode)
	}

	return nil
}
