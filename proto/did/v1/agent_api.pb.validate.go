// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: did/v1/agent_api.proto

package didV1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Ticket with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Ticket) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Ticket with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in TicketMultiError, or nil if none found.
func (m *Ticket) ValidateAll() error {
	return m.validate(true)
}

func (m *Ticket) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Timestamp

	// no validation rules for NonceValue

	// no validation rules for KeyId

	// no validation rules for Document

	// no validation rules for Proof

	// no validation rules for Signature

	if len(errors) > 0 {
		return TicketMultiError(errors)
	}

	return nil
}

// TicketMultiError is an error wrapping multiple validation errors returned by
// Ticket.ValidateAll() if the designated constraints aren't met.
type TicketMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m TicketMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m TicketMultiError) AllErrors() []error { return m }

// TicketValidationError is the validation error returned by Ticket.Validate if
// the designated constraints aren't met.
type TicketValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e TicketValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e TicketValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e TicketValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e TicketValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e TicketValidationError) ErrorName() string { return "TicketValidationError" }

// Error satisfies the builtin error interface
func (e TicketValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sTicket.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = TicketValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = TicketValidationError{}

// Validate checks the field values on PingResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PingResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PingResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PingResponseMultiError, or
// nil if none found.
func (m *PingResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PingResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ok

	if len(errors) > 0 {
		return PingResponseMultiError(errors)
	}

	return nil
}

// PingResponseMultiError is an error wrapping multiple validation errors
// returned by PingResponse.ValidateAll() if the designated constraints aren't met.
type PingResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PingResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PingResponseMultiError) AllErrors() []error { return m }

// PingResponseValidationError is the validation error returned by
// PingResponse.Validate if the designated constraints aren't met.
type PingResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PingResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PingResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PingResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PingResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PingResponseValidationError) ErrorName() string { return "PingResponseValidationError" }

// Error satisfies the builtin error interface
func (e PingResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPingResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PingResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PingResponseValidationError{}

// Validate checks the field values on ProcessRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ProcessRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProcessRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ProcessRequestMultiError,
// or nil if none found.
func (m *ProcessRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ProcessRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Task

	if all {
		switch v := interface{}(m.GetTicket()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ProcessRequestValidationError{
					field:  "Ticket",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ProcessRequestValidationError{
					field:  "Ticket",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTicket()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProcessRequestValidationError{
				field:  "Ticket",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ProcessRequestMultiError(errors)
	}

	return nil
}

// ProcessRequestMultiError is an error wrapping multiple validation errors
// returned by ProcessRequest.ValidateAll() if the designated constraints
// aren't met.
type ProcessRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProcessRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProcessRequestMultiError) AllErrors() []error { return m }

// ProcessRequestValidationError is the validation error returned by
// ProcessRequest.Validate if the designated constraints aren't met.
type ProcessRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProcessRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProcessRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProcessRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProcessRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProcessRequestValidationError) ErrorName() string { return "ProcessRequestValidationError" }

// Error satisfies the builtin error interface
func (e ProcessRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProcessRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProcessRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProcessRequestValidationError{}

// Validate checks the field values on ProcessResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ProcessResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ProcessResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ProcessResponseMultiError, or nil if none found.
func (m *ProcessResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ProcessResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Ok

	if len(errors) > 0 {
		return ProcessResponseMultiError(errors)
	}

	return nil
}

// ProcessResponseMultiError is an error wrapping multiple validation errors
// returned by ProcessResponse.ValidateAll() if the designated constraints
// aren't met.
type ProcessResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ProcessResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ProcessResponseMultiError) AllErrors() []error { return m }

// ProcessResponseValidationError is the validation error returned by
// ProcessResponse.Validate if the designated constraints aren't met.
type ProcessResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProcessResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProcessResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProcessResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProcessResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProcessResponseValidationError) ErrorName() string { return "ProcessResponseValidationError" }

// Error satisfies the builtin error interface
func (e ProcessResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProcessResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProcessResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProcessResponseValidationError{}

// Validate checks the field values on QueryRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QueryRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QueryRequestMultiError, or
// nil if none found.
func (m *QueryRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Method

	// no validation rules for Subject

	if len(errors) > 0 {
		return QueryRequestMultiError(errors)
	}

	return nil
}

// QueryRequestMultiError is an error wrapping multiple validation errors
// returned by QueryRequest.ValidateAll() if the designated constraints aren't met.
type QueryRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryRequestMultiError) AllErrors() []error { return m }

// QueryRequestValidationError is the validation error returned by
// QueryRequest.Validate if the designated constraints aren't met.
type QueryRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryRequestValidationError) ErrorName() string { return "QueryRequestValidationError" }

// Error satisfies the builtin error interface
func (e QueryRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryRequestValidationError{}

// Validate checks the field values on QueryResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *QueryResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on QueryResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in QueryResponseMultiError, or
// nil if none found.
func (m *QueryResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *QueryResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Document

	// no validation rules for Proof

	if len(errors) > 0 {
		return QueryResponseMultiError(errors)
	}

	return nil
}

// QueryResponseMultiError is an error wrapping multiple validation errors
// returned by QueryResponse.ValidateAll() if the designated constraints
// aren't met.
type QueryResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m QueryResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m QueryResponseMultiError) AllErrors() []error { return m }

// QueryResponseValidationError is the validation error returned by
// QueryResponse.Validate if the designated constraints aren't met.
type QueryResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e QueryResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e QueryResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e QueryResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e QueryResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e QueryResponseValidationError) ErrorName() string { return "QueryResponseValidationError" }

// Error satisfies the builtin error interface
func (e QueryResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sQueryResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = QueryResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = QueryResponseValidationError{}
