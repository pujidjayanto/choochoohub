package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
)

type Client interface {
	FireRequest(ctx context.Context, method, url string, headers map[string]string, params, body, out any) error
}

type fiberClient struct {
	log           *logrus.Logger
	agent         *fiber.Agent
	clientTimeout time.Duration
}

func NewClient(log *logrus.Logger) Client {
	return &fiberClient{
		log:           log,
		agent:         fiber.AcquireAgent(),
		clientTimeout: 2 * time.Second,
	}
}

func (c *fiberClient) FireRequest(
	ctx context.Context,
	method, url string,
	headers map[string]string,
	params, body, out any,
) error {
	req := c.agent.Request()

	req.Header.SetMethod(method)
	req.SetRequestURI(url)

	// set headers
	for k, v := range headers {
		req.Header.Set(k, v)
	}

	// encode body if provided
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("marshal body: %w", err)
		}
		req.SetBody(jsonBody)
		req.Header.Set("Content-Type", "application/json")
	}

	c.log.WithField("headers", headers).WithField("body", req.Body()).WithField("url", url).Info("request info")

	var resp fasthttp.Response
	// context timeout handling
	done := make(chan error, 1)
	go func() {
		err := c.agent.Do(req, &resp)
		done <- err
	}()

	select {
	case <-ctx.Done():
		return fmt.Errorf("request canceled or timeout: %w", ctx.Err())
	case err := <-done:
		if err != nil {
			return fmt.Errorf("fiber agent error: %w", err)
		}
	}

	defer fiber.ReleaseAgent(c.agent)

	bodyBytes := resp.Body()
	status := resp.StatusCode()

	if status >= 400 {
		return fmt.Errorf("http error %d: %s", status, string(bodyBytes))
	}

	if out != nil {
		if err := json.Unmarshal(bodyBytes, out); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}

	return nil
}
