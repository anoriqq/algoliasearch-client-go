// Code generated by go generate. DO NOT EDIT.

package opt

import "encoding/json"

// FiltersOption is a wrapper for an Filters option parameter. It holds
// the actual value of the option that can be accessed by calling Get.
type FiltersOption struct {
	value string
}

// Filters wraps the given value into a FiltersOption.
func Filters(v string) *FiltersOption {
	return &FiltersOption{v}
}

// Get retrieves the actual value of the option parameter.
func (o *FiltersOption) Get() string {
	if o == nil {
		return "attribute"
	}
	return o.value
}

// MarshalJSON implements the json.Marshaler interface for
// FiltersOption.
func (o FiltersOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON implements the json.Unmarshaler interface for
// FiltersOption.
func (o *FiltersOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = "attribute"
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

// Equal returns true if the given option is equal to the instance one. In case
// the given option is nil, we checked the instance one is set to the default
// value of the option.
func (o *FiltersOption) Equal(o2 *FiltersOption) bool {
	if o == nil {
		return o2 == nil || o2.value == "attribute"
	}
	if o2 == nil {
		return o == nil || o.value == "attribute"
	}
	return o.value == o2.value
}

// FiltersEqual returns true if the two options are equal.
// In case of one option being nil, the value of the other must be nil as well
// or be set to the default value of this option.
func FiltersEqual(o1, o2 *FiltersOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
