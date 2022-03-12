# Fare App Calculation System

## Build and Running on Local Maching
To build use `make build` or `go build -v -o fare-app cmd/*.go`.
This will generate `fare-app` binary that can be run using `./fare-app`.


To trigger both build and run, use `make run`

## Flags
This application support two optional flags
- `-input`: data input that will be processed by the application.
  If this is empty, application will start in console mode.

- `-config`: json config file that will be used as application configuration.
  If this is empty, it will use default config specified in `./internal/model/config.go

### Console Mode
In console mode, the app will wait for user input. Each data will be processed line by line,
until reaching empty line, which triggers fare calculation.

If there's no error found, the output will be printed in standard output.
#### Example
run the app `./fare-app`, and then input the following data:
```
00:00:00.000 0.0
00:01:00.123 480.9
00:02:00.125 1141.2

```
this will trigger the output of `440`

### File Mode
In file mode, the app will directly process inputted file. 
Each data will be processed line by line, until reaching end of file,
which triggers fare calculation.

If there's no error found, the output will be printed in standard output.
#### Example
run the app `./fare-app -input sample_data.txt` to run with provided sample data,
inside the sample data is:
```
00:00:00.000 0.0
00:01:00.123 480.9
00:02:00.125 1141.2
00:03:00.100 1850.8
```
this will trigger the output of `520`


## Running the app without compiling on local
If building the app on local is not possible, the `/bin` folder includes some
popular OS and architecture. Simply calls specific binary to use it.
For example, on Linux 32-bit calls `./bin/linux_x86/fare-app` to run the app.

All above flags and command is the same for all type of builds. For example
`./bin/linux_x86/fare-app -input sample_data.txt -config config.json`
will trigger fare calculation on `sample_data.txt` using configuration from `config.json`

Note that, the binary included in the root folder is built on Mac, so it won't run on other OS.
