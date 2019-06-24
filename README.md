# Docker Hands on

http://reireias-slides.firebaseapp.com/docker

## dockerfile
Create simple golang web server image.

```bash
cd dockerfile

docker build -t go-hello .

docker run --rm -p 8888:8080 go-hello

curl http://localhost:8888
```

## docker-compose
Creat golang web server + DB.

```bash
docker compose up -d

curl http://localhost:8888
curl http://localhost:8888/tasks
```
