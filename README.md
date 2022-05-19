# Whatsapp Bot project


### Beforehand
Docker instalation is required to run this project, if you dont have, you should follow https://www.docker.com/get-started/



## Quickstart
You can start copying the .env.example to .env to set basic setup of db:
```bash
  cp .env.example .env
```

This project is using docker compose to mount the developer enviroment. There are 3 containers to run:
- DB(postgres),
- Client(React/typescript)
- Server(Golang)

Run all with this command:
```bash
  docker-compose up -d --build
```
Server: http://localhost:5000/api/v1/messages
Client: http://localhost:3000
KeyCloak: http://localhost:8080


For server development, the container is using hotreload, so just save the file you are changing and be happy. The react container does not have this feature yet.