package property

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBool(t *testing.T) {
	assert.Equal(t, "true", Bool(true).String(), `Bool.String(): "true"`)
	assert.Equal(t, "false", Bool(false).String(), `Bool.String(): "false"`)
	type S struct {
		B Bool `json:"b"`
	}
	t.Run("Marshal", func(t *testing.T) {
		v := S{B: true}
		data, err := json.Marshal(v)
		require.NoError(t, err, "json.Marshal must not return an error")
		assert.Equal(t, []byte(`{"b":"true"}`), data)
	})
	t.Run("Unmarshal", func(t *testing.T) {
		var s S
		assert.NoError(t, json.Unmarshal([]byte(`{ "b": "true" }`), &s), `json.Unmarshal: "true"`)
		assert.True(t, bool(s.B))
		assert.NoError(t, json.Unmarshal([]byte(`{ "b": "false" }`), &s), `json.Unmarshal: "false"`)
		assert.False(t, bool(s.B))
		assert.NoError(t, json.Unmarshal([]byte(`{ "b": true }`), &s), `json.Unmarshal: true`)
		assert.True(t, bool(s.B))
		assert.NoError(t, json.Unmarshal([]byte(`{ "b": false }`), &s), `json.Unmarshal: true`)
		assert.False(t, bool(s.B))
		assert.Error(t, json.Unmarshal([]byte(`{ "b": "foo" }`), &s), `json.Unmarshal: "foo"`)
	})
}

func TestEnumeration(t *testing.T) {
	assert.Equal(t, "a;b;c", Enumeration{"a", "b", "c"}.String(), "Enumeration.String()")
	type S struct {
		E Enumeration `json:"e"`
	}
	t.Run("Marshal", func(t *testing.T) {
		v := S{E: Enumeration{"a", "b", "c"}}
		data, err := json.Marshal(v)
		require.NoError(t, err, "json.Marshal must not return an error")
		assert.Equal(t, []byte(`{"e":"a;b;c"}`), data)
	})
	t.Run("Unmarshal", func(t *testing.T) {
		var s S
		assert.NoError(t, json.Unmarshal([]byte(`{ "e": "a;b;c" }`), &s), `json.Unmarshal: "a;b;c"`)
		assert.Equal(t, Enumeration{"a", "b", "c"}, s.E)
		assert.NoError(t, json.Unmarshal([]byte(`{ "e": ["d", "e", "f"] }`), &s), `json.Unmarshal: ["a", "b", "c"]`)
		assert.Equal(t, Enumeration{"d", "e", "f"}, s.E)
		assert.Error(t, json.Unmarshal([]byte(`{ "e": 1 }`), &s), `json.Unmarshal: 1`)
	})
}

func TestDate(t *testing.T) {
	assert.Equal(t, "2024-02-23", Date(time.Date(2024, 2, 23, 0, 0, 0, 0, time.UTC)).String())
	type S struct {
		D Date `json:"d"`
	}
	t.Run(("MarshalJSON"), func(t *testing.T) {
		v := S{D: Date(time.Date(2024, 2, 23, 0, 0, 0, 0, time.UTC))}
		data, err := json.Marshal(v)
		require.NoError(t, err, "json.Marshal must not return an error")
		assert.Equal(t, []byte(`{"d":"2024-02-23"}`), data)
	})
	t.Run("UnmarshalJSON", func(t *testing.T) {
		var s S
		assert.NoError(t, json.Unmarshal([]byte(`{ "d": "2024-02-23" }`), &s), `json.Unmarshal: "2024-02-23"`)
		assert.Equal(t, Date(time.Date(2024, 2, 23, 0, 0, 0, 0, time.UTC)), s.D)
		assert.NoError(t, json.Unmarshal([]byte(`{ "d": "2024-02-22T00:00:00Z" }`), &s), `json.Unmarshal: "2024-02-23T00:00:00Z"`)
		assert.Equal(t, Date(time.Date(2024, 2, 22, 0, 0, 0, 0, time.UTC)), s.D)
		assert.Error(t, json.Unmarshal([]byte(`{ "d": 1 }`), &s), `json.Unmarshal: 1`)
	})
	t.Run("Wrappers", func(t *testing.T) {
		now := time.Now()
		d := Date(now)

		// Add
		duration := 48 * time.Hour
		assert.Equal(t, time.Time(d).Add(duration), d.Add(duration))

		// AddDate
		years, months, days := 1, 2, 3
		assert.Equal(t, time.Time(d).AddDate(years, months, days), d.AddDate(years, months, days))

		// After
		past := now.Add(-duration)
		assert.Equal(t, now.After(past), d.After(past))

		// AppendFormat
		layout := "2006-01-02 15:04:05"
		assert.Equal(t, now.AppendFormat([]byte{}, layout), d.AppendFormat([]byte{}, layout))

		// Before
		assert.Equal(t, now.Before(past), d.Before(past))

		// Clock
		hour, min, sec := now.Clock()
		assertHour, assertMin, assertSec := d.Clock()
		assert.Equal(t, hour, assertHour)
		assert.Equal(t, min, assertMin)
		assert.Equal(t, sec, assertSec)

		// Compare
		assert.Equal(t, now.Compare(past), d.Compare(past))

		// Date
		year, month, day := now.Date()
		assertYear, assertMonth, assertDay := d.Date()
		assert.Equal(t, year, assertYear)
		assert.Equal(t, month, assertMonth)
		assert.Equal(t, day, assertDay)

		// Day
		assert.Equal(t, now.Day(), d.Day())

		// Equal
		assert.Equal(t, now.Equal(now), d.Equal(now))

		// Format
		assert.Equal(t, now.Format(layout), d.Format(layout))

		// GoString
		assert.Equal(t, now.GoString(), d.GoString())

		// GobDecode and GobEncode
		encoded, err := d.GobEncode()
		require.NoError(t, err)
		var decoded Date
		require.NoError(t, decoded.GobDecode(encoded))
		// assert.Equal(t, d, decoded) // always fails because of wall, ext
		assert.Zero(t, decoded.Compare(now))

		// Hour
		assert.Equal(t, now.Hour(), d.Hour())

		// ISOWeek
		isoYear, isoWeek := now.ISOWeek()
		assertIsoYear, assertIsoWeek := d.ISOWeek()
		assert.Equal(t, isoYear, assertIsoYear)
		assert.Equal(t, isoWeek, assertIsoWeek)

		// In
		location := time.FixedZone("Test", 3600)
		assert.Equal(t, now.In(location), d.In(location))

		// IsDST
		assert.Equal(t, now.IsDST(), d.IsDST())

		// IsZero
		zeroDate := Date(time.Time{})
		assert.Equal(t, zeroDate.IsZero(), true)

		// Local
		assert.Equal(t, now.Local(), d.Local())

		// Location
		assert.Equal(t, now.Location(), d.Location())

		// MarshalBinary and UnmarshalBinary
		binary, err := d.MarshalBinary()
		require.NoError(t, err)
		var binaryDecoded Date
		require.NoError(t, binaryDecoded.UnmarshalBinary(binary))
		// assert.Equal(t, d, binaryDecoded) // always fails because of wall, ext
		assert.Zero(t, binaryDecoded.Compare(now))
		require.Error(t, binaryDecoded.UnmarshalBinary(append(binary, 1)))

		// MarshalText and UnmarshalText
		text, err := d.MarshalText()
		require.NoError(t, err)
		var textDecoded Date
		require.NoError(t, textDecoded.UnmarshalText(text))
		// assert.Equal(t, d, textDecoded) // always fails because of wall, ext
		assert.Zero(t, textDecoded.Compare(now))
		require.Error(t, textDecoded.UnmarshalText(append(text, 1)))

		// Minute
		assert.Equal(t, now.Minute(), d.Minute())

		// Month
		assert.Equal(t, now.Month(), d.Month())

		// Nanosecond
		assert.Equal(t, now.Nanosecond(), d.Nanosecond())

		// Round
		roundDuration := time.Hour
		assert.Equal(t, now.Round(roundDuration), d.Round(roundDuration))

		// Second
		assert.Equal(t, now.Second(), d.Second())

		// Sub
		assert.Equal(t, now.Sub(past), d.Sub(past))

		// Truncate
		truncateDuration := time.Hour
		assert.Equal(t, now.Truncate(truncateDuration), d.Truncate(truncateDuration))

		// UTC
		assert.Equal(t, now.UTC(), d.UTC())

		// Unix
		assert.Equal(t, now.Unix(), d.Unix())

		// UnixMicro
		assert.Equal(t, now.UnixMicro(), d.UnixMicro())

		// UnixMilli
		assert.Equal(t, now.UnixMilli(), d.UnixMilli())

		// UnixNano
		assert.Equal(t, now.UnixNano(), d.UnixNano())

		// Weekday
		assert.Equal(t, now.Weekday(), d.Weekday())

		// Year
		assert.Equal(t, now.Year(), d.Year())

		// YearDay
		assert.Equal(t, now.YearDay(), d.YearDay())

		// Zone
		zoneName, offset := now.Zone()
		assertZoneName, assertOffset := d.Zone()
		assert.Equal(t, zoneName, assertZoneName)
		assert.Equal(t, offset, assertOffset)

		// ZoneBounds
		zoneStart, zoneEnd := now.ZoneBounds()
		assertZoneStart, assertZoneEnd := d.ZoneBounds()
		assert.Equal(t, zoneStart, assertZoneStart)
		assert.Equal(t, zoneEnd, assertZoneEnd)
	})
}

func TestDateTime(t *testing.T) {
	assert.Equal(t, "2024-02-23T16:14:31.229Z", DateTime(time.Date(2024, 2, 23, 16, 14, 31, 229000000, time.UTC)).String())
	type S struct {
		D DateTime `json:"d"`
	}
	t.Run(("MarshalJSON"), func(t *testing.T) {
		v := S{D: DateTime(time.Date(2024, 2, 23, 16, 14, 31, 229000000, time.UTC))}
		data, err := json.Marshal(v)
		require.NoError(t, err, "json.Marshal must not return an error")
		assert.Equal(t, []byte(`{"d":"2024-02-23T16:14:31.229Z"}`), data)
	})
	t.Run("UnmarshalJSON", func(t *testing.T) {
		var s S
		assert.NoError(t, json.Unmarshal([]byte(`{ "d": "2024-02-23T16:14:31.229Z" }`), &s), `json.Unmarshal: "2024-02-23T16:14:31.229Z"`)
		assert.Equal(t, DateTime(time.Date(2024, 2, 23, 16, 14, 31, 229000000, time.UTC)), s.D)
		assert.Equal(t, DateTime(time.Date(2024, 2, 23, 16, 14, 31, 229000000, time.UTC)), s.D)
		assert.Error(t, json.Unmarshal([]byte(`{ "d": "2024-02-22" }`), &s), `json.Unmarshal: "2024-02-22"`)
		assert.Error(t, json.Unmarshal([]byte(`{ "d": 1 }`), &s), `json.Unmarshal: 1`)
	})
	t.Run("Wrappers", func(t *testing.T) {
		now := time.Now()
		d := DateTime(now)

		// Add
		duration := 48 * time.Hour
		assert.Equal(t, time.Time(d).Add(duration), d.Add(duration))

		// AddDate
		years, months, days := 1, 2, 3
		assert.Equal(t, time.Time(d).AddDate(years, months, days), d.AddDate(years, months, days))

		// After
		past := now.Add(-duration)
		assert.Equal(t, now.After(past), d.After(past))

		// AppendFormat
		layout := "2006-01-02 15:04:05"
		assert.Equal(t, now.AppendFormat([]byte{}, layout), d.AppendFormat([]byte{}, layout))

		// Before
		assert.Equal(t, now.Before(past), d.Before(past))

		// Clock
		hour, min, sec := now.Clock()
		assertHour, assertMin, assertSec := d.Clock()
		assert.Equal(t, hour, assertHour)
		assert.Equal(t, min, assertMin)
		assert.Equal(t, sec, assertSec)

		// Compare
		assert.Equal(t, now.Compare(past), d.Compare(past))

		// Date
		year, month, day := now.Date()
		assertYear, assertMonth, assertDay := d.Date()
		assert.Equal(t, year, assertYear)
		assert.Equal(t, month, assertMonth)
		assert.Equal(t, day, assertDay)

		// Day
		assert.Equal(t, now.Day(), d.Day())

		// Equal
		assert.Equal(t, now.Equal(now), d.Equal(now))

		// Format
		assert.Equal(t, now.Format(layout), d.Format(layout))

		// GoString
		assert.Equal(t, now.GoString(), d.GoString())

		// GobDecode and GobEncode
		encoded, err := d.GobEncode()
		require.NoError(t, err)
		var decoded DateTime
		require.NoError(t, decoded.GobDecode(encoded))
		// assert.Equal(t, d, decoded) // always fails because of wall, ext
		assert.Zero(t, decoded.Compare(now))

		// Hour
		assert.Equal(t, now.Hour(), d.Hour())

		// ISOWeek
		isoYear, isoWeek := now.ISOWeek()
		assertIsoYear, assertIsoWeek := d.ISOWeek()
		assert.Equal(t, isoYear, assertIsoYear)
		assert.Equal(t, isoWeek, assertIsoWeek)

		// In
		location := time.FixedZone("Test", 3600)
		assert.Equal(t, now.In(location), d.In(location))

		// IsDST
		assert.Equal(t, now.IsDST(), d.IsDST())

		// IsZero
		zeroDate := DateTime(time.Time{})
		assert.Equal(t, zeroDate.IsZero(), true)

		// Local
		assert.Equal(t, now.Local(), d.Local())

		// Location
		assert.Equal(t, now.Location(), d.Location())

		// MarshalBinary and UnmarshalBinary
		binary, err := d.MarshalBinary()
		require.NoError(t, err)
		var binaryDecoded DateTime
		require.NoError(t, binaryDecoded.UnmarshalBinary(binary))
		// assert.Equal(t, d, binaryDecoded) // always fails because of wall, ext
		assert.Zero(t, binaryDecoded.Compare(now))
		require.Error(t, binaryDecoded.UnmarshalBinary(append(binary, 1)))

		// MarshalText and UnmarshalText
		text, err := d.MarshalText()
		require.NoError(t, err)
		var textDecoded DateTime
		require.NoError(t, textDecoded.UnmarshalText(text))
		// assert.Equal(t, d, textDecoded) // always fails because of wall, ext
		assert.Zero(t, textDecoded.Compare(now))
		require.Error(t, textDecoded.UnmarshalText(append(text, 1)))

		// Minute
		assert.Equal(t, now.Minute(), d.Minute())

		// Month
		assert.Equal(t, now.Month(), d.Month())

		// Nanosecond
		assert.Equal(t, now.Nanosecond(), d.Nanosecond())

		// Round
		roundDuration := time.Hour
		assert.Equal(t, now.Round(roundDuration), d.Round(roundDuration))

		// Second
		assert.Equal(t, now.Second(), d.Second())

		// Sub
		assert.Equal(t, now.Sub(past), d.Sub(past))

		// Truncate
		truncateDuration := time.Hour
		assert.Equal(t, now.Truncate(truncateDuration), d.Truncate(truncateDuration))

		// UTC
		assert.Equal(t, now.UTC(), d.UTC())

		// Unix
		assert.Equal(t, now.Unix(), d.Unix())

		// UnixMicro
		assert.Equal(t, now.UnixMicro(), d.UnixMicro())

		// UnixMilli
		assert.Equal(t, now.UnixMilli(), d.UnixMilli())

		// UnixNano
		assert.Equal(t, now.UnixNano(), d.UnixNano())

		// Weekday
		assert.Equal(t, now.Weekday(), d.Weekday())

		// Year
		assert.Equal(t, now.Year(), d.Year())

		// YearDay
		assert.Equal(t, now.YearDay(), d.YearDay())

		// Zone
		zoneName, offset := now.Zone()
		assertZoneName, assertOffset := d.Zone()
		assert.Equal(t, zoneName, assertZoneName)
		assert.Equal(t, offset, assertOffset)

		// ZoneBounds
		zoneStart, zoneEnd := now.ZoneBounds()
		assertZoneStart, assertZoneEnd := d.ZoneBounds()
		assert.Equal(t, zoneStart, assertZoneStart)
		assert.Equal(t, zoneEnd, assertZoneEnd)
	})
}
