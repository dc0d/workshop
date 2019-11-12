package repositories_test

import (
	"testing"
	"time"

	"gitlab.com/dc0d/go-workshop/interface_adapters/repositories"
	"gitlab.com/dc0d/go-workshop/model"

	"github.com/stretchr/testify/require"
)

func Test_time_source(t *testing.T) {
	t.Run("create time factory", func(t *testing.T) {
		var _ model.TimeSource = repositories.NewTimeSource()
	})

	t.Run("test time factory", func(t *testing.T) {
		assert := require.New(t)

		src := repositories.NewTimeSource()

		utc := src.NowUTC()
		nearUTC := time.Now().UTC()

		diff := int64(nearUTC.Sub(utc))
		maxDelta := int64(time.Microsecond * 10)
		assert.LessOrEqual(diff, maxDelta)
	})
}
