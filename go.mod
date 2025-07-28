module trung_go

go 1.24

require (
	rsc.io/quote/v4 v4.0.1
	trung_go/greetings v0.0.0-00010101000000-000000000000
)

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	golang.org/x/tour v0.1.0
	rsc.io/sampler v1.3.0 // indirect
)

replace trung_go/greetings => ./greetings
