Bounty Beacon
=============

A CLI tool to monitor and claim rewards from the Octopus Energy Octoplus partner program.

## Project Layout

- `main.go`: Thin entrypoint that delegates to the CLI package.
- `cli/`: Cobra/Viper command handling, structured logging, and user-facing output.
- `lib/client/`: Octopus client implementation (one method per file).
- `lib/operations/`: GraphQL operation requests and operation-specific types.
- `sanitize_har.sh`, `fetch_js.sh`: HAR analysis helpers.

## Current Progress

### 1. HAR Sanitization
We have developed a robust shell script `sanitize_har.sh` to prepare HAR files for analysis.
- **Filtering:** Keeps only relevant domains (e.g., `octopus.energy`).
- **Sanitization:** Redacts `Authorization` headers, `Cookie` values, and `POST` bodies to protect sensitive session data.
- **Robustness:** Includes type-checking to handle inconsistent JSON structures in HAR files.

### 2. API Reverse Engineering
Analyzed the Octopus GraphQL schema to identify key operations:
- **Endpoint:** `https://api.backend.octopus.energy/v1/graphql/`
- **Queries:** 
    - `getOctoplusRewards`: Retrieves existing voucher codes and history.
    - `getOctoplusOfferBySlug`: Checks specific offer availability (e.g., `caffe-nero`).
- **Authentication:** Requires a JWT Bearer token 

### 3. Go CLI Implementation
Implemented a Cobra/Viper-based Go CLI with the following commands:
- `rewards`: Lists all currently held vouchers and their expiry dates.
- `check`: Checks if the Caffè Nero offer is in stock or available to claim.
- `claim`: Attempts to claim the configured offer.
- `watch`: Polls availability and can auto-claim when stock appears.

## Usage

```bash
export OCTOPUS_REFRESH_TOKEN="your_refresh_token_here"
export OCTOPUS_CLIENT_ID="your_octopus_oauth_client_id"
go run . login
go run . check --offer=caffe-nero
go run . watch --offer=caffe-nero --interval=10s --auto-claim

# Optional env-based configuration via Viper
export CLAIM_POLL_INTERVAL=1s
export CLAIM_POLL_TIMEOUT=45s
export LOG_LEVEL=info
export LOG_FORMAT=text
go run . claim
```

### Useful CLI flags

- Global: `--log-level`, `--log-format`
- `check`: `--offer`
- `claim`: `--offer`, `--claim-poll-interval`, `--claim-timeout`
- `watch`: `--offer`, `--interval`, `--auto-claim`, `--claim-poll-interval`, `--claim-timeout`

## Kubernetes CronJob (Monday morning)

This repo includes:

- `Dockerfile` using a `FROM scratch` runtime image
- `k8s/cronjob.yaml` scheduled for Monday 03:00 (`Europe/London`)
- `k8s/secret.example.yaml` for the refresh token secret

The cron job runs an init container for `login`, then the main container runs `watch --auto-claim`.

Build and push:

```bash
docker build -t gregarendse/bountybeacon:latest .
docker push gregarendse/bountybeacon:latest
```

Create secret and apply CronJob:

```bash
kubectl apply -f k8s/secret.example.yaml
kubectl apply -f k8s/cronjob.yaml
```

If you want a different time, update `spec.schedule` in `k8s/cronjob.yaml`.

# References

- [HAR Schema Definition](http://www.softwareishard.com/blog/har-12-spec/)
- https://developer.octopus.energy/guides/graphql/api-basics/
- https://api.octopus.energy/v1/graphql
- https://developer.octopus.energy/
