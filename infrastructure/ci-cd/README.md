# CI/CD for Spa Management System

This directory contains all configuration for automated testing, building, security scanning, and deployment of services and apps.

## Folder Structure

```
ci-cd/
├── github/
│   └── workflows/        # All GitHub Actions YAMLs
├── jenkins/
│   ├── agent/            # Jenkins agent Dockerfile
│   ├── pipelines/        # All dedicated Jenkinsfiles (per service/app)
│   └── README.md         # Jenkins-specific docs & automation
├── .windsurf-guide.md    # Pipeline guide, env var reference
├── README.md             # This file (overview)
```

## Where to Find Pipelines

- **GitHub Actions:**
  - All workflows in `github/workflows/`
  - Covers backend, frontend, meta checks, and multi-environment deployment
- **Jenkins:**
  - All dedicated pipelines in `jenkins/pipelines/` (one Jenkinsfile per service/app)
  - Custom Jenkins agent Dockerfile in `jenkins/agent/`
  - Jenkins automation and docs in `jenkins/README.md`

## Usage

- **GitHub Actions:**
  - Runs on push/PR to environment branches
  - Matrix builds, Docker, Trivy, Kubernetes deploy, notifications
- **Jenkins:**
  - One pipeline per service/app for granular control
  - Parameterized for environment and branch
  - Use Job DSL or REST API for automated job creation

## Multi-Environment Branching & Deployment

Branch-to-environment mapping:
- `develop` → DEV (`spa-dev` namespace)
- `sit` → SIT (`spa-sit` namespace)
- `uat` → UAT (`spa-uat` namespace)
- `nft` → NFT (`spa-nft` namespace)
- `main` → PRD (`spa-system` namespace)

On push or PR to these branches, the pipeline will:
- Build and tag Docker images with `<env>-<commit-sha>`
- Deploy to the corresponding Kubernetes namespace
- Notify on deployment status

See `.windsurf-guide.md` for required secrets and environment variables.

---
For Jenkins details, see `jenkins/README.md`.
For GitHub Actions details, see `github/workflows/`.
For pipeline guide and environment variable requirements, see `.windsurf-guide.md`.
