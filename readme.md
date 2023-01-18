# epoch

A Go library for working with time intervals and time parsing.

## Installation

To install the library, use `go get`:

```bash
go get github.com/aahainc/epoch
```

## Usage

### Parsing Intervals

The library provides a function to parse time intervals from strings in the format of `value+unit`, where `value` is a
float number and `unit` is one of `s`, `m`, `h`, `d`, `w`, `mo`, `y`.
For example, `5m` stands for 5 minutes.

```golang
interval, err := epoch.ParseInterval("5m")
if err != nil {
// handle error
}
fmt.Println(interval) // 5m

```

### Parsing Time

The library also provides a function to parse time from strings in the format of `time.RFC3339` or unix timestamp
format.

```golang
t, err := epoch.ParseTime("2006-01-02T15:04:05Z")
if err != nil {
// handle error
}
fmt.Println(t)
```

### Safe Duration

A `Duration()` method is provided for `Interval` struct, but it panics on non-finite interval (years, months).
`IsSafeDuration()` method is provided to check if interval is finite, before calling `Duration()`

```golang
interval, _ := epoch.ParseInterval("5m")
if interval.IsSafeDuration() {
fmt.Println(interval.Duration())
}
```

## Examples

See `examples/` for more complete examples.

## Contributing

We welcome contributions to the `epoch` library. To contribute, fork the repository and submit a pull request.

## License

The `epoch` library is released under the [MIT License](https://opensource.org/licenses/MIT).