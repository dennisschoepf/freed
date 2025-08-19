# freed

A unique feed aggregator and bookmarking service that tries to free you from too much content. Currently only a CLI program, but I plan to add a web server and frontend in the future. In active development.

## Installation

The easiest way to install the program is by using `go install`:

```sh
go install ...
```

## Usage

### Add a RSS feed or youtube channel

```sh
freed add
```

### List all feeds

```sh
freed ls
```

### Show a feed entry for the day

```sh
freed today
```

### Sync the added feeds

```sh
freed sync
```
