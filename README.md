# VPN Manager

A production-style control plane for provisioning WireGuard VPN users with automated IP allocation, device management, and observability.

---

## 🚀 Features

- Generate WireGuard key pairs
- Allocate VPN IP addresses dynamically
- Register and manage VPN peers
- Generate client configuration files
- REST API for device lifecycle (create, list, delete)
- PostgreSQL-backed persistence
- Prometheus metrics instrumentation
- Grafana dashboards for observability
- Fully containerized backend (Docker)
- Infrastructure provisioning via Terraform

---

## 🧠 Architecture
       ┌───────────────┐
       │   Grafana     │
       └──────▲────────┘
              │
       ┌──────┴────────┐
       │  Prometheus   │
       └──────▲────────┘
              │
       ┌──────┴────────┐
       │   Go API      │
       │ (vpn-manager) │
       └──────▲────────┘
              │
    ┌─────────┴─────────┐
    │   PostgreSQL      │
    └───────────────────┘

       (WireGuard runs on host)


---

## ⚙️ Tech Stack

- **Go** (backend + API)
- **PostgreSQL** (persistent storage)
- **WireGuard** (VPN data plane)
- **Docker** (containerization)
- **Terraform** (infrastructure as code)
- **Prometheus** (metrics collection)
- **Grafana** (visualization)
- **Linux (Debian)**

---

## 🔍 Observability

The system exposes Prometheus metrics:

- `vpn_requests_total` → request volume
- `vpn_request_duration_seconds` → latency
- `vpn_errors_total` → error tracking

Prometheus scrapes metrics and Grafana visualizes them.

---

## 🐳 Containerization

- Go backend is built using a multi-stage Dockerfile
- Runs independently from host environment
- Handles missing system dependencies (e.g., WireGuard) gracefully

---

## 🏗️ Infrastructure (Terraform)

Terraform manages:

- Prometheus container
- Grafana container

Enables reproducible infrastructure setup.

---

## ⚠️ Design Note

WireGuard operations are treated as **optional inside containerized environments**:

- API remains functional even if `wg` is unavailable
- Prevents hard dependency on host-level networking tools

---

## 📌 Future Improvements

- Run WireGuard inside container with proper privileges (`NET_ADMIN`)
- Containerize PostgreSQL (remove host dependency)
- Add authentication & user management
- Implement rate limiting and access control
- Multi-region / multi-node VPN orchestration
- Improve Grafana dashboards

---

## 📄 Summary

This project demonstrates:

- Backend system design
- Infrastructure automation
- Observability integration
- Containerization challenges & solutions
- Real-world debugging across networking and system boundaries