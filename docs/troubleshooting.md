# Troubleshooting Guide

This guide lists common issues and solutions for the Spa Management System in local, CI/CD, and production environments.

---

## 1. Local Development
- **Cannot connect to PostgreSQL/Redis:**
  - Ensure services are running (try `docker-compose up`)
  - Check `.env` for correct connection strings
- **Port conflicts:**
  - Change service/app port in `.env` or config files
- **Frontend build errors:**
  - Run `npm ci` to reinstall dependencies
  - Clear cache: `npm run clean` or delete `node_modules`

## 2. CI/CD Pipeline
- **Migration failures:**
  - Check migration logs in Jenkins/GitHub Actions
  - Ensure DB credentials/secrets are set in CI/CD
- **Docker build errors:**
  - Verify Dockerfile syntax and base images
- **Kubernetes deploy errors:**
  - Check image tags and deployment YAMLs
  - Use `kubectl describe pod` for logs

## 3. Production
- **Service not starting:**
  - Check logs (`kubectl logs <pod>`, cloud provider dashboard)
  - Verify secrets/configs are mounted
- **API Gateway issues:**
  - Check gateway logs and route configs
- **Performance issues:**
  - Check DB/Redis resource usage
  - Review HPA and resource limits in k8s manifests

---

For unresolved issues, contact the DevOps team or consult the #spa-support Slack channel.
