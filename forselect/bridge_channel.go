package forselect

// Bridge bridges a channel of channels into a single channel
func Bridge[T any](done <-chan interface{}, 
	channels <-chan (<-chan T)) <-chan T {
		
		valStream := make(chan T)
		go func ()  {
			defer close(valStream)
			processor := func(in T)T { return in }
			for {
				var stream <-chan T 
				select {
				case maybeStream, ok := <- channels:
					if !ok {
						return
					}
					stream = maybeStream
				case <-done:
					return
				}
				for value := range OrDoneProcessor[T,T](done, stream, processor) {
					select {
					case valStream <- value:
					case <- done:
					}
				}
			}
		}()
		return valStream


}