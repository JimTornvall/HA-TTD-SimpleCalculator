# SimpleCalculator

- **Author:** Jim Törnvall
- **Year:** 2024
- **School:** Åland University of Applied Sciences
- **Course:** Test Driven Development
- **Status:** [![Go Build and Test](https://github.com/JimTornvall/HA-TTD-SimpleCalculator/actions/workflows/go.yml/badge.svg)](https://github.com/JimTornvall/HA-TTD-SimpleCalculator/actions/workflows/go.yml)

## Info

This is a simple TDD project created for school.

## Packages
- Testify
  - https://github.com/stretchr/testify
  - Assert: Used for assertions in the tests
  - Suite: Used for creating test suites, used to mimic @BeforeEach and @AfterEach in JUnit
  - Mock: Isnt used as i like Mockio better
- Mockio
  - https://github.com/ovechkin-dm/mockio
  - Mockito for golang, used for mocking interfaces in the tests

## Setup

```shell
go mod tidy
```

## Run

```shell
go run
```

## Test 

```shell
go test -v ./...
```

## TODO

- One more test for the separator spec + newlines among the numbers
- Refactor all the duplicated code in add.go
- Build the actual calculator main.go
