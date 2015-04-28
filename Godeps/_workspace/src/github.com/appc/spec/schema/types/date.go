package types

import (
	"encoding/json"
	"fmt"
	"time"
)

// Date wraps time.Time to marshal/unmarshal to/from JSON strings in strict
// accordance with RFC3339
// TODO(jonboulle): golang's implementation seems slightly buggy here;
// according to http://tools.ietf.org/html/rfc3339#section-5.6 , applications
// may choose to separate the date and time with a space instead of a T
// character (for example, `date --rfc-3339` on GNU coreutils) - but this is
// considered an error by go's parser. File a bug?
type Date time.Time

func NewDate(s string) (*Date, error) {
	t, err := time.Parse(time.RFC3339, s)
	if err != nil {
		return nil, fmt.Errorf("bad Date: %v", err)
	}
	d := Date(t)
	return &d, nil
}

func (d Date) String() string {
	return time.Time(d).Format(time.RFC3339)
}

func (d *Date) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	nd, err := NewDate(s)
	if err != nil {
		return err
	}
	*d = *nd
	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}
