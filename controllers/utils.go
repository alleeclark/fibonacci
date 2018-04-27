package controllers

// take the first num of items off of an incoming stream and then exit
func take(done <-chan interface{}, valueStream <-chan interface{}, num int) <-chan interface{} {
	takeStream := make(chan interface{})
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}

// repeatedFunc repeatedly will calls a function passed in as an argument until you tell it to stop
func repeatFunc(done <-chan interface{}, n int, fn func(n int) int) <-chan interface{} {
	valueStream := make(chan interface{})
	go func() {
		defer close(valueStream)
		for {
			select {
			case <-done:
				return
			default:
				for i := 0; i <= n; i++ {
					valueStream <- fn(i)
				}
			}
		}
	}()
	return valueStream
}
