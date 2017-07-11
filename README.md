# arena-go

A Golang interface to the [are.na](https://www.are.na/) [API](https://dev.are.na/documentation). Not quite started, but will be a work in progress. Partly using this to learn some basic Go stuff.

## installation

    go get github.com/xanderseren/arena-go

## usage

All interaction starts with a `arena.Client`. Create one with your token:

```Go
import "github.com/xanderseren/arena-go"

client := arena.NewClient(token)
arena.GetChannel(channelID)
```    

## todo

- ...
