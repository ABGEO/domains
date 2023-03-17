package services

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"

	"github.com/abgeo/domains/config"
	"github.com/pkg/errors"
)

var errInvalidRecaptchaResponse = errors.New("invalid reCAPTCHA response")

type RecaptchaService struct {
	config *config.Config
}

func NewRecaptchaService(conf *config.Config) *RecaptchaService {
	return &RecaptchaService{config: conf}
}

func (service *RecaptchaService) Validate(remoteIP, token string) error {
	var recaptchaResponse struct {
		Success bool `json:"success"`
	}

	response, err := http.PostForm(
		"https://www.google.com/recaptcha/api/siteverify",
		url.Values{"secret": {service.config.Recaptcha.SecretKey}, "remoteip": {remoteIP}, "response": {token}},
	)
	if err != nil {
		return errors.Wrap(err, "unable to call reCAPTCHA API")
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "unable to read reCAPTCHA API response")
	}

	if err = json.Unmarshal(body, &recaptchaResponse); err != nil {
		return errors.Wrap(err, "unable to read reCAPTCHA API response")
	}

	if !recaptchaResponse.Success {
		return errInvalidRecaptchaResponse
	}

	return nil
}
