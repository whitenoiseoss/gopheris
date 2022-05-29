package internal

import (
	"testing"

	"github.com/whitenoiseoss/gopheris/discord/intents"
)

func TestNewBitmask(t *testing.T) {
	bm := NewBitmask()
	wantedFlags := 0

	if bm.Flags != wantedFlags {
		t.Fatalf("Got %b, wanted %b", bm.Flags, wantedFlags)
	}
}

func TestAddAndCheckFlag(t *testing.T) {
	bm := NewBitmask()
	bm.AddFlag(intents.DIRECT_MESSAGES)

	want := true
	got := bm.CheckFlag(intents.DIRECT_MESSAGES)

	if got != want {
		t.Fatalf("Got %t, wanted %t", got, want)
	}
}

func TestAddAndCheckFlagFalse(t *testing.T) {
	bm := NewBitmask()
	bm.AddFlag(intents.DIRECT_MESSAGE_REACTIONS)

	want := false
	got := bm.CheckFlag(intents.GUILD_MESSAGE_REACTIONS)

	if got != want {
		t.Fatalf("Got %t, wanted %t", got, want)
	}
}

func TestAddFlagsCheckFlag(t *testing.T) {
	bm := NewBitmask()
	bm.AddFlags(intents.DIRECT_MESSAGE_REACTIONS, intents.GUILD_MESSAGE_REACTIONS)

	want := true
	got := bm.CheckFlag(intents.GUILD_MESSAGE_REACTIONS)

	if got != want {
		t.Fatalf("AddFlags did not add expected flag. Got %t, wanted %t.", got, want)
	}
}

func TestAddAndCheckFlags(t *testing.T) {
	bm := NewBitmask()
	bm.AddFlags(intents.DIRECT_MESSAGE_REACTIONS, intents.GUILD_MEMBERS)

	want := true
	got := bm.CheckFlags(intents.DIRECT_MESSAGE_REACTIONS, intents.GUILD_MEMBERS)

	if got != want {
		t.Fatalf("Got %t, wanted %t", got, want)
	}
}

func TestAddAndCheckFlagsFalse(t *testing.T) {
	bm := NewBitmask()
	bm.AddFlags(intents.DIRECT_MESSAGE_REACTIONS, intents.GUILD_MESSAGE_REACTIONS)

	want := false
	got := bm.CheckFlags(intents.DIRECT_MESSAGE_REACTIONS, intents.GUILD_MEMBERS)

	if got != want {
		t.Fatalf("Got %t, wanted %t", got, want)
	}
}

func TestRemoveFlag(t *testing.T) {
	bm := NewBitmask()
	bm.AddFlags(intents.DIRECT_MESSAGE_REACTIONS, intents.GUILD_MESSAGE_REACTIONS)
	bm.RemoveFlag(intents.DIRECT_MESSAGE_REACTIONS)

	want := false
	got := bm.CheckFlag(intents.DIRECT_MESSAGE_REACTIONS)

	if got != want {
		t.Fatalf("Got %t, wanted %t", got, want)
	}
}

func TestRemoveFlags(t *testing.T) {
	bm := NewBitmask()
	bm.AddFlags(intents.DIRECT_MESSAGE_REACTIONS, intents.GUILD_MESSAGE_REACTIONS)
	bm.RemoveFlags(intents.DIRECT_MESSAGE_REACTIONS, intents.GUILD_MESSAGE_REACTIONS)

	want := 0
	got := bm.Flags

	if got != want {
		t.Fatalf("Flags were not cleared. Got %b, wanted %b", got, want)
	}
}

func TestToggleFlag(t *testing.T) {
	bm := NewBitmask()
	bm.AddFlag(intents.GUILD_MESSAGE_REACTIONS)
	bm.ToggleFlag(intents.GUILD_MESSAGE_REACTIONS)

	want := false
	got := bm.CheckFlag(intents.GUILD_MESSAGE_REACTIONS)

	if got != want {
		t.Fatalf("Flag was not toggled. Got %t, wanted %t", got, want)
	}
}

func TestToggleFlagOn(t *testing.T) {
	bm := NewBitmask()
	bm.AddFlags(intents.DIRECT_MESSAGE_REACTIONS, intents.GUILD_MESSAGE_REACTIONS)
	bm.ToggleFlag(intents.GUILD_MESSAGE_REACTIONS)
	bm.ToggleFlag(intents.GUILD_MESSAGE_REACTIONS)

	want := true
	got := bm.CheckFlag(intents.GUILD_MESSAGE_REACTIONS)

	if got != want {
		t.Fatalf("Double toggle did not return original state. Got %t, wanted %t", got, want)
	}
}

func TestToggleFlags(t *testing.T) {
	bm := NewBitmask()
	bm.AddFlags(intents.DIRECT_MESSAGE_REACTIONS, intents.GUILD_MESSAGE_REACTIONS, intents.GUILD_MEMBERS)
	bm.ToggleFlags(intents.DIRECT_MESSAGE_REACTIONS, intents.GUILD_MESSAGE_REACTIONS)

	want := intents.GUILD_MEMBERS
	got := bm.Flags

	if got != want {
		t.Fatalf("Flags were not toggled. Got %b, wanted %b", got, want)
	}
}
