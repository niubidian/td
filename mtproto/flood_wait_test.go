package mtproto

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"golang.org/x/xerrors"

	"github.com/gotd/td/tgerr"
)

func TestAsFloodWait(t *testing.T) {
	err := func() error {
		return xerrors.Errorf("failed to perform operation: %w",
			tgerr.New(400, "FLOOD_WAIT_10"),
		)
	}()

	d, ok := AsFloodWait(err)
	assert.True(t, ok)
	assert.Equal(t, time.Second*10, d)
}
