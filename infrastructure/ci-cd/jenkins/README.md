# Jenkins CI/CD for Spa Management System

## Pipeline Folder Structure

```
infrastructure/ci-cd/
├── Jenkinsfile                        # (Legacy/monorepo pipeline, optional)
├── Jenkinsfile.customer-service       # Dedicated pipeline for customer-service (BE)
├── Jenkinsfile.appointment-service    # Dedicated pipeline for appointment-service (BE)
├── Jenkinsfile.staff-service          # Dedicated pipeline for staff-service (BE)
├── Jenkinsfile.branch-service         # Dedicated pipeline for branch-service (BE)
├── Jenkinsfile.marketing-service      # Dedicated pipeline for marketing-service (BE)
├── Jenkinsfile.finance-service        # Dedicated pipeline for finance-service (BE)
├── Jenkinsfile.analytics-service      # Dedicated pipeline for analytics-service (BE)
├── Jenkinsfile.shell                  # Dedicated pipeline for shell (FE)
├── Jenkinsfile.customer-app           # Dedicated pipeline for customer-app (FE)
├── Jenkinsfile.appointment-app        # Dedicated pipeline for appointment-app (FE)
├── Jenkinsfile.staff-app              # Dedicated pipeline for staff-app (FE)
├── Jenkinsfile.branch-app             # Dedicated pipeline for branch-app (FE)
├── Jenkinsfile.marketing-app          # Dedicated pipeline for marketing-app (FE)
├── Jenkinsfile.finance-app            # Dedicated pipeline for finance-app (FE)
├── Jenkinsfile.analytics-app          # Dedicated pipeline for analytics-app (FE)
├── jenkins-agent.Dockerfile           # Custom agent image (Docker, Go, Node.js, Trivy, kubectl)
```

## Jenkins Pipeline Usage

- Each Jenkinsfile is a dedicated pipeline for a single service or frontend app.
- Pipelines are parameterized with dropdowns for environment (`ENV`) and branch (`GIT_BRANCH`).
- Each backend service pipeline now includes an automated **Migrate DB** stage:
  - Runs `go run migrate.go` in `backend/pkg/database` to apply all SQL migrations before build/test.
  - Ensures the database schema is up-to-date for every deployment.
- Each pipeline performs:
  - Checkout
  - Migrate DB (for backend services)
  - Build & Test (Go for BE, npm for FE)
  - Docker build/tag/push
  - Trivy security scan
  - Kubernetes deploy (updates only the relevant deployment)
  - Email notification on success/failure
- All secrets are managed via Jenkins credentials.

## How to Create Jenkins Jobs Automatically

You can automate job creation using the Jenkins Job DSL plugin or by scripting with the Jenkins REST API.

### 1. Using Job DSL (recommended for code-as-config)
Create a file `jobs.groovy` in this folder with:

```groovy
folder('spa-management-system')

def services = [
  'customer-service', 'appointment-service', 'staff-service', 'branch-service',
  'marketing-service', 'finance-service', 'analytics-service',
  'shell', 'customer-app', 'appointment-app', 'staff-app', 'branch-app',
  'marketing-app', 'finance-app', 'analytics-app'
]

services.each { svc ->
  pipelineJob("spa-management-system/${svc}") {
    definition {
      cpsScm {
        scm {
          git {
            remote { url('https://your.git.repo.url') }
            branch('*/main')
          }
        }
        scriptPath("infrastructure/ci-cd/Jenkinsfile.${svc}")
      }
    }
    parameters {
      choiceParam('ENV', ['dev', 'sit', 'uat', 'nft', 'prd'], 'Target environment')
      choiceParam('GIT_BRANCH', ['develop', 'sit', 'uat', 'nft', 'main'], 'Branch to build')
    }
  }
}
```

- Install the [Job DSL plugin](https://plugins.jenkins.io/job-dsl/).
- Add a seed job in Jenkins pointing to this `jobs.groovy`.
- Run the seed job to create/update all pipelines.

### 2. Using Jenkins REST API (for scripting)
You can POST XML job configs to Jenkins using the REST API. See the [Jenkins API docs](https://www.jenkins.io/doc/book/using/remote-access-api/) for details.

---

**Tip:** Use the custom agent image (`jenkins-agent.Dockerfile`) for all jobs to ensure a consistent build environment.

If you need a sample `jobs.groovy` file or help with Jenkins configuration as code, just ask!
