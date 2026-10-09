# Vesper

Vesper is an experimental peer-to-peer networking project written in Go. The goal is to explore node discovery, communication, and file sharing across a decentralized network.

> **Work in progress:** Vesper currently focuses on discovering peers on a local network.

## Current features

- Automatic peer discovery using UDP broadcasts
- Unique node IDs and configurable node names
- Peer tracking with last-seen timestamps
- Removal of inactive peers after 10 seconds
- Docker Compose lab with three nodes

## Getting started

### Requirements

- Go 1.27.1 (for running locally)
- Docker with Compose (for the multi-node lab)

Clone the repository:

```bash
git clone https://github.com/fillipehmeireles/vesper.git
cd vesper
```

Run a local node:

```bash
go run ./cmd/main.go --name node01
```

Or start the three-node Docker lab:

```bash
make lab-up
```

Stop the lab with `make lab-down`, or remove its containers, volumes, and locally built images with `make lab-clean`.

## How it works

Nodes periodically broadcast discovery messages over UDP (port `6969`). When a node receives a discovery message from another node, it updates its peer table. Peers that have not been seen for more than 10 seconds are removed from the active table.

## Roadmap

- [x] UDP-based peer discovery
- [x] Peer tracking and inactivity detection
- [x] Multi-node Docker lab
- [ ] TCP-based peer communication
- [ ] Peer-to-peer file transfer
- [ ] Discovery and connectivity across the Internet

## Status

Vesper is a personal learning project and is under active development. Its networking protocol and APIs may change.
