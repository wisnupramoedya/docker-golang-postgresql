## Structure of DB
* title: string
* genre: array of string
* description: string
* director: string
* year: number

Material: [Docker Compose with Golang](https://www.mitrais.com/news-updates/how-to-dockerize-a-restful-api-with-golang-and-postgres/)

To run dev setting to code.
```sh
# to start
docker-compose -f docker-compose.dev.yml up -d
# to stop
docker-compose -f docker-compose.dev.yml down
```

To run prod setting to code.
```sh
# to start
docker-compose up -d
# to stop

docker-compose down
```

CompileDaemon could not hot reload.