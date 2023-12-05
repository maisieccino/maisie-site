package music

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"go.uber.org/zap"
)

type SpotifyClient struct {
	logger       zap.Logger
	clientID     string
	clientSecret string

	clientToken string
	tokenExpiry time.Time
	client      *http.Client
}

func NewSpotifyClient(logger zap.Logger, clientID, clientSecret string) *SpotifyClient {
	return &SpotifyClient{
		clientID:     clientID,
		clientSecret: clientSecret,
		logger:       *logger.With(zap.String("component", "spotify")),
		client:       http.DefaultClient,
	}
}

func (c *SpotifyClient) getToken(ctx context.Context) error {
	str := c.clientID + ":" + c.clientSecret
	authCredentials := base64.StdEncoding.EncodeToString([]byte(str))
	data := url.Values(map[string][]string{
		"grant_type": {"client_credentials"},
	})
	resp, err := c.client.Do(&http.Request{
		Method: http.MethodPost,
		Body:   io.NopCloser(strings.NewReader(data.Encode())),
		Header: map[string][]string{
			"Authorization": {"Basic " + authCredentials},
		},
	})

	innerLogger := c.logger.With(
		zap.Int("response_code", resp.StatusCode),
		zap.Error(err),
	)
	if err != nil {
		innerLogger.Warn("error while fetching auth token")
		return err
	}
	if resp.StatusCode != http.StatusOK {
		innerLogger.Warn("did not receive OK response fetching auth token")
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(body))
	}

	respData := struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		ExpiresIn   int    `json:"expires_in"`
	}{}

	d := json.NewDecoder(resp.Body)
	if err := d.Decode(&respData); err != nil {
		innerLogger.With(zap.Error(err)).Warn("error parsing response with token")
	}

	c.clientToken = respData.AccessToken
	c.tokenExpiry = time.Now().Add(time.Duration(respData.ExpiresIn) * time.Second)
	return nil
}
