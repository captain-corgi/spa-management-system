# Kubernetes Manifests for Spa Management System

This directory contains the Kubernetes manifests for deploying the entire Spa Management System, following DDD, microservices, and best-practice cloud-native architecture.

## Structure

```
base/      # Namespace, ConfigMap, Secret, Ingress, HPAs, general configs
backend/   # Deployments and Services for all backend microservices
frontend/  # Deployments and Services for all frontend micro-frontends
database/  # PostgreSQL and Redis deployments, PVs, and PVCs
```

## Usage

### 1. Create Namespace and Base Resources
```
kubectl apply -f base/
```

### 2. Deploy Database Layer
```
kubectl apply -f database/
```

### 3. Deploy Backend Services
```
kubectl apply -f backend/
```

### 4. Deploy Frontend Applications
```
kubectl apply -f frontend/
```

### 5. (Optional) Apply All at Once
```
kubectl apply -f .
```

## Notes
- **ConfigMap/Secret:** Edit `base/configmap.yaml` and `base/secret.yaml` for environment-specific configuration.
- **Ingress:** Edit `base/ingress.yaml` for domain and path routing.
- **Autoscaling:** HPAs are in `base/hpa-backend.yaml` and `base/hpa-frontend.yaml`.
- **Persistent Storage:** PVs and PVCs for PostgreSQL and Redis are in `database/`.
- **Resource Limits:** All deployments have CPU/memory requests and limits set. Adjust as needed.
- **Health Checks:** All pods implement readiness and liveness probes.

## Best Practices
- Use `kubectl -n spa-system get all` to monitor resources.
- Use `kubectl describe` and `kubectl logs` for troubleshooting.
- Keep secrets out of version control in production.

---
For more details, see `.windsurf-guide.md` in this directory.
