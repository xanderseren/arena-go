# arena-go

A Golang interface to the [are.na](https://www.are.na/) [API](https://dev.are.na/documentation).

## installation

    go get github.com/xanderseren/arena-go

## usage

    import "github.com/xanderseren/arena-go"
    arena.channels.channel('faq')

All interaction starts with a `arena.Client`. Create one with your token:

```Go
client := arena.NewClient(token)
```    

## todo

- tests
- documentation
- authentication flow

## api parity progress

- [ ] Authentication
- [ ] Blocks
- [ ] Channels
- [ ] Search
- [ ] Users
