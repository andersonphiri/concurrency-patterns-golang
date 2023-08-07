package forselect

// TChannel receives from a single stream then sends to two channels
func TChannelTwo[TIn any, TOut any](done <-chan interface{},
	inputStream <-chan TIn,
	processor func(TIn) TOut) (<-chan TOut, <-chan TOut) {
	out_one := make(chan TOut)
	out_two := make(chan TOut)
	go func() {
		defer close(out_one)
		defer close(out_two)
		for input := range OrDoneProcessor[TIn, TOut](done, inputStream, processor) {
			var out1, out2 = out_one, out_two
			for i := 0; i < 2; i++ {
				select {
				case <-done:
					return
				case out1 <- input:
					out1 = nil
				case out2 <- input:
					out2 = nil
				}
			}
		}
	}()
	return out_one, out_two

}
