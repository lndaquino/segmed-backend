# Segmed challenge
### This is a basic image management API to Segmed challenge.
&nbsp;&nbsp;
# Getting started
### To run this API you must follow the steps bellow:

## 1. Dependencies
### Assure you have the following dependencies installed:

`1 - download Docker Desktop if you are in Windows or Mac ==>` [install docker desktop](https://www.docker.com/products/docker-desktop)

`2 - download Docker and docker-compose if you are running a Linux machine ==>` [docker - steps 1 and 2](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-docker-on-ubuntu-18-04) [docker-compose - only step 1](https://www.digitalocean.com/community/tutorials/how-to-install-docker-compose-on-ubuntu-18-04)

`3 - install Golang according your operating system (api tested on version go1.15.7) ==>` [install Go](https://golang.org/doc/install)

`4 - download "migration tool"` [migrate](https://github.com/golang-migrate/migrate)
```sh
$   go get -tags 'postgres' -u github.com/golang-migrate/migrate/cmd/migrate
important: you must run this command in a directory without go.mod or in a directory without a parent directory with go.mod. The installation directory must be in your path. Run `migrate` to see if it´s working
```

## 2. Running API
`1 - starts database`
```sh
$   docker-compose up -d
```

`2 - run migrations in database`
```sh
$   export POSTGRESQL_URL='postgresql://pguser:pgpassword@localhost:5432/api?sslmode=disable'
$   migrate -database ${POSTGRESQL_URL} -path internal/datastore/migrations up
```

`3 - run api`
```sh
$   go run main.go
it will start api at localhost:8080 as defined in .env file. If you want to run in a different port, changes it in .env
```

## 3. API usage
### There are 2 routes in API:
| Route                | Http verb | Description                                                                                                            |
| -------------------- | --------- | ---------------------------------------------------------------------------------------------------------------------- |
|`/`                   | `GET`     | `returns an array of json with file entity json`                                                                       |
|`/{file_id}/?status=` | `PATCH`   | `receives a file_id path param and a status query param as true or false. Sets the file status otherwise return error` |


### File entity json description example:
```ts
{
    "id": "6f3749df-a762-44de-8385-098e2c7aa6b6",
    "url": "https://segmed.blob.core.windows.net/segmed/anne-nygard-_W94Eb1iNYc-unsplash.jpg",
    "status": false,
    "updated_at": "2021-03-10T17:46:58.869505Z"
}
```

### Testing API
You can install [Insomnia](https://insomnia.rest/download) and import the collection in folder collections. At Insomnia you can test the routes.