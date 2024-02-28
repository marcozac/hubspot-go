package hubspot

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/marcozac/hubspot-go/util"
)

const (
	PropertyTypeString      PropertyType = "string"
	PropertyTypePhoneNumber PropertyType = "phone_number" // Alias for "string"
	PropertyTypeNumber      PropertyType = "number"
	PropertyTypeDate        PropertyType = "date"
	PropertyTypeDateTime    PropertyType = "datetime"
	PropertyTypeEnumeration PropertyType = "enumeration"
	PropertyTypeBool        PropertyType = "bool"

	PropertyFieldTypeText                PropertyFieldType = "text"
	PropertyFieldTypeTextarea            PropertyFieldType = "textarea"
	PropertyFieldTypeDate                PropertyFieldType = "date"
	PropertyFieldTypeFile                PropertyFieldType = "file"
	PropertyFieldTypeNumber              PropertyFieldType = "number"
	PropertyFieldTypeSelect              PropertyFieldType = "select"
	PropertyFieldTypeRadio               PropertyFieldType = "radio"
	PropertyFieldTypeCheckbox            PropertyFieldType = "checkbox"
	PropertyFieldTypeBooleanCheckbox     PropertyFieldType = "booleancheckbox"
	PropertyFieldTypeCalculationEquation PropertyFieldType = "calculation_equation"

	PropertyReferencedObjectTypeOwner    PropertyReferencedObjectType = "OWNER"
	PropertyReferencedObjectTypeContact  PropertyReferencedObjectType = "CONTACT"
	PropertyReferencedObjectTypeCompany  PropertyReferencedObjectType = "COMPANY"
	PropertyReferencedObjectTypeDeal     PropertyReferencedObjectType = "DEAL"
	PropertyReferencedObjectTypeSequence PropertyReferencedObjectType = "SEQUENCE"
)

type (
	// Property is a HubSpot property as returned by the properties API.
	//
	// Do not confuse it with Object properties, which are the properties of a
	// HubSpot object (e.g. a contact, a company, a deal).
	Property struct {
		// Name is the internal name of the property as used in the API.
		Name string `json:"name"`

		// Label is the label of the property as displayed in the HubSpot UI.
		Label string `json:"label"`

		// GroupName is the name of the group the property belongs to.
		GroupName string `json:"groupName"`

		// Type is the type of the property, e.g. "string", "number", "enumeration".
		Type PropertyType `json:"type"`

		// FieldType is the type of the field, e.g. "text", "select", "checkbox".
		FieldType PropertyFieldType `json:"fieldType"`

		// ReferencedObjectType is the type of the object the property references.
		//
		// From HubSpot's documentation:
		//
		// Should be set to 'OWNER' when 'externalOptions' is true, which causes
		// the property to dynamically pull option values from the current
		// HubSpot users.
		ReferencedObjectType PropertyReferencedObjectType `json:"referencedObjectType,omitempty"`

		// Description is a description of the property.
		Description string `json:"description,omitempty"`

		// DisplayOrder is the order in which the property is displayed.
		DisplayOrder int `json:"displayOrder,omitempty"`

		// Options is a list of options for enumeration properties.
		Options []PropertyOption `json:"options,omitempty"`

		// ExternalOptions is true if the property has external options.
		ExternalOptions bool `json:"externalOptions,omitempty"`

		// Calculated is true if the property is calculated.
		Calculated bool `json:"calculated,omitempty"`

		// CalculationFormula is the formula used to calculate the property.
		//
		// From HubSpot's documentation:
		//
		// Applicable only for 'enumeration' type properties. Should be set to
		// true in conjunction with a 'referencedObjectType' of 'OWNER'.
		// Otherwise false.
		CalculationFormula string `json:"calculationFormula,omitempty"`

		// HasUniqueValue is true if the property has unique values.
		HasUniqueValue bool `json:"hasUniqueValue,omitempty"`

		// Hidden is true if the property is hidden.
		Hidden bool `json:"hidden,omitempty"`

		// ShowCurrencySymbol is true if the currency symbol is shown.
		ShowCurrencySymbol bool `json:"showCurrencySymbol,omitempty"`

		// FormField is true if the property is a form field.
		FormField bool `json:"formField,omitempty"`

		// Archived is true if the property is archived.
		Archived bool `json:"archived,omitempty"`

		// ModificationMetadata is the modification metadata of the property.
		ModificationMetadata PropertyModificationMetadata `json:"modificationMetadata,omitempty"`

		// CreatedUserId is the ID of the user who created the property.
		CreatedUserId string `json:"createdUserId,omitempty"`

		// UpdatedAt is the time the property was last updated.
		UpdatedAt DateTime `json:"updatedAt,omitempty"`

		// UpdatedUserId is the ID of the user who last updated the property.
		UpdatedUserId string `json:"updatedUserId,omitempty"`

		// CreatedAt is the time the property was created.
		CreatedAt DateTime `json:"createdAt,omitempty"`
	}

	// PropertyType is the type of a property.
	PropertyType string

	// PropertyFieldType is the type of the field of a property.
	PropertyFieldType string

	// PropertyReferencedObjectType is the type of the object a property references.
	// It should be set to 'OWNER' when 'externalOptions' is true.
	PropertyReferencedObjectType string

	// PropertyOption is an option for an enumeration property.
	PropertyOption struct {
		// Label is the label of the option.
		Label string `json:"label"`

		// Value is the value of the option.
		Value string `json:"value"`

		// Hidden is true if the option is hidden.
		Hidden bool `json:"hidden"`

		// Description is a description of the option.
		Description string `json:"description,omitempty"`

		// DisplayOrder is the order in which the option is displayed.
		DisplayOrder int `json:"displayOrder,omitempty"`
	}

	// PropertyModificationMetadata is the modification metadata of a property.
	PropertyModificationMetadata struct {
		// Archivable is true if the property is archivable.
		Archivable bool `json:"archivable"`

		// ReadOnlyDefinition is true if the property definition is read-only.
		ReadOnlyDefinition bool `json:"readOnlyDefinition"`

		// ReadOnlyValue is true if the property value is read-only.
		ReadOnlyValue bool `json:"readOnlyValue"`
	}

	PropertyGroup struct {
		// Name is the internal name of the group as used in the API.
		Name string `json:"name"`

		// Label is the label of the group as displayed in the HubSpot UI.
		Label string `json:"label"`

		// DisplayOrder is the order in which the group is displayed.
		DisplayOrder int `json:"displayOrder,omitempty"`

		// Archived is true if the group is archived.
		Archived bool `json:"archived,omitempty"`
	}

	PropertyWithHistory struct {
		SourceID        string   `json:"sourceId,omitempty"`
		SourceLabel     string   `json:"sourceLabel,omitempty"`
		SourceType      string   `json:"sourceType,omitempty"`
		UpdatedByUserID int64    `json:"updatedByUserId,omitempty"`
		Value           string   `json:"value,omitempty"`
		Timestamp       DateTime `json:"timestamp,omitempty"`
	}
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

type Int int

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

var _ json.Marshaler = Int(0)

func (i Int) MarshalJSON() ([]byte, error) {
	return util.MarshalStringerAsJSON(i)
}

var _ json.Unmarshaler = (*Int)(nil)

func (i *Int) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		// Try to unmarshal as an int.
		return json.Unmarshal(data, (*int)(i))
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*i = Int(v)
	return nil
}

type Int32 int32

func (i Int32) String() string {
	return strconv.FormatInt(int64(i), 10)
}

var _ json.Marshaler = Int32(0)

func (i Int32) MarshalJSON() ([]byte, error) {
	return util.MarshalStringerAsJSON(i)
}

var _ json.Unmarshaler = (*Int32)(nil)

func (i *Int32) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		// Try to unmarshal as an int32.
		return json.Unmarshal(data, (*int32)(i))
	}
	v, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return err
	}
	*i = Int32(v)
	return nil
}

type Int64 int64

func (i Int64) String() string {
	return strconv.FormatInt(int64(i), 10)
}

var _ json.Marshaler = Int64(0)

func (i Int64) MarshalJSON() ([]byte, error) {
	return util.MarshalStringerAsJSON(i)
}

var _ json.Unmarshaler = (*Int64)(nil)

func (i *Int64) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		// Try to unmarshal as an int64.
		return json.Unmarshal(data, (*int64)(i))
	}
	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return err
	}
	*i = Int64(v)
	return nil
}

type Float32 float32

func (f Float32) String() string {
	return strconv.FormatFloat(float64(f), 'f', -1, 32)
}

var _ json.Marshaler = Float32(0)

func (f Float32) MarshalJSON() ([]byte, error) {
	return util.MarshalStringerAsJSON(f)
}

var _ json.Unmarshaler = (*Float32)(nil)

func (f *Float32) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		// Try to unmarshal as a float32.
		return json.Unmarshal(data, (*float32)(f))
	}
	v, err := strconv.ParseFloat(s, 32)
	if err != nil {
		return err
	}
	*f = Float32(v)
	return nil
}

type Float64 float64

func (f Float64) String() string {
	return strconv.FormatFloat(float64(f), 'f', -1, 64)
}

var _ json.Marshaler = Float64(0)

func (f Float64) MarshalJSON() ([]byte, error) {
	return util.MarshalStringerAsJSON(f)
}

var _ json.Unmarshaler = (*Float64)(nil)

func (f *Float64) UnmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		// Try to unmarshal as a float64.
		return json.Unmarshal(data, (*float64)(f))
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}
	*f = Float64(v)
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
	*e = strings.Split(s, ";")
	return nil
}

// Date is a time.Time compatible with HubSpot's date properties.
//
// It returns a string in the format "2006-01-02" when formatted with String()
// and marshals to and from JSON as a string in the same format.
//
// The other methods are essentially wrappers around the time.Time ones.
//
// In order to respect the JSON "omitempty" tag for the zero value, Date must
// be a pointer in the struct definition. See [LineItemDefaultProperties] for
// an example.
type Date time.Time

func (d Date) String() string {
	return time.Time(d).Format(time.DateOnly)
}

var _ json.Marshaler = (*Date)(nil)

func (d *Date) MarshalJSON() ([]byte, error) {
	return util.MarshalStringerAsJSON(d)
}

var _ json.Unmarshaler = (*Date)(nil)

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

// DateTime is a time.Time compatible with HubSpot's date properties.
//
// It returns a string in the format "2006-01-02T15:04:05.999Z07:00" (ISO 8601
// with millisecond precision) when formatted with String() and marshals to and
// from JSON as a string in the same format.
//
// The other methods are essentially wrappers around the time.Time ones.
//
// In order to respect the JSON "omitempty" tag for the zero value, DateTime
// must be a pointer in the struct definition. See [ContactDefaultProperties]
// for an example.
type DateTime time.Time

func (d DateTime) String() string {
	return time.Time(d).Format(util.RFC3339Milli)
}

var _ json.Marshaler = (*DateTime)(nil)

func (d *DateTime) MarshalJSON() ([]byte, error) {
	return util.MarshalStringerAsJSON(d)
}

var _ json.Unmarshaler = (*DateTime)(nil)

func (d *DateTime) UnmarshalJSON(data []byte) error {
	if err := d.unmarshalJSON(data); err != nil {
		// Try to unmarshal as a time.Time (strict RFC 3339)
		return json.Unmarshal(data, (*time.Time)(d))
	}
	return nil
}

func (d *DateTime) unmarshalJSON(data []byte) error {
	s, err := strconv.Unquote(string(data))
	if err != nil {
		return err
	}
	t, err := time.Parse(util.RFC3339Milli, s)
	if err != nil {
		return err
	}
	*d = DateTime(t)
	return nil
}

func (d DateTime) Add(dn time.Duration) time.Time {
	return time.Time(d).Add(dn)
}

func (d DateTime) AddDate(years int, months int, days int) time.Time {
	return time.Time(d).AddDate(years, months, days)
}

func (d DateTime) After(u time.Time) bool {
	return time.Time(d).After(u)
}

func (d DateTime) AppendFormat(b []byte, layout string) []byte {
	return time.Time(d).AppendFormat(b, layout)
}

func (d DateTime) Before(u time.Time) bool {
	return time.Time(d).Before(u)
}

func (d DateTime) Clock() (hour int, min int, sec int) {
	return time.Time(d).Clock()
}

func (d DateTime) Compare(u time.Time) int {
	return time.Time(d).Compare(u)
}

func (d DateTime) Date() (year int, month time.Month, day int) {
	return time.Time(d).Date()
}

func (d DateTime) Day() int {
	return time.Time(d).Day()
}

func (d DateTime) Equal(u time.Time) bool {
	return time.Time(d).Equal(u)
}

func (d DateTime) Format(layout string) string {
	return time.Time(d).Format(layout)
}

func (d DateTime) GoString() string {
	return time.Time(d).GoString()
}

func (d *DateTime) GobDecode(data []byte) error {
	return d.UnmarshalBinary(data)
}

func (d DateTime) GobEncode() ([]byte, error) {
	return d.MarshalBinary()
}

func (d DateTime) Hour() int {
	return time.Time(d).Hour()
}

func (d DateTime) ISOWeek() (year int, week int) {
	return time.Time(d).ISOWeek()
}

func (d DateTime) In(loc *time.Location) time.Time {
	return time.Time(d).In(loc)
}

func (d DateTime) IsDST() bool {
	return time.Time(d).IsDST()
}

func (d DateTime) IsZero() bool {
	return time.Time(d).IsZero()
}

func (d DateTime) Local() time.Time {
	return time.Time(d).Local()
}

func (d DateTime) Location() *time.Location {
	return time.Time(d).Location()
}

func (d DateTime) MarshalBinary() ([]byte, error) {
	return time.Time(d).MarshalBinary()
}

func (d DateTime) MarshalText() ([]byte, error) {
	return time.Time(d).MarshalText()
}

func (d DateTime) Minute() int {
	return time.Time(d).Minute()
}

func (d DateTime) Month() time.Month {
	return time.Time(d).Month()
}

func (d DateTime) Nanosecond() int {
	return time.Time(d).Nanosecond()
}

func (d DateTime) Round(dn time.Duration) time.Time {
	return time.Time(d).Round(dn)
}

func (d DateTime) Second() int {
	return time.Time(d).Second()
}

func (d DateTime) Sub(u time.Time) time.Duration {
	return time.Time(d).Sub(u)
}

func (d DateTime) Truncate(dn time.Duration) time.Time {
	return time.Time(d).Truncate(dn)
}

func (d DateTime) UTC() time.Time {
	return time.Time(d).UTC()
}

func (d DateTime) Unix() int64 {
	return time.Time(d).Unix()
}

func (d DateTime) UnixMicro() int64 {
	return time.Time(d).UnixMicro()
}

func (d DateTime) UnixMilli() int64 {
	return time.Time(d).UnixMilli()
}

func (d DateTime) UnixNano() int64 {
	return time.Time(d).UnixNano()
}

func (d *DateTime) UnmarshalBinary(data []byte) error {
	var t time.Time
	if err := (&t).UnmarshalBinary(data); err != nil {
		return err
	}
	*d = DateTime(t)
	return nil
}

func (d *DateTime) UnmarshalText(data []byte) error {
	var t time.Time
	if err := (&t).UnmarshalText(data); err != nil {
		return err
	}
	*d = DateTime(t)
	return nil
}

func (d DateTime) Weekday() time.Weekday {
	return time.Time(d).Weekday()
}

func (d DateTime) Year() int {
	return time.Time(d).Year()
}

func (d DateTime) YearDay() int {
	return time.Time(d).YearDay()
}

func (d DateTime) Zone() (name string, offset int) {
	return time.Time(d).Zone()
}

func (d DateTime) ZoneBounds() (start time.Time, end time.Time) {
	return time.Time(d).ZoneBounds()
}
