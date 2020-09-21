package main

func main() {
	var value interface{}

	reallyLongCaluculation := func(
		done <-chan interface{},
		value interface{},
	) interface{} {
		intermediateResult := longCalculation(done, value)
		select {
		case <-done:
			return
		default:
		}

		return longCalculation(done, intermediateResult)
	}

	select {
	case <-done:
		return
	case value = <-valueStream:
	}

	result := reallyLongCaluculation(value)

	select {
	case <-done:
		return
	case resultStream <- result:
	}
}
