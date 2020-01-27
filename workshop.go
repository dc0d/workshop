package workshop

func BusStops(routes [][]int) int {
	state := newState(routes)

next_minute:
	for minute := 0; minute < 480; minute++ {
		state.gossipAt(minute)

		if !state.allMet() {
			continue next_minute
		}

		return minute + 1
	}

	return -1
}

func groupGossipsByStop(routes [][]int, minute int) (result []gossips) {
	zipped := zipAt(routes, minute)

	gossipGroup := make(map[int]gossips)
	for _routeNumber, busStop := range zipped {
		if gossipGroup[busStop] == nil {
			gossipGroup[busStop] = newGossips()
		}
		gossipGroup[busStop].add(routeNumber(_routeNumber))
	}

	for _, gossips := range gossipGroup {
		result = append(result, gossips)
	}

	return
}

func zipAt(routes [][]int, minute int) (busStops []int) {
	for _, route := range routes {
		busStops = append(busStops, route[minute%len(route)])
	}
	return
}

type state struct {
	state  map[routeNumber]gossips
	routes [][]int
}

func newState(routes [][]int) *state {
	state := &state{
		state:  make(map[routeNumber]gossips),
		routes: routes,
	}
	for i := range routes {
		routeNumber := routeNumber(i)
		item := newGossips()
		item.add(routeNumber)
		state.state[routeNumber] = item
	}
	return state
}

func (s *state) allMet() bool {
	for _, gossips := range s.state {
		if len(gossips) < len(s.state) {
			return false
		}
	}
	return true
}

func (s *state) gossipsFor(r routeNumber) gossips   { return s.state[r] }
func (s *state) mergeWith(r routeNumber, g gossips) { s.gossipsFor(r).merge(g) }

func (s *state) gossipAt(minute int) {
	gossipGroup := groupGossipsByStop(s.routes, minute)

	for _, gossips := range gossipGroup {
		for gossip := range gossips {
			for k := range gossips {
				merged := newGossips()
				merged.merge(s.gossipsFor(gossip))
				merged.merge(s.gossipsFor(k))

				s.mergeWith(gossip, merged)
				s.mergeWith(k, merged)
			}
		}
	}
}

type gossips map[routeNumber]struct{}

func newGossips() gossips {
	return make(gossips)
}

func (g gossips) add(r routeNumber) { g[r] = struct{}{} }
func (g gossips) merge(gossips gossips) {
	for k := range gossips {
		g.add(k)
	}
}

type (
	routeNumber int
)
