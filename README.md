# The API Application

An API Application is an API implementation which builds it's contents by 
including a number of builders, an them activating them with various 
implementations.

The reasoning for the application approach is:

1. any api consumer will want to hardcode it's go library includes for the 
coach providers, an likely build it's builders before activating them
2. most api consumers will want some dynamic run-time configuration of the 
builders which cannot be hardcoded into the go code, and which will come 
from some central configuration.

So the typical app building approach is:


1. Create an instance of the application struct that you like
2. Create any builders from implementations that you like, including their
go library, and configuring their structs
3. Add each builder that should be avaialable to the application using the
AddBuilder() method
4. Decide what implementations you would like to activate (perhaps from some
configuration source) and for each implemtnatipo
5. Activate builders in a sequence that makes sense for you app.  Sometimes
one builder may require operations from a different builder, so this sequence
is important.
