# vanilla-api
[![CircleCI](https://circleci.com/gh/austin1237/vanilla-api.svg?style=svg)](https://circleci.com/gh/austin1237/vanilla-api)<br />
An api that only uses the standard lib

## Prerequistes
You must have the following installed/configured for this to work correctly<br />
1. [Docker](https://www.docker.com/community-edition)
2. [Docker-Compose](https://docs.docker.com/compose/)

## Development Environment
To start the api service which listens on localhost:8080 enter the following command
```bash
docker-compose up
```

### Tests
To run the test suite use the command
```bash
make test
```

