package channels

import "log"

func exampleDeclareReadOnlyChannel() {
	// read only channel.
	// means we can only read from this channel. we cannot send data to thee channel
	var dataStream <-chan interface{} 
	// read from
	dataStreamItem := <-dataStream
	log.Println(dataStreamItem)
	dataStream2 := make(<-chan interface{})
	dataStreamItem = <-dataStream2
	log.Println(dataStreamItem)
}

func sendDataOnly(ch chan <- interface{}) {
	ch <- 300
}

func readDataOnly(ch <- chan interface{}) {
	data := <- ch
	log.Println(data)
}

func exampleDeclareWriteOnlyChannel() {
	// write only channel.
	// means we can only write to this channel. we cannot read data from thee channel
	var dataStreamOut chan <- interface{} 
	// write to
	dataStreamOut <- 20
	dataStream2 := make(chan <- interface{})
	dataStream2 <- 400
	// readDataOnly(dataStream2) won't work
}

func universalChannel() {
	ch := make(chan interface{})
	sendDataOnly(ch)
	readDataOnly(ch)
}