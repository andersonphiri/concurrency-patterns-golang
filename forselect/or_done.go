package forselect

func OrDoneProcessor[TIn any, TOut any](done <-chan interface{},
	inputStream <-chan TIn, processor func(TIn) TOut) <-chan TOut {
	stream := make(chan TOut)
	go func() {
		defer close(stream)
		for {
			select {
			case <-done:
				return
			case pull, ok := <-inputStream:
				if !ok {
					return
				}
				select {
				case stream <- processor(pull):
				case <-done:
					return
				}
			}
		}

	}()
	return stream
}
