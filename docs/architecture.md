# Architecture Overview

## Core Principles

- Microservices with bounded contexts.
- Database-per-service ownership model.
- SOLID principles in each service codebase.
- API Gateway as external entry point.

## Initial Runtime

- Communication: REST (sync) for initial phase.
- Datastores: PostgreSQL.
- Local development: Docker Compose.
- Production target: Kubernetes.
