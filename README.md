# weather-service

## Table of Contents

- [Weather Forecast Service](#weather-forecast-service)
  - [Table of Contents](#table-of-contents)
  - [Overview](#overview)
  - [Getting started](#getting-started)
    - [First Time Setup](#first-time-setup)
    - [Prerequisites](#prerequisites)
  - [Shortcuts](#shortcuts)
  - [Development](#development)
    - [Project Structure Organization](#project-structure-organization)
      - [Project Root](#project-root)
      - [Project Source (`./src`)](#project-source-src)
    - [Building the Code](#building-the-code)
    - [Unit Tests](#unit-tests)
    - [Component Tests](#component-tests)
    - [Integration Tests](#integration-tests)
    - [Local Testing](#local-testing)
    - [Enhancements](#enhancements)
      - [New APIs](#new-apis)
  - [Deployment](#deployment)
    - [Deployment to DEV](#deployment-to-dev)
    - [Deployment to UAT](#deployment-to-uat)
    - [Deployment to PRODUCTION](#deployment-to-production)
  - [Owners](#owners)

## Overview

To support a simple weather forecast server

| Entity | Description |
| ----------- | ----------- |
| Forecast | The weather forecast for the given input |

## Getting Started

### First Time Setup

1. Clone repo
```
git clone git@github.com:djohnson-engineer/weather-service.git
```

2. (Assuming go is installed on your machine, if not head to prerequisites below) Run the following command: 
```
PORT=8950 HOST=0.0.0.0 go run src/cmd/main.go
```

3. Head to the following sample endpoint: [Sample Endpoint for Boston](http://localhost:8950/api/weather/v1/forecast?latitude=42.361145&longitude=-71.057083)

### Prerequisites

1.  [Golang](https://go.dev/dl/). This service is written in go. Make sure $GOPATH, $GOPRIVATE environment variable are set.

```
brew install go
```
or head to [download link](https://go.dev/dl/):

2.  [Visual Studio Code](https://code.visualstudio.com/) for development and debugging the service

```
brew install --cask visual-studio-code
```

3. [mockery](https://github.com/vektra/mockery) to generate mocks from interfaces to make speed up unit testing development

```
go install github.com/vektra/mockery/v2/...@latest
```

4. [gin](https://github.com/gin-gonic/gin) for a high-performance simple web framework to take care of middleware and validation

```
go get -u github.com/gin-gonic/gin
```

5. [wire](https://github.com/google/wire) to handle dependency injenction

```
go install github.com/google/wire/cmd/wire@latest
```

## Shortcuts

Noting down in a single location for ease of reading a few shortcuts taken and changes needed:
1. No Auth Layer
1. No Common Validation Layer
1. No Swagger UI/API docs or similar
1. No component and obviously no integration testing
1. Minimal, quickly thrown together Unit tests for POC
1. No dedicated logging library
1. No monitoring integration
1. No caching
1. Minimal handling on HTTP Response codes
1. Minimal Server Implementation
1. No handling for API Response headers
1. And probably many more that will come up in conversation depending on the reader's interests

## Development

### Project Structure Organization

#### Project Root

| Folder | Description |
| ----------- | ----------- |
| .github | For github workflows/owners |
| .vscode | Local Dev Scripts |
| ./reports | Output for lint/tests |
| ./src | Source code base folder |

#### Project Source (`./src`)

| Folder | Description |
| ----------- | ----------- |
| /cmd | app, config and `main.go` |
| /container | Dependency Injection |
| /controllers | Controllers (handlers) for REST APIs |
| /datasource | APIs, DBs, Pipelines, as datasources for this service |
| /domain | Internal data models |
| /interfaces | Interfaces for usage in controllers |
| /logger | Service logging utility |
| /managers | Connector between controllers and datasources |
| /mocks | Mockery-generated mocks for unit testing |
| /models | Service Output Models |
| /testutils | Utility functions shared across packages for testing |
| /translation | Connector between managers and service output |
| /utils | Utility functions shared across packages |

### Building the Code

```
go build ./...
```

### Unit Tests

```
go test ./...
```

or for code coverage:

```
make cover
```

### Component Tests

TBD

### Integration Tests

TBD

### Local Testing

- Running the Service from Visual Studio Code:

In order to debug the service, one can run the service directly from the VS Code `Run and Debug` -> `Launch` and add breakpoints

- API Health End-point

TODO: Swagger UI or similar

```
http://localhost:8950/api/weather/v1/health
```

### Enhancements

#### New APIs

TBD

## Deployment

### Deployment to Dev

TBD

### Deployment to UAT

TBD

### Deployment to Production

TBD

## Owners

Team: TeamName

Email: TeamName@jackhenry.com

Slack: [#TeamName](Link)

Teams: [TeamName](Link)

Jira/Azure Devops/Etc.: [TeamName Board](Link)