module example.com/hello

go 1.22.5

require (
	example.com/greetings v0.0.0-00010101000000-000000000000
	golang.org/x/example/hello v0.0.0-20240716161537-39e772fc2670
)

replace example.com/greetings => ../helloworld/greetings
