package openhue

func (o *On) Switch() *On {
	v := !(*o.On)
	return &On{On: &v}
}
