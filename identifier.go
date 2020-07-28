package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var partyIDPattern = regexp.MustCompile("urn:oid:([0-9\\.]+):(.*)")

// PartyID is a data type uniquely identifying a party in the Nuts Network.
// It's represented as a URN-encoded OID: https://www.ietf.org/rfc/rfc8141.txt
// For example: urn:oid:1.2.3.4:foo
type PartyID struct {
	oid   string
	value string
}

// MarshalJSON marshals the PartyID to JSON.
func (i PartyID) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// MarshalJSON unmarshals the PartyID from JSON.
func (i *PartyID) UnmarshalJSON(bytes []byte) error {
	var str string
	if err := json.Unmarshal(bytes, &str); err != nil {
		return err
	}
	if str == "" {
		*i = PartyID{}
		return nil
	}
	if partyID, err := ParsePartyID(str); err != nil {
		return err
	} else {
		*i = partyID
		return nil
	}
}

// ParsePartyID tries to parse the given input as URN-encoded OID (for example: urn:oid:1.2.3.4:foo)
func ParsePartyID(input string) (PartyID, error) {
	parts := partyIDPattern.FindStringSubmatch(input)
	if len(parts) != 3 {
		return PartyID{}, fmt.Errorf("invalid PartyID: %s", input)
	} else {
		return NewPartyID(parts[1], parts[2])
	}
}

// NewPartyID creates a new PartyID
func NewPartyID(qualifier string, value string) (PartyID, error) {
	if strings.TrimSpace(qualifier) == "" {
		return PartyID{}, errors.New("PartyID qualifier is empty")
	}
	if strings.TrimSpace(value) == "" {
		return PartyID{}, errors.New("PartyID value is empty")
	}
	return PartyID{oid: qualifier, value: value}, nil
}

// IsZero tests whether this identifier is empty a.k.a. 'zero'.
func (i PartyID) IsZero() bool {
	return i.value == "" && i.oid == ""
}

// String returns the PartyID as fully-qualified URN-encoded OID.
func (i PartyID) String() string {
	if i.value == "" {
		return ""
	}
	return fmt.Sprintf("urn:oid:%s:%s", i.oid, i.value)
}

// Value returns the value part of the PartyID.
func (i PartyID) Value() string {
	return i.value
}

// OID returns the OID of the PartyID
func (i PartyID) OID() string {
	return i.oid
}
