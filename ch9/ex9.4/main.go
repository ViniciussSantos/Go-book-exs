package main

func pipeline(number int) (in, out chan int) {
	out = make(chan int)
	first := out
	for i := 0; i < number; i++ {
		in = out
		out = make(chan int)
		go func(in chan int, out chan int) {
			for v := range in {
				out <- v
			}
			close(out)
		}(in, out)
	}
	return first, out

}
