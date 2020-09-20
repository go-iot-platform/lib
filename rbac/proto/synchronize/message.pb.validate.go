// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: synchronize/message.proto

package synchronize

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
var _message_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Response with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	// no validation rules for Message

	// no validation rules for ErrorMessage

	return nil
}

// ResponseValidationError is the validation error returned by
// Response.Validate if the designated constraints aren't met.
type ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResponseValidationError) ErrorName() string { return "ResponseValidationError" }

// Error satisfies the builtin error interface
func (e ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResponseValidationError{}

// Validate checks the field values on ZoneRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ZoneRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for GatewayMac

	for idx, item := range m.GetZones() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ZoneRequestValidationError{
					field:  fmt.Sprintf("Zones[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ZoneRequestValidationError is the validation error returned by
// ZoneRequest.Validate if the designated constraints aren't met.
type ZoneRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ZoneRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ZoneRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ZoneRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ZoneRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ZoneRequestValidationError) ErrorName() string { return "ZoneRequestValidationError" }

// Error satisfies the builtin error interface
func (e ZoneRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sZoneRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ZoneRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ZoneRequestValidationError{}

// Validate checks the field values on SceneRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SceneRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// SceneRequestValidationError is the validation error returned by
// SceneRequest.Validate if the designated constraints aren't met.
type SceneRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SceneRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SceneRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SceneRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SceneRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SceneRequestValidationError) ErrorName() string { return "SceneRequestValidationError" }

// Error satisfies the builtin error interface
func (e SceneRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSceneRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SceneRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SceneRequestValidationError{}

// Validate checks the field values on ScheduleRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ScheduleRequest) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ScheduleRequestValidationError is the validation error returned by
// ScheduleRequest.Validate if the designated constraints aren't met.
type ScheduleRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ScheduleRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ScheduleRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ScheduleRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ScheduleRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ScheduleRequestValidationError) ErrorName() string { return "ScheduleRequestValidationError" }

// Error satisfies the builtin error interface
func (e ScheduleRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sScheduleRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ScheduleRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ScheduleRequestValidationError{}

// Validate checks the field values on Zone with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *Zone) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Valid

	// no validation rules for ZoneId

	// no validation rules for Active

	// no validation rules for ZoneType

	// no validation rules for ZoneName

	// no validation rules for DevicesInzone

	// no validation rules for ParentIndex

	// no validation rules for IconName

	// no validation rules for ExDevId

	// no validation rules for ExZoneId

	// no validation rules for ChildZoneId

	// no validation rules for SecMode

	// no validation rules for EsTimeout

	// no validation rules for EsStatus

	// no validation rules for Operation

	// no validation rules for Dim

	// no validation rules for HomeSecMode

	// no validation rules for ZoneSecMode

	// no validation rules for Mac

	return nil
}

// ZoneValidationError is the validation error returned by Zone.Validate if the
// designated constraints aren't met.
type ZoneValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ZoneValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ZoneValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ZoneValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ZoneValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ZoneValidationError) ErrorName() string { return "ZoneValidationError" }

// Error satisfies the builtin error interface
func (e ZoneValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sZone.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ZoneValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ZoneValidationError{}

// Validate checks the field values on Request with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for DevId

	// no validation rules for Reachable

	// no validation rules for Active

	// no validation rules for Mac

	// no validation rules for DevType

	// no validation rules for SubDevType

	// no validation rules for ApplianceType

	// no validation rules for DevName

	// no validation rules for GatewayMac

	// no validation rules for UniqueName

	// no validation rules for Zoneid

	if v, ok := interface{}(m.GetStatus()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RequestValidationError{
				field:  "Status",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// RequestValidationError is the validation error returned by Request.Validate
// if the designated constraints aren't met.
type RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestValidationError) ErrorName() string { return "RequestValidationError" }

// Error satisfies the builtin error interface
func (e RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestValidationError{}

// Validate checks the field values on ChangeGatewayRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ChangeGatewayRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for MacAddress

	// no validation rules for Name

	// no validation rules for Serial

	return nil
}

// ChangeGatewayRequestValidationError is the validation error returned by
// ChangeGatewayRequest.Validate if the designated constraints aren't met.
type ChangeGatewayRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChangeGatewayRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChangeGatewayRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChangeGatewayRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChangeGatewayRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChangeGatewayRequestValidationError) ErrorName() string {
	return "ChangeGatewayRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ChangeGatewayRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChangeGatewayRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChangeGatewayRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChangeGatewayRequestValidationError{}

// Validate checks the field values on Status with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Status) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Red

	// no validation rules for Green

	// no validation rules for Blue

	// no validation rules for Warm

	// no validation rules for Natural

	// no validation rules for Cool

	// no validation rules for Dimm

	// no validation rules for SensorSecMode

	// no validation rules for OpMode

	// no validation rules for Sleep

	// no validation rules for Onoff

	// no validation rules for Light

	// no validation rules for Utlight

	// no validation rules for Ltlight

	// no validation rules for Temp

	// no validation rules for Uttemp

	// no validation rules for Lttemp

	// no validation rules for Motion

	// no validation rules for Enmotion

	// no validation rules for Motiontimeout

	// no validation rules for Slider

	// no validation rules for Ulslider

	// no validation rules for Llslider

	// no validation rules for Humidity

	// no validation rules for Uhumidity

	// no validation rules for Lhumidity

	// no validation rules for Sonoff

	// no validation rules for Ledstatusen

	// no validation rules for MConfig

	// no validation rules for MVal

	// no validation rules for DhtTemp

	// no validation rules for DhtHumidity

	// no validation rules for Plug_1

	// no validation rules for Current_1

	// no validation rules for Voltage_1

	// no validation rules for Power_1

	// no validation rules for Plug_2

	// no validation rules for Current_2

	// no validation rules for Voltage_2

	// no validation rules for Power_2

	// no validation rules for SlightOnoff

	// no validation rules for SlightSw_1

	// no validation rules for SlightSw_2

	// no validation rules for Slight_1

	// no validation rules for Slight_2

	// no validation rules for Slight_3

	// no validation rules for Slight_4

	// no validation rules for SlightSw_1Name

	// no validation rules for SlightSw_2Name

	// no validation rules for SlightName_1

	// no validation rules for SlightName_2

	// no validation rules for SlightName_3

	// no validation rules for SlightName_4

	// no validation rules for Secmode

	// no validation rules for Motionhwpresent

	// no validation rules for TpsMotion

	// no validation rules for TpsDoorlock

	// no validation rules for TpsCo

	// no validation rules for TpsSmoke

	// no validation rules for TpsAlarm1

	// no validation rules for TpsAlarm2

	// no validation rules for TpsTamper

	// no validation rules for TpsBattery

	// no validation rules for TpsSecm

	// no validation rules for TpiRev1

	// no validation rules for TpiSecm

	// no validation rules for Onoff_2

	// no validation rules for Onoff_3

	// no validation rules for GvswitchSw_1Name

	// no validation rules for GvswitchSw_2Name

	// no validation rules for GvswitchSw_3Name

	return nil
}

// StatusValidationError is the validation error returned by Status.Validate if
// the designated constraints aren't met.
type StatusValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StatusValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StatusValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StatusValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StatusValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StatusValidationError) ErrorName() string { return "StatusValidationError" }

// Error satisfies the builtin error interface
func (e StatusValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStatus.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StatusValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StatusValidationError{}
