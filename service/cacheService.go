package service

import (
	"SendMailGoBot/model"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetMailPending(m *model.MailDto) (*model.MailDto, error) {
	url := "http//:localhost:8082/mail/pending"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var mailDto model.MailDto
	if err := json.NewDecoder(resp.Body).Decode(&mailDto); err != nil {
		return nil, err
	}

	return &mailDto, nil
}
