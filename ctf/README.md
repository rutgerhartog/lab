# Archivist

The archivist listens to a NATS topic and pushes the results to a postgres database.

## Environment variables

| Variable | Description | Default |
| -------- | ----------- | ------- |
| `NATS_URL` | URL of the NATS server | `nats://172.17.0.1:4222` |
| `NATS_SUBSCRIBER_TOPIC` | Subscriber topic for NATS | orders |
| `POSTGRES_HOST` | The hostname or IP address of the PostgreSQL database | `127.0.0.1`
| `POSTGRES_PORT` | The (TCP) port the database listens on | `5432`
| `POSTGRES_USER` | The username for the database | `postgres`
| `POSTGRES_DB` | The database name | `postgres`
| `POSTGRES_PASSWORD` | The password belonging to the postgres user | |

## Example

The archivist requires data on the NATS topic to be of a specific format:

```golang
type Order struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Address string    `json:"address"`
	Amount  float32   `json:"amount"`
}
```

For example, the following is valid:

```json
{
    "id": "8fb49e20-7063-4342-a6c0-85b3b09086de",
    "name": "Bo ter Ham",
    "address": "Bakkerstraat 1",
    "amount": 1337.25
}
```

but this is not:

```json
{
    "id": "123",
    "name": "Bo ter Ham",
    "address": "Bakkerstraat 1",
    "amount": 1337.25
}
```

as the UUID is not correct, nor is this:

```json
{
    "id": "8fb49e20-7063-4342-a6c0-85b3b09086de",
    "name": "Bo ter Ham",
    "amount": 1337.25
}
```
as the `address` field is missing.


## Testing 

For testing, you'll want to use the `nats` CLI tool: https://docs.nats.io/using-nats/nats-tools/nats_cli. With this installed, you can send a payload to the NATS server like this:

```sh
$ cat payload.json | nats publish 
```
where `payload.json` is a file containing a correct `order` object. Moreover, this requires your environment variables like `$NATS_URL` to be set correctly.
