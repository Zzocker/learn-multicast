# Multicast

Multicast is project for learning various Gossip and Epidemic protocol. It has server for simulating the protocols and client for interacting with it. The server simulate `gorutine` as `node` in a unstructured overlay, with each `gorutine` having open channel to subset of `gorutine`. Client makes `Set` call to a random `gorutine`, then this `gorutine` multicast the updated to value to the overlay. Client will make `Get` to any random `gorutine` in order to get the updated value. Supported multicast protocols are:

- Anti-Entropy
- Rumour Mongering
- Gossip Protocol

## Multicast server

```proto
service Multicast{
    rpc Set(Data) returns (Empty);
    rpc Get(Empty) returns (Data);
}
```

## How to use

- Start Multicast server: `go run cmd/server/main.go -c 16`
- Set `54` : `go run cmd/cli/main.go set -v 54`