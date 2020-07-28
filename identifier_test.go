package core

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPartyID_IsZero(t *testing.T) {
	t.Run("zero", func(t *testing.T) {
		assert.True(t, PartyID{}.IsZero())
	})
	t.Run("non zero", func(t *testing.T) {
		partyID, _ := NewPartyID("foo", "bar")
		assert.False(t, partyID.IsZero())
	})
}

func TestPartyID_MarshalJSON(t *testing.T) {
	partyID, _ := NewPartyID("123", "foo")
	actual, _ := partyID.MarshalJSON()
	assert.Equal(t, quote(partyID.String()), string(actual))
}

func TestPartyID_Unmarshal(t *testing.T) {
	expected := "urn:oid:1.2.3:foo"
	actual := PartyID{}
	_ = json.Unmarshal([]byte(quote(expected)), &actual)
	assert.Equal(t, expected, actual.String())
}

func TestPartyID_Marshal(t *testing.T) {
	partyID, _ := NewPartyID("123", "foo")
	actual, _ := json.Marshal(partyID)
	assert.Equal(t, quote(partyID.String()), string(actual))
}

func TestPartyID_UnmarshalJSON(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		id := PartyID{}
		expected := "urn:oid:1.2.3:value"
		err := id.UnmarshalJSON([]byte(quote(expected)))
		assert.NoError(t, err)
		assert.Equal(t, expected, id.String())
	})
	t.Run("ok - empty string", func(t *testing.T) {
		id := PartyID{}
		err := id.UnmarshalJSON([]byte(quote("")))
		assert.NoError(t, err)
		assert.True(t, id.IsZero())
	})
	t.Run("error - no bytes", func(t *testing.T) {
		id := PartyID{}
		err := id.UnmarshalJSON([]byte(""))
		assert.Error(t, err)
		assert.True(t, id.IsZero())
	})
	t.Run("error - invalid format", func(t *testing.T) {
		id := PartyID{}
		err := id.UnmarshalJSON([]byte(quote("foobar")))
		assert.EqualError(t, err, "invalid PartyID: foobar")
	})
}

func TestPartyID_NewPartyID(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		partyID, err := NewPartyID("1.2.3", "foo")
		assert.NoError(t, err)
		assert.Equal(t, "urn:oid:1.2.3:foo", partyID.String())
	})
	t.Run("error - qualifier empty", func(t *testing.T) {
		partyID, err := NewPartyID("    ", "foo")
		assert.EqualError(t, err, "PartyID qualifier is empty")
		assert.True(t, partyID.IsZero())
	})
	t.Run("error - value empty", func(t *testing.T) {
		partyID, err := NewPartyID("1.2.3", "  ")
		assert.EqualError(t, err, "PartyID value is empty")
		assert.True(t, partyID.IsZero())
	})
}

func TestPartyID_Value(t *testing.T) {
	partyID, _ := NewPartyID("1.2.3", "foo")
	assert.Equal(t, "foo", partyID.Value())
}

func TestPartyID_String(t *testing.T) {
	t.Run("non zero", func(t *testing.T) {
		partyID, _ := NewPartyID("1.2.3", "foo")
		assert.Equal(t, "urn:oid:1.2.3:foo", partyID.String())
	})
	t.Run("zero", func(t *testing.T) {
		assert.Equal(t, "", PartyID{}.String())
	})
}

func quote(input string) string {
	return `"` + input + `"`
}