package repositories_test

import (
	"testing"

	"github.com/dc0d/workshop/interface_adapters/repositories"
	model "github.com/dc0d/workshop/domain_model"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func Test_create_event_store(t *testing.T) {
	var _ model.EventStorage = repositories.NewEventStore(nil, nil)
}

func Test_call_event_storage(t *testing.T) {
	t.Run("call load method", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		assert := require.New(t)

		storage := NewMockEventStore(ctrl)
		storage.
			EXPECT().
			Load(gomock.Any()).
			Return(nil, nil)

		store := repositories.NewEventStore(storage, nil)
		_, err := store.Load("STREAM_ID")
		assert.NoError(err)
	})

	t.Run("call append method", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		assert := require.New(t)

		storage := NewMockEventStore(ctrl)
		storage.
			EXPECT().
			Append().
			Return(nil)

		publisher := NewMockEventPublisher(ctrl)
		publisher.
			EXPECT().
			Publish(gomock.AssignableToTypeOf([]model.EventRecord{})).
			Return(nil)

		store := repositories.NewEventStore(storage, publisher)
		err := store.Append()
		assert.NoError(err)
	})
}
