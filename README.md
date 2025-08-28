# freed

A unique feed aggregator that tries to free you from consuming too much content.

## Installation

The easiest way to install the program is by using `go install`:

```sh
go install ...
```

## Usage

### Add a RSS feed or youtube channel

Currently only RSS and Atom feeds and are supported.

```sh
freed add "https://feed.url"
```

### List all feeds

Lists all currently stored feed with their metadata.

```sh
freed list # or `freed ls`
```

### Show a feed entry for the day

Shows/opens an article from one of your feeds and marks it as read.

```sh
freed today
```

### Sync the added feeds

Manually syncs all feeds in the database. An automatic sync is performed on running any command if the last sync was more than 24 hours ago.

```sh
freed sync
```
