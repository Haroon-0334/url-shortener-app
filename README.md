# URL Shortener - app

Minimal Go URL shortener service with health and metrics endpoints.

## Run locally

1. Build: `docker build -t url-shortener:local .`
2. Run: `docker run -p 8080:8080 url-shortener:local`
3. Health: `curl localhost:8080/healthz`

## CI
Configured GitHub Actions to run tests, build and push to ECR, and update manifests repo.