module Basics

go 1.18

//@NOTE
//When your code imports packages contained in other modules, you manage those dependencies through your code's own module. That module is defined by a go.mod file that tracks the modules that provide those packages.
//That go.mod file stays with your code, including in your source code repository.

//@NOTE Third party are added as requirement for codebase. We have just addded quote
require rsc.io/quote/v4 v4.0.1

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/sampler v1.3.0 // indirect
)
