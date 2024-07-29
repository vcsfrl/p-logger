# p-logger
Code challenge - Logger

### Usage

```bash
# make sure you have docker and docker compose installed
make dc-install
make dc-test-cover
make dc-run-example

# for other options see commands below
```

## Commands

```bash
# NOTE: Makefile and commands are for Unix-like systems. I only tested on Linux.

# get help
make

# install dependencies
make install

# run tests
make test

# run tests with coverage
make test-cover

# run example code
make run-example

# install dependencies with docker compose
make dc-install

# run tests with docker compose
make dc-test

# run tests with coverage with docker compose
make dc-test-cover

# run example code with docker compose
make dc-run-example


#####################
# using go commands

# install dependencies
go mod vendor
go mod tidy

# run tests
go test ./... -race -cpu 24 -v

# run example code
go run example/main.go
 
```

## Boundary conditions

 - *Usability: Reasonable default configuration, supports different drivers and runs out
of the box.*
   - Just call `logger.NewLogger(nil)` to get a logger instance with the defaults. This will write log messages to stdout. The minimum log level is `INFO`. Any log with severity less than `INFO` will be ignored (`DEBUG`).
 - *Configurable: Customizable by a configuration file without touching the core code or
the code which uses this package.*
   - The logger can be configured from a `json` file. See `example/config.json` with all the available options.
 - *Extensible: Without touching code the package should be extensible by new drivers.*
   - The package is designed to be extensible. To create a new driver just implement the `logger.OutputWriter` and register a constructor method of type `logger.OutputWriterConstructor` with the `logger.RegisterOutputWriter` function. The registered constructor can be referenced in the configuration file. See `example/main.go` how an external driver is defined (`JsonOutputWriter`) and registered (`logger.RegisterOutputWriter("json_stdout", NewJsonOutputWriter)`).
 - *Programming languages have to be PHP or Go.*
   - GO
 - *The code has to be unit tested.*
   - I used TDD. The design evolved with the tests. 
 - *How you design and implement this is up to you. Because of that this project is
openly described on purpose.*
 - *The code has to be published on a public git repository.*