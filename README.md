# Personal blog APIs

## Getting started

### Setup environment variables
- Rename `.env.sample` file to `.env`
- Add your configs to `.env` file

### Create Docker network
- Run `make createnetwork`

### Start database container
- The first time, run `make rundb`
- Otherwise, run `make startdb` instead

### Run migrations
- Run `make migrateup`

### Start
- Run `make start`