// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: log-serial/log-serial.proto

package logserial

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _log_serial_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on DataRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DataRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for CustomerNumber

	// no validation rules for ThingName

	// no validation rules for ThingDisplayName

	// no validation rules for ThingSerial

	// no validation rules for PropertyName

	// no validation rules for DataType

	// no validation rules for Value

	// no validation rules for PropertyId

	// no validation rules for ThingId

	for idx, item := range m.GetAlerts() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return DataRequestValidationError{
					field:  fmt.Sprintf("Alerts[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for TemplateId

	// no validation rules for TemplateName

	return nil
}

// DataRequestValidationError is the validation error returned by
// DataRequest.Validate if the designated constraints aren't met.
type DataRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataRequestValidationError) ErrorName() string { return "DataRequestValidationError" }

// Error satisfies the builtin error interface
func (e DataRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDataRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataRequestValidationError{}

// Validate checks the field values on AlertData with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *AlertData) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	// no validation rules for Value

	// no validation rules for SettingValue

	// no validation rules for AlertModule

	// no validation rules for Status

	// no validation rules for PriorityId

	for idx, item := range m.GetTriggers() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return AlertDataValidationError{
					field:  fmt.Sprintf("Triggers[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// AlertDataValidationError is the validation error returned by
// AlertData.Validate if the designated constraints aren't met.
type AlertDataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AlertDataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AlertDataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AlertDataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AlertDataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AlertDataValidationError) ErrorName() string { return "AlertDataValidationError" }

// Error satisfies the builtin error interface
func (e AlertDataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAlertData.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AlertDataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AlertDataValidationError{}

// Validate checks the field values on Trigger with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Trigger) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Micro

	// no validation rules for Endpoint

	// no validation rules for Body

	// no validation rules for Params

	// no validation rules for Namespace

	return nil
}

// TriggerValidationError is the validation error returned by Trigger.Validate
// if the designated constraints aren't met.
type TriggerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TriggerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TriggerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TriggerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TriggerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TriggerValidationError) ErrorName() string { return "TriggerValidationError" }

// Error satisfies the builtin error interface
func (e TriggerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTrigger.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TriggerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TriggerValidationError{}

// Validate checks the field values on DataResponse with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *DataResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for StatusCode

	return nil
}

// DataResponseValidationError is the validation error returned by
// DataResponse.Validate if the designated constraints aren't met.
type DataResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DataResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DataResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DataResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DataResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DataResponseValidationError) ErrorName() string { return "DataResponseValidationError" }

// Error satisfies the builtin error interface
func (e DataResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDataResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DataResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DataResponseValidationError{}
