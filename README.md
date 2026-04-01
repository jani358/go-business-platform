# Project Business

Go-based backend platform built with microservices, Docker, CI/CD, Kubernetes, and PostgreSQL.

## Services

- `api-gateway`: Single entry point for clients.
- `auth-service`: Authentication and authorization.
- `product-service`: Product CRUD (first MVP service).
- `order-service`: Order management.

## Folder Structure

```text
project-business/
  services/
    api-gateway/
    auth-service/
    product-service/
    order-service/
  pkg/
  deploy/
    docker/
    k8s/
      base/
      overlays/
  scripts/
  .github/
    workflows/
  docs/
```

## Getting Started

1. Install Go 1.22+ and Docker.
2. Start with `product-service` CRUD implementation.
3. Add Dockerfiles and compose setup.
4. Add CI/CD workflows.
5. Deploy to Kubernetes.

