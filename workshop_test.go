package workshop_test

import (
	"fmt"
	"testing"

	"github.com/dc0d/workshop"

	assert "github.com/stretchr/testify/require"
)

var _ *assert.Assertions

func Test_all_meet_at_same_bus_stop(t *testing.T) {
	type (
		expectation struct {
			routes   [][]int
			busStops int
		}
	)

	var (
		expectations = []expectation{
			{
				routes: [][]int{
					[]int{1},
					[]int{1},
				},
				busStops: 1,
			},
			{
				routes: [][]int{
					[]int{2, 1},
					[]int{3, 1},
				},
				busStops: 2,
			},
			{
				routes: [][]int{
					[]int{2, 1, 4},
					[]int{3, 1, 5},
					[]int{3, 1, 7},
				},
				busStops: 2,
			},
		}
	)

	for testCase, exp := range expectations {
		var (
			routes   = exp.routes
			busStops = exp.busStops
		)

		t.Run(fmt.Sprintf("test case number %v", testCase), func(t *testing.T) {
			var (
				assert = assert.New(t)
			)

			assert.Equal(busStops, workshop.BusStops(routes))
		})
	}
}

func Test_eventually_all_meet(t *testing.T) {
	type (
		expectation struct {
			routes   [][]int
			busStops int
		}
	)

	var (
		expectations = []expectation{
			{
				routes: [][]int{
					[]int{1, 2, 1},
					[]int{1, 3, 1},
					[]int{2, 3, 2},
				},
				busStops: 3,
			},
			{
				routes: [][]int{
					[]int{1, 2, 1},
					[]int{3, 4, 5, 1},
					[]int{2, 3, 2},
				},
				busStops: 12,
			},
			{
				routes: [][]int{
					[]int{3, 1, 2, 3},
					[]int{3, 2, 3, 1},
					[]int{4, 2, 3, 4, 5},
				},
				busStops: 5,
			},
		}
	)

	for testCase, exp := range expectations {
		var (
			routes   = exp.routes
			busStops = exp.busStops
		)

		t.Run(fmt.Sprintf("test case number %v", testCase), func(t *testing.T) {
			var (
				assert = assert.New(t)
			)

			assert.Equal(busStops, workshop.BusStops(routes))
		})
	}
}

func Test_never_meet(t *testing.T) {
	type (
		expectation struct {
			routes   [][]int
			busStops int
		}
	)

	var (
		expectations = []expectation{
			{
				routes: [][]int{
					[]int{2, 1, 2},
					[]int{5, 2, 8},
				},
				busStops: -1,
			},
		}
	)

	for testCase, exp := range expectations {
		var (
			routes   = exp.routes
			busStops = exp.busStops
		)

		t.Run(fmt.Sprintf("test case number %v", testCase), func(t *testing.T) {
			var (
				assert = assert.New(t)
			)

			assert.Equal(busStops, workshop.BusStops(routes))
		})
	}
}
