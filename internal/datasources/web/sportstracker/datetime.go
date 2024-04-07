package sportstracker

import (
	"encoding/json"
	"time"
)

type Datetime struct {
	time.Time
}

func (dt *Datetime) UnmarshalJSON(data []byte) error {
	var milliseconds int64
	if err := json.Unmarshal(data, &milliseconds); err != nil {
		return err
	}

	dt.Time = time.UnixMilli(milliseconds)
	return nil
}

func (dt Datetime) MarshalJSON() ([]byte, error) {
	return json.Marshal(dt.Format(time.RFC3339Nano))
}
