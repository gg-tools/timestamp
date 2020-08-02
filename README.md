# Timestamp Tool

## Installation

Simply run:

`go install github.com/gg-tools/timestamp`

Make sure your `PATH` includes the `$GOPATH/bin` directory so your commands can be easily used:

`export PATH=$PATH:$GOPATH/bin`

## Usage

Get timestamp of a specific time:

```shell
$ timestamp
1596355497 (2020-08-02 16:04:57)
$ timestamp 1988-08-13
1596326400 (2020-08-02 00:00:00)
$ timestamp "1988-08-13 18:06:06"
1596391566 (2020-08-02 18:06:06)
```

Get timestamp of now plus 1 year 3 days 4 hours 5 minutes 6 seconds and minus 2 months:

```shell
$ timestamp +1y -2m 3d 4h 5i 6s
1622895151 (2021-06-05 20:12:31)
```

### Time

Get current time:

```shell
timestamp time
```

Get time from timestamps:

```shell
$ timestamp time 587433600
1988-08-13 09:00:00 (587433600)
$ timestamp t 587433600 1597276800
1988-08-13 09:00:00 (587433600)
2020-08-13 08:00:00 (1597276800)
```
