package sportstracker

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/microhod/sweaty-swapper/internal/domain"
)

func (c *Client) ListActivities(ctx context.Context) ([]domain.Activity, error) {
	var workouts []Workout
	var offset int

	var err error
	for err == nil {
		var page []Workout
		page, err = c.getWorkoutsPage(ctx, c.pageSize, offset)

		workouts = append(workouts, page...)
		offset += c.pageSize
	}
	if !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("listing workouts: %w", err)
	}

	err = c.attachGPX(ctx, workouts)
	if err != nil {
		return nil, err
	}

	activities := make([]domain.Activity, len(workouts))
	for i, workout := range workouts {
		activities[i], err = workout.ToActivity()
		if err != nil {
			return nil, fmt.Errorf("converting to domain activity: %w", err)
		}
	}
	return activities, nil
}

func (c *Client) getWorkoutsPage(ctx context.Context, limit, offset int) ([]Workout, error) {
	url, err := url.JoinPath(c.baseURL, "/v1/workouts")
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	query := request.URL.Query()
	query.Add("sortonst", "true")
	query.Add("limit", fmt.Sprint(limit))
	query.Add("offset", fmt.Sprint(offset))
	request.URL.RawQuery = query.Encode()

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got non-OK status: %d", response.StatusCode)
	}

	var body workoutsResponse
	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		return nil, err
	}

	if len(body.Payload) == 0 {
		return nil, io.EOF
	}
	return body.Payload, nil
}

type workoutsResponse struct {
	Payload []Workout `json:"payload"`
}

func (c *Client) attachGPX(ctx context.Context, workouts []Workout) error {
	for i, workout := range workouts {
		gpx, err := c.exportGPX(ctx, workout)
		if err != nil {
			return fmt.Errorf("attaching GPX to workout [%s]: %w", workout.Key, err)
		}
		workouts[i].GPX = gpx
	}
	return nil
}

func (c *Client) exportGPX(ctx context.Context, workout Workout) (GPX, error) {
	url, err := url.JoinPath(c.baseURL, "/v1/workout/exportGpx", workout.Key)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("got non-OK status: %d", response.StatusCode)
	}

	buffer := new(bytes.Buffer)
	_, err = io.Copy(buffer, response.Body)
	return buffer.Bytes(), err
}
