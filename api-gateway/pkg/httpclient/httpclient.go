package httpclient

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Client interface {
	FireRequest(ctx context.Context, method, url string, headers map[string]string, body, out any) error
}

type fiberClient struct {
	log *logrus.Logger
}

func NewClient(log *logrus.Logger) Client {
	return &fiberClient{
		log: log,
	}
}

func (c *fiberClient) FireRequest(
	ctx context.Context,
	method, url string,
	headers map[string]string,
	body, out any,
) error {
	// Acquire a new agent per request
	agent := fiber.AcquireAgent()
	defer fiber.ReleaseAgent(agent)

	req := agent.Request()
	req.Header.SetMethod(method)
	req.SetRequestURI(url)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal body: %w", err)
		}
		req.SetBody(jsonBody)
		req.Header.Set("Content-Type", "application/json")
	}

	_ = agent.Parse()

	// Logging request
	c.log.WithFields(logrus.Fields{
		"method":  method,
		"url":     url,
		"headers": &req.Header,
		"body":    body,
	}).Info("Outgoing HTTP request")

	// Send request
	statusCode, respBody, errs := agent.Bytes()
	c.log.WithFields(logrus.Fields{
		"statusCode": statusCode,
		"respBody":   respBody,
		"errs":       errs,
	}).Info("Outgoing HTTP response")
	if len(errs) > 0 {
		return fmt.Errorf("http errors: %v", errs)
	}

	if statusCode >= 400 {
		return fmt.Errorf("http error %d: %s", statusCode, string(respBody))
	}

	// Decode response if out provided
	if out != nil {
		if err := json.Unmarshal(respBody, out); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}

	return nil
}
