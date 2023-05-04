# OSG-ARCH
This task manager is simple crm project for development team

## Run project
1. Clone project
2. Install docker, docker-compose and make
3. source .env file
4. Run command `make init_db`
5. Run command `make compose`

## To test connection rest api
```bash
curl -X GET http://localhost:8080/ping
```

if response is `pong` then connection is ok

