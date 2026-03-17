# VPN Manager

A Go-based control plane for provisioning WireGuard VPN users automatically.

## Features

- Generate WireGuard key pairs
- Allocate VPN IP addresses
- Register peers automatically
- Generate client configuration files

## Architecture

Control Plane (Go)
        │
        ▼
WireGuard Server

## Tech Stack

- Go
- WireGuard
- Linux (Debian)

## Future Work

- Persistent database for users/devices
- REST API for provisioning
- Metrics & monitoring
- Multi-node VPN infrastructure