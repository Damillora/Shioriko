# Shioriko

A booru-like software written in Go and Svelte.

## Installation

The easiest way to get started is to use Docker:
```bash
docker pull damillora/shioriko
```

## Requirements

* PostgreSQL database

## Configuration

Shioriko is configured using environment variables:

* `POSTGRES_DATABASE`: DSN string of Postgres Database, see [Gorm documentation](https://gorm.io/docs/connecting_to_the_database.html)
* `AUTH_SECRET`: Secret used to sign JWTs
* `DATA_DIR`: Data directory to store images
* `BASE_URL`: Accesible URL of the instance
* `DISABLE_REGISTRATION`: Optional, disable registration on the instance

## Contributing
Shioriko is still in an early stage, but contributions are welcome!

## License
[MIT](https://choosealicense.com/licenses/mit/)