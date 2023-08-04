package forselect


func IterateAndSendChannel() {
	var queue chan <- string 
	done := make(chan struct{})
	for _, str := range []string{"my", "name", "is", "anderson"} {
		select {
		case <- done:
			return
		case queue <- str :
		}
	}

}