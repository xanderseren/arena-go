# arena-go

A Golang interface to the [are.na](https://www.are.na/) [API](https://dev.are.na/documentation). A work in progress. Partly using this to learn some basic Go stuff.

## installation

    go get github.com/xanderseren/arena-go

## usage

All interaction starts with a `arena.Client`:

```Go
import arena "github.com/xanderseren/arena-go"

client := arena.NewClient(token)
block, err := client.Blocks.Get(blockID, nil) // blockID can be ID or Slug
fmt.Println(block.Title)

```    

## todo

- [x] Authentication
- [x] Better interface solution for various Block and Channel functions

#### Blocks
- [x] GET /v2/blocks/:id
```Go
arena.Blocks.Get(ID, arguments)
```    
- [x] GET /v2/blocks/:id/channels
```Go
arena.Blocks.ListChannels(ID, arguments)
```
- [ ] POST /v2/channels/:slug/blocks
- [ ] PUT /v2/blocks/:id
- [ ] PUT /v2/channels/:channel_id/blocks/:id/selection
- [ ] DELETE /v2/channel/:channel_id/blocks/:id

#### Channels
- [ ] GET /v2/channels
- [x] GET /v2/channels/:slug
```Go
arena.Channels.Get(ID, arguments)
```
- [ ] GET /v2/channels/:slug/thumb
- [ ] GET /v2/channels/:id/connections
- [x] GET /v2/channels/:id/contents
```Go
arena.Channels.Contents(ID, arguments)
```
- [ ] POST /v2/channels
- [ ] PUT /v2/channels/:slug
- [ ] PUT /v2/channels/:slug/sort
- [ ] DELETE /v2/channels/:slug
- [ ] POST /v2/channels/:slug/blocks
- [ ] PUT /v2/channels/:channel_id/blocks/:id/selection
- [ ] GET /v2/channels/:id/collaborators
- [ ] POST /v2/channels/:id/collaborators
- [ ] DELETE /v2/channels/:id/collaborators

#### Users
- [ ] GET /v2/users/:id
- [ ] GET /v2/users/:id/channel
- [ ] GET /v2/users/:id/channels
- [ ] GET /v2/users/:id/following
- [ ] GET /v2/users/:id/followers

#### Search
- [ ] GET /v2/search?q=:q
- [ ] GET /v2/search/users?q=:q
- [x] GET /v2/search/channels?q=:q
```Go
arena.Channels.Search(arena.Arguments{"q": query})
```
- [ ] GET /v2/search/blocks?q=:q
```Go
arena.Blocks.Search(arena.Arguments{"q": query})
```
