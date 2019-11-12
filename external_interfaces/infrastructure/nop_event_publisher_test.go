package infrastructure_test

import (
	"testing"

	"gitlab.com/dc0d/go-workshop/external_interfaces/infrastructure"
	"gitlab.com/dc0d/go-workshop/model"

	"github.com/stretchr/testify/require"
)

func Test_nop_event_publisher(t *testing.T) {
	t.Run("create nop event publisher", func(t *testing.T) {
		var _ model.EventPublisher = infrastructure.NewNopEventPublisher()
	})

	t.Run("call publish returns nil error", func(t *testing.T) {
		assert := require.New(t)

		publisher := infrastructure.NewNopEventPublisher()

		err := publisher.Publish()
		assert.NoError(err)

		err = publisher.Publish(model.EventRecord{})
		assert.NoError(err)
	})
}
