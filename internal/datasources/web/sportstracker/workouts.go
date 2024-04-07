package sportstracker

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const defaultWorkoutsPageSize = 50

type WorkoutClient struct {
	client   httpClient
	pageSize int
}

func (w *WorkoutClient) ListWorkouts() ([]Workout, error) {
	var workouts []Workout
	var offset int

	var err error
	for err == nil {
		var page []Workout
		page, err = w.getWorkoutsPage(w.pageSize, offset)

		workouts = append(workouts, page...)
		offset += w.pageSize
	}
	if !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("failed to list workouts: %w", err)
	}

	return w.attachGPX(workouts)
}

func (w *WorkoutClient) getWorkoutsPage(limit, offset int) ([]Workout, error) {
	url, err := url.JoinPath(apiserverURL, "/v1/workouts")
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	query := request.URL.Query()
	query.Add("sortonst", "true")
	query.Add("limit", fmt.Sprint(limit))
	query.Add("offset", fmt.Sprint(offset))
	request.URL.RawQuery = query.Encode()

	response, err := w.client.Do(request)
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

type Workout struct {
	Key         string   `json:"workoutKey"`
	Activity    Activity `json:"activityId"`
	Created     Datetime `json:"created"`
	Description string   `json:"description"`
	Photos      []Photo  `json:"photos"`
	Videos      []Video  `json:"videos"`
	GPX         GPX      `json:"gpx"`
}

type Photo struct {
	Key          string `json:"key"`
	URL          string `json:"url"`
	Height       int    `json:"height"`
	Width        int    `json:"width"`
}

type Video struct {
	Key          string `json:"key"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
	Height       int    `json:"height"`
	Width        int    `json:"width"`
}

type GPX []byte

func (w *WorkoutClient) attachGPX(workouts []Workout) ([]Workout, error) {
	for i, workout := range workouts {
		gpx, err := w.ExportGPX(workout)
		if err != nil {
			return nil, err
		}
		workouts[i].GPX = gpx
	}
	return workouts, nil
}

func (w *WorkoutClient) ExportGPX(workout Workout) (GPX, error) {
	url, err := url.JoinPath(apiserverURL, "/v1/workout/exportGpx", workout.Key)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	response, err := w.client.Do(request)
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
