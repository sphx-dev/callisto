# Sphinx Callisto

### DB Setup
Before Starting we need two databases, one for Sphinx Callisto and one for Hasura.

DB_NAME: sphx_testnet
HASURA_GRAPHQL_METADATA_DATABASE_URL: hasura

Here I've create two databases sphx_testnet and hasura. You can create databases with any name.
But make sure to update the environment variables in the configuration files. 
`DB_NAME` would be used by Callisto and Hasura both.

```bash

### Config

Following is the configuration file for the Sphinx Callisto Local Node. Change the `DB_*` 
variables to your local database credentials. 

```yaml
chain:
  bech32_prefix: sphx
  modules:
    [
      actions,
      auth,
      bank,
      consensus,
      distribution,
      message_type,
      "daily refetch",
      mint,
      staking,
      gov,
      upgrade,
      marginacc,
      markets,
      orders,
      positions,
    ]
node:
  type: remote
  config:
    rpc:
      client_name: juno
      # address: https://rpc1.sphx.dev
      address: http://localhost:26657
      max_connections: 20
    grpc:
      # address: https://grpc.sphx.dev
      address: http://localhost:9090
      insecure: true
    api:
      # address: https://rest1.sphx.dev
      address: http://localhost:1317
parsing:
  workers: 1
  start_height: 1
  average_block_time: 500ms
  listen_new_blocks: true
  parse_old_blocks: true
  parse_genesis: true
database:
  url: postgresql://{DB_USER}:{DB_PASS}@{DB_HOST}:5432/{DB_NAME}?sslmode=disable&search_path=public
  max_open_connections: 1
  max_idle_connections: 1
  partition_size: 100000
  partition_batch: 1000
  ssl_mode_enable: "false"
  ssl_root_cert: ""
  ssl_cert: ""
  ssl_key: ""
logging:
  level: debug
  format: text
actions:
  host: 127.0.0.1
  port: 3000
```

### Running Hasura
After successfully running Callisto, you can run Hasura to interact with the database.
For local development, `./hasura/docker-compose.yml` is provided. Update the environment variables
in the file.
* PG_DATABASE_URL would be same as the `database.url` in the Sphinx Callisto configuration file if you are using the same database and same host.
Currently docker containers will use host's network to connect to the database(so postgres should be running on the host machine).
If don't have postgres running on the host machine then uncomment postgres service in the `docker-compose.yml` file and update the environment variables to point to the postgres service. Also comment `network_mode: host` from the containers service.

* After Hasura is running, you can access the console at `http://localhost:8080/console` and run queries.

Now Hasura needs metadata to be imported. You can import the metadata from `./hasura/metadata` directory by running `hasura metadata apply --endpoint http://127.0.0.1:8080`.
