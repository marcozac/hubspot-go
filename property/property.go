package property

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/marcozac/hubspot-go/util"
)

type Bool bool

func (b Bool) String() string {
	if b {
		return "true"
	}
	return "false"
}

var _ json.Marshaler = Bool(true)

func (b Bool) MarshalJSON() ([]byte, error) {
	return util.MarshalStringerAsJSON(b)
}

var (
	_ json.Unmarshaler = (*Bool)(nil)

	trueStringBytes  = []byte(`"true"`)
	falseStringBytes = []byte(`"false"`)
)

func (b *Bool) UnmarshalJSON(data []byte) error {
	switch {
	case bytes.Equal(data, trueStringBytes):
		*b = true
	case bytes.Equal(data, falseStringBytes):
		*b = false
	default:
		// Try to unmarshal as a bool.
		return json.Unmarshal(data, (*bool)(b))
	}
	return nil
}

type Enumeration []string

func (e Enumeration) String() string {
	return strings.Join(e, ";")
}

var _ json.Marshaler = Enumeration{}

func (e Enumeration) MarshalJSON() ([]byte, error) {
	return util.MarshalStringerAsJSON(e)
}

var _ json.Unmarshaler = (*Enumeration)(nil)

func (e *Enumeration) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		// Try to unmarshal as a slice of strings.
		return json.Unmarshal(data, (*[]string)(e))
	}
	for _, v := range strings.Split(s, ";") {
		*e = append(*e, v)
	}
	return nil
}

// Date is a time.Time compatible with HubSpot's date properties.
//
// It returns a string in the format "2006-01-02" when formatted with String()
// and marshals to and from JSON as a string in the same format.
// The other methods are (almost) wrappers around the time.Time ones.
type Date time.Time

func (d Date) String() string {
	return time.Time(d).Format(time.DateOnly)
}

func (d Date) MarshalJSON() ([]byte, error) {
	return util.MarshalStringerAsJSON(d)
}

func (d *Date) UnmarshalJSON(data []byte) error {
	if err := d.unmarshalJSON(data); err != nil {
		// Try to unmarshal as a time.Time (strict RFC 3339)
		return json.Unmarshal(data, (*time.Time)(d))
	}
	return nil
}

func (d *Date) unmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	t, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) Add(dn time.Duration) time.Time {
	return time.Time(d).Add(dn)
}

func (d Date) AddDate(years int, months int, days int) time.Time {
	return time.Time(d).AddDate(years, months, days)
}

func (d Date) After(u time.Time) bool {
	return time.Time(d).After(u)
}

func (d Date) AppendFormat(b []byte, layout string) []byte {
	return time.Time(d).AppendFormat(b, layout)
}

func (d Date) Before(u time.Time) bool {
	return time.Time(d).Before(u)
}

func (d Date) Clock() (hour int, min int, sec int) {
	return time.Time(d).Clock()
}

func (d Date) Compare(u time.Time) int {
	return time.Time(d).Compare(u)
}

func (d Date) Date() (year int, month time.Month, day int) {
	return time.Time(d).Date()
}

func (d Date) Day() int {
	return time.Time(d).Day()
}

func (d Date) Equal(u time.Time) bool {
	return time.Time(d).Equal(u)
}

func (d Date) Format(layout string) string {
	return time.Time(d).Format(layout)
}

func (d Date) GoString() string {
	return time.Time(d).GoString()
}

func (d *Date) GobDecode(data []byte) error {
	return d.UnmarshalBinary(data)
}

func (d Date) GobEncode() ([]byte, error) {
	return d.MarshalBinary()
}

func (d Date) Hour() int {
	return time.Time(d).Hour()
}

func (d Date) ISOWeek() (year int, week int) {
	return time.Time(d).ISOWeek()
}

func (d Date) In(loc *time.Location) time.Time {
	return time.Time(d).In(loc)
}

func (d Date) IsDST() bool {
	return time.Time(d).IsDST()
}

func (d Date) IsZero() bool {
	return time.Time(d).IsZero()
}

func (d Date) Local() time.Time {
	return time.Time(d).Local()
}

func (d Date) Location() *time.Location {
	return time.Time(d).Location()
}

func (d Date) MarshalBinary() ([]byte, error) {
	return time.Time(d).MarshalBinary()
}

func (d Date) MarshalText() ([]byte, error) {
	return time.Time(d).MarshalText()
}

func (d Date) Minute() int {
	return time.Time(d).Minute()
}

func (d Date) Month() time.Month {
	return time.Time(d).Month()
}

func (d Date) Nanosecond() int {
	return time.Time(d).Nanosecond()
}

func (d Date) Round(dn time.Duration) time.Time {
	return time.Time(d).Round(dn)
}

func (d Date) Second() int {
	return time.Time(d).Second()
}

func (d Date) Sub(u time.Time) time.Duration {
	return time.Time(d).Sub(u)
}

func (d Date) Truncate(dn time.Duration) time.Time {
	return time.Time(d).Truncate(dn)
}

func (d Date) UTC() time.Time {
	return time.Time(d).UTC()
}

func (d Date) Unix() int64 {
	return time.Time(d).Unix()
}

func (d Date) UnixMicro() int64 {
	return time.Time(d).UnixMicro()
}

func (d Date) UnixMilli() int64 {
	return time.Time(d).UnixMilli()
}

func (d Date) UnixNano() int64 {
	return time.Time(d).UnixNano()
}

func (d *Date) UnmarshalBinary(data []byte) error {
	var t time.Time
	if err := (&t).UnmarshalBinary(data); err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d *Date) UnmarshalText(data []byte) error {
	var t time.Time
	if err := (&t).UnmarshalText(data); err != nil {
		return err
	}
	*d = Date(t)
	return nil
}

func (d Date) Weekday() time.Weekday {
	return time.Time(d).Weekday()
}

func (d Date) Year() int {
	return time.Time(d).Year()
}

func (d Date) YearDay() int {
	return time.Time(d).YearDay()
}

func (d Date) Zone() (name string, offset int) {
	return time.Time(d).Zone()
}

func (d Date) ZoneBounds() (start time.Time, end time.Time) {
	return time.Time(d).ZoneBounds()
}
