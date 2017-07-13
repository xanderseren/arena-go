# arena-go

A Golang interface to the [are.na](https://www.are.na/) [API](https://dev.are.na/documentation). A work in progress. Partly using this to learn some basic Go stuff.

## installation

    go get github.com/xanderseren/arena-go

## usage

All interaction starts with a `arena.Client`:

```Go
import arena "github.com/xanderseren/arena-go"

client := arena.NewClient(token)
arena.GetChannel(channelID) //channel ID or slug
arena.GetBlock(blockID) // block ID or slug
```    

## todo

- [x] Authentication
- [ ] Better methods solution for various Block and Channel calls

#### Blocks
- [x] GET /v2/blocks/:id
- [ ] GET /v2/blocks/:id/channels
- [ ] POST /v2/channels/:slug/blocks
- [ ] PUT /v2/blocks/:id
- [ ] PUT /v2/channels/:channel_id/blocks/:id/selection
- [ ] DELETE /v2/channel/:channel_id/blocks/:id

#### Channels
- [ ] GET /v2/channels
- [x] GET /v2/channels/:slug
- [ ] GET /v2/channels/:slug/thumb
- [ ] GET /v2/channels/:id/connections
- [ ] GET /v2/channels/:id/connections
- [ ] GET /v2/channels/:id/contents
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
- [ ] GET /v2/search/channels?q=:q
- [ ] GET /v2/search/blocks?q=:q
