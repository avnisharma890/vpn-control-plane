# VPN Manager — System Design

## Overview

This project is a control plane for managing WireGuard VPN users.

It automates:
- key generation
- IP allocation
- peer registration
- peer removal
- live configuration reload

---

## Architecture

User (future API)
        │
        ▼
Go Control Plane
        │
        ├── PostgreSQL (source of truth)
        │
        └── WireGuard (data plane)

---

## Core Concepts

### Control Plane vs Data Plane

- Control Plane → Go application
- Data Plane → WireGuard interface (wg0)

The control plane decides *what should exist*.
The data plane enforces it.

---

## Current Features

### 1. Device Creation

Flow:

1. Generate key pair
2. Allocate IP
3. Store device in PostgreSQL
4. Append peer to wg0.conf
5. Reload WireGuard using syncconf
6. Return client config

---

### 2. Device Deletion

Flow:

1. Identify device by public key
2. Remove from PostgreSQL
3. Remove peer block from wg0.conf
4. Reload WireGuard

---

## Database Schema

### devices

| column      | type      | description |
|------------|----------|-------------|
| id         | SERIAL   | primary key |
| public_key | TEXT     | unique device identifier |
| vpn_ip     | TEXT     | assigned VPN IP |
| created_at | TIMESTAMP| creation time |

---

## Important Learnings

### 1. State Management

- wg0.conf was initially the source of truth
- now PostgreSQL is the source of truth

### 2. RBAC

- PostgreSQL enforces role-based access
- app user must be explicitly granted permissions

### 3. Infrastructure Sync

- DB state must match WireGuard state
- inconsistency = broken system

---

## Known Limitations

- Config parsing is naive (string-based)
- No API layer yet
- No user/device mapping (only devices)
- No concurrency handling for IP allocation

---

## Next Steps

- Add GET /devices
- Build REST API
- Add user model
- Improve config parsing
- Add metrics

## API Layer

### GET /devices

Returns all registered VPN devices from PostgreSQL.

Flow:
Client → API → DB → JSON response

### DELETE /devices/:id

Deletes a VPN device.

Flow:
Request → API → DB lookup → WireGuard removal → reload → response