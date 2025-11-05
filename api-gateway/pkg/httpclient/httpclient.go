package httpclient

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Response struct {
	StatusCode int
	Body       []byte
	Headers    map[string]string
	Duration   time.Duration
}

type Client interface {
	Fire(ctx context.Context, method, url string, headers map[string]string, body any) (*Response, error)
}

type fiberClient struct {
	log     *logrus.Logger
	timeout time.Duration
}

func NewClient(log *logrus.Logger) Client {
	return &fiberClient{
		log:     log,
		timeout: 5 * time.Second,
	}
}

func (c *fiberClient) Fire(ctx context.Context, method, url string, headers map[string]string, body any) (*Response, error) {
	startTime := time.Now()

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
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("marshal body: %w", err)
		}
		req.SetBody(bodyBytes)
		req.Header.Set("Content-Type", "application/json")
	}

	deadline, ok := ctx.Deadline()
	if ok {
		agent.Timeout(time.Until(deadline))
	} else {
		agent.Timeout(c.timeout)
	}

	c.log.WithFields(logrus.Fields{
		"method": method,
		"url":    url,
	}).Info("Outgoing HTTP request")

	_ = agent.Parse()

	// Send request
	statusCode, respBody, errs := agent.Bytes()
	duration := time.Since(startTime)

	logFields := logrus.Fields{
		"method":      method,
		"url":         url,
		"status_code": statusCode,
		"duration_ms": duration.Milliseconds(),
		"type":        "outgoing_response",
	}

	if len(errs) > 0 {
		logFields["errors"] = errs
		c.log.WithFields(logFields).Error("HTTP request failed with errors")
		return nil, fmt.Errorf("http request failed: %v", errs)
	}

	return &Response{
		StatusCode: statusCode,
		Body:       respBody,
		Duration:   duration,
	}, nil
}
