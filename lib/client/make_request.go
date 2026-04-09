package client

import (
	"bytes"
	"cmp"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const defaultHTTPTimeout = 30 * time.Second

var httpClient = &http.Client{Timeout: defaultHTTPTimeout}

func MakeRequest[TVariables any, TData any](ctx context.Context, token string, payload GraphQLRequest[TVariables], endpoint ...string) (*GraphQLResponse[TData], error) {
	url := BackendEndpoint
	if len(endpoint) > 0 {
		url = cmp.Or(endpoint[0], BackendEndpoint)
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal GraphQL request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Authorization", token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://octopus.energy/")
	req.Header.Set("Referer", "https://octopus.energy/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:147.0) Gecko/20100101 Firefox/147.0")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read HTTP response body: %w", err)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		trimmed := string(resBody)
		if len(trimmed) > 4*1024 {
			trimmed = trimmed[:4*1024] + "..."
		}
		return nil, fmt.Errorf("graphql request failed with status %s body=%q", resp.Status, trimmed)
	}

	var result GraphQLResponse[TData]
	if err := json.Unmarshal(resBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func MakeUnauthenticatedRequest[TVariables any, TData any](ctx context.Context, payload GraphQLRequest[TVariables], endpoint ...string) (*GraphQLResponse[TData], error) {
	url := ApiOctopusEndpoint
	if len(endpoint) > 0 {
		url = cmp.Or(endpoint[0], ApiOctopusEndpoint)
	}
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal GraphQL request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://octopus.energy/")
	req.Header.Set("Referer", "https://octopus.energy/")
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:147.0) Gecko/20100101 Firefox/147.0")

	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read HTTP response body: %w", err)
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		trimmed := string(resBody)
		if len(trimmed) > 4*1024 {
			trimmed = trimmed[:4*1024] + "..."
		}
		return nil, fmt.Errorf("graphql request failed with status %s body=%q", resp.Status, trimmed)
	}

	var result GraphQLResponse[TData]
	if err := json.Unmarshal(resBody, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

