package _134

func canCompleteCircuit(gas []int, cost []int) int {
	length := len(gas)
	rest, run, start := 0, 0, 0

	for i := 0; i < length; i++ {
		run += (gas[i] - cost[i])
		rest += (gas[i] - cost[i])
		if run < 0 {
			start = i + 1
			run = 0
		}
	}

	if rest < 0 {
		return -1
	}

	return start
}
