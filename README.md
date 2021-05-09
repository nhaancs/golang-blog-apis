# Personal blog APIs

## Getting started

### Setup environment variables
- Rename `.env.sample` file to `.env`
- Add your configs to `.env` file

### Start database container
- The first time, run `make rundb`
- Otherwise, run `make startdb` instead

### Run migrations
- The first time, run `make buildmigrator` and `make migrateup`
- Then if you have any new migrations, just run `make migrateup`

### Start
- Run `make start`