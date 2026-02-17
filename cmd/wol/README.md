# wol

## Abstract

I needed a simple tool to wake up my media center **and** wanted to replace the [`wakeonlan`](https://github.com/jpoliv/wakeonlan) package (written in Perl) available for Debian/Ubuntu (at least).

## Installation

```bash
$ go install github.com/da-rod/wakeonlan/cmd/wol@latest
```

## Usage

```bash
$ wol -h
Usage of wol:
  -mac string
        MAC Address to wake up
```
