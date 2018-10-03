# goplugin_example
This is a sample to illustrate on create and use goplugins


## Requires

Go version 1.8 (or) above

## Overview

Simple greeter application that uses two language fonts: english and telugu.
Takes language of choice (telugu and english are supported as cmd line arguments in the application)

### Creating go plugin

1. Create a plugin package for a particular language. eg:- tel to represent a lanuage telugu
2. Create a plugin with name 'greeter.go' under tel package.
3. Use capitilization rule to export Function,Variables to be exported as part of the plugin.
4. The function Greet() and variable Greeter are exported
5. When this plugin's Greet() is invoked, it will print a message in telugu.

### Notes about the plugin

1. The plugin package is a regular go package
2. The package declared should be main
3. The exported functions (or) variables can be of any type

### Compiling the plugin

1. `go build -buildmode=plugin -o tel/tel.so tel/greeter.go`
This will be compiled using the normal go tool chain and will generate tel.so in tel package
2. `go build -buildmode=plugin -o eng/eng.so eng/greeter.go`
This will compiled using the normal go tool chain and will generate eng.so in eng package

### Using the go plugin in application

1. import the package plugin
2. Define/select type for imported package. 
3. Determine the type of .so file to be loaded based on cmd line arguments i.e -lang
4. Open the plugin package. plug.Open will load the .so file for symbols
5. Lookup for symbol in the package. plug.Lookup("Greeter")
6. Assert for loaded symbol of desired type. 
7. Invoke the interface function.



### Running the application

Run `go run greeter.go -lang telugu`
This will print the greeting in language telugu

Run `go run greeter.go -lang english`
This will print the greeting in language english




