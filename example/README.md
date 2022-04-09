# application example

This is simple application server example

- Listen one port app endpoint
- Optinal another monitoring endpoint

## parameter configuration method

`main.Message` is configurable paremeter

Order

1. Commandline argument `-message` by `flag.Parse()` in `vermouth.Parse()`
2. Environment variable `MESSAGE=` by `vermouth.Parse`
3. Compile time embed value `-ldflags "-X main.Message="` in `Makefile`
4. Hardcoded value in source