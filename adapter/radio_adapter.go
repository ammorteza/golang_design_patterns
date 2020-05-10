package main

type RadioAdapter struct {
	Radio
}

func (this *RadioAdapter)Plug220Volt() string{
	return this.Plug12Volt()
}