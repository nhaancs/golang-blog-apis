# Deployment

## Client setup
- Update environment variables in `.env` file
- Provide permission to run `./deploy/deploy.sh`, `./deploy/migratedb.sh`, `./deploy/setupserver.sh` files
    ```bash
    make setpermissions
    ```
    
## Server setup (first time only)
- Setup ssh connection from client to server
- Run command
    ```bash
    make setupserver
    ```

## Deploy
- Do migrations (if any)
    ```bash
    make migratedb
    ```
- Deploy
    ```bash
    make deploy
    ```