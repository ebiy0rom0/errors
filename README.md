# errors [![Unit-Test](https://github.com/ebiy0rom0/errors/actions/workflows/unittest.yml/badge.svg)](https://github.com/ebiy0rom0/errors/actions/workflows/unittest.yml) [![codecov](https://codecov.io/gh/ebiy0rom0/errors/branch/develop/graph/badge.svg?token=VBTPX64FKX)](https://codecov.io/gh/ebiy0rom0/errors)

Package errors is simple error handler for Go language.  
It's reference pkg/errors and privides similar functionality.

## Installation

This package can be installed with the `go get` command:
```
go get github.com/ebiy0rom0/errros
```

## How to use

This package obtains stack traces for all functions.  
If want to issue your own errors(e.g. validation checks) use `errors.New` or `errors.Errorf`.
```go
  // Cases requiring 0 or more for param
  if param <= 0 {
    return errors.New("param is less than 0.")
    // or new format error is:
    // return errors.Errorf("param is %d but required 0 or more.", param)
  }
```

If using an existing error or an error returned by a function, use `errors.Wrap` or `errors.Wrapf`. Both can wrap variables that satisfy the error interface.
```go
  // Cases failed file open and returned os.ErrNotExist
  name := "example.log"
  fp, err := os.Open(name)
  if err != nil {
    return errors.Wrap(err, "no log output distination.")
    // or wrap format error is:
    // return errors.Wrapf(err, "no log file %s", name)
  }
```

## Formatted printing of errors
This package implement fmt.Formatter and supports `%s` and `%v` by the fmt package.
```
%s  print the error. If for wrapped error, print at chained recursively.
%v  simultaneously printed stack trace obtained at oldest.
```

## License

MIT License