# phoebe

![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/Damillora/phoebe/dev.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/Damillora/phoebe)](https://goreportcard.com/report/github.com/Damillora/phoebe)
![GitHub License](https://img.shields.io/github/license/Damillora/phoebe)

<img src="./phoebe-logo.svg " alt="project logo" width="300" height="200">

a booru-style image gallery and organizer.

Built with the Go language and Svelte framework, phoebe is designed for personal image gathering.

## Features
* Upload and organize images
* Basic tagging system
* Search and autocomplete based on tags
* Similarity search and duplicate detection using perceptual hash

## Installation

The easiest way to get started is to use Docker:
```bash
docker pull damillora/phoebe
docker run -e POSTGRES_DATABASE=<PostgreSQL DSN> -e AUTH_SECRET=<secret> -e DATA_DIR=/data -e BASE_URL=http://localhost:8080 -p 8080:8080 -v "./data:/data" damillora/phoebe
```

## Requirements

* PostgreSQL database

## Configuration

phoebe is configured using environment variables:

* `POSTGRES_DATABASE`: DSN string of Postgres Database, see [Gorm documentation](https://gorm.io/docs/connecting_to_the_database.html)
* `AUTH_SECRET`: Secret used to sign JWTs
* `DATA_DIR`: Data directory to store images
* `BASE_URL`: Accesible URL of the instance
* `DISABLE_REGISTRATION`: Optional, disable registration on the instance

## Contributing
phoebe is still in an early stage, but contributions are welcome!

## License
phoebe is licensed under the [MIT license](https://choosealicense.com/licenses/mit/).