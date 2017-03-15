package property

// ValueChan returns a channel value that will update a property whenever it is changed
func ValueChan(p Property) chan interface{} {
	val := make(chan interface{})
	go func() {
		p.Set(<-listener) // @TODO should this be looping?  I think so.
	}()
	return val
}
