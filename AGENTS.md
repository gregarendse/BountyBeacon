AGENTS
=========

This project, **BountyBeacon**, is a Go-based CLI tool designed to monitor and claim rewards from the Octopus Energy Octoplus partner program. It includes utilities for reverse engineering the Octopus API from HAR files and provides a programmatic way to interact with the backend.

## Project Overview

- **Purpose**: Automate checking and claiming of rewards (like Caffè Nero or Greggs drinks) from Octopus Energy.
- **Technologies**: Go (Golang) for the main CLI, Bash and `jq` for HAR file processing and sanitization.
- **Architecture**:
    - `main.go`: Thin entrypoint that calls the `cli` package.
    - `cli/`: Cobra/Viper command layer and user interaction/output.
    - `lib/client/`: Octopus client implementation (method-oriented, split by file).
    - `lib/operations/`: GraphQL operations and operation-specific request/response types.
    - `sanitize_har.sh`: Utility to filter and redact sensitive information from HAR files.

## Building and Running

### Prerequisites
- Go 1.25+
- `jq` (for HAR processing)

### Key Commands

- **Run the CLI**:
  ```bash
  go run . <command>
  ```
  Available commands: `login`, `rewards`, `check`, `claim`, `watch`.

- **Authentication**:
  1. Set your API key (preferred — uses `obtainKrakenToken` GraphQL mutation, no captcha required):
     ```bash
     export OCTOPUS_API_KEY="your_api_key"
     ```
     Get your API key from: https://octopus.energy/dashboard/new/accounts/personal-details/api-access
  2. Or set your refresh token (legacy — tokens expire after ~2 weeks):
     ```bash
     export OCTOPUS_REFRESH_TOKEN="your_refresh_token"
     export OCTOPUS_CLIENT_ID="your_octopus_oauth_client_id"
     ```
  3. Perform initial login to save configuration:
     ```bash
     go run . login
     ```
  This saves tokens and account info to `~/.bountybeacon.json`.

- **Check Offer Status**:
  ```bash
  go run . check --offer=caffe-nero
  ```

- **List Rewards**:
  ```bash
  go run . rewards
  ```

- **Claim Offer**:
  ```bash
  go run . claim --offer=caffe-nero --claim-poll-interval=1s --claim-timeout=45s
  ```

- **Watch Offer**:
  ```bash
  go run . watch --offer=caffe-nero --interval=30s --auto-claim
  ```

- **Sanitize a HAR file**:
  ```bash
  ./sanitize_har.sh <input.har> <output.har> <domain_to_keep> [--login-only] [--api-only]
  ```

## Development Conventions

- **API Endpoints**:
    - Main GraphQL: `https://api.backend.octopus.energy/v1/graphql/`
    - Token: `https://auth.octopus.energy/token/`
- **Headers**: API requests include `Authorization` (raw token string), `Origin`, and `Referer` for parity with observed browser calls.
- **Configuration**: Session details are stored in `~/.bountybeacon.json`.
- **CLI Config (Viper)**:
    - Env vars: `OCTOPUS_API_KEY`, `OCTOPUS_REFRESH_TOKEN`, `OCTOPUS_CLIENT_ID`, `CLAIM_POLL_INTERVAL`, `CLAIM_POLL_TIMEOUT`, `LOG_LEVEL`, `LOG_FORMAT`
    - Flags: `--offer`, `--interval`, `--auto-claim`, `--claim-poll-interval`, `--claim-timeout`, `--log-level`, `--log-format`
- **Security**: Never commit `~/.bountybeacon.json` or unsanitized HAR files. Use `sanitize_har.sh` before sharing or processing HAR data.
