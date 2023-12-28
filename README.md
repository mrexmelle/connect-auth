# connect-authx

## Compiling

```
$ make clean && make
```

## Building Docker image

```
$ make docker-image
```

## Pushing Docker image to Container Registry

Note that only the owner of the repository is allowed to push the image. 

```
$ make docker-release
```

## Running

### For local environment

```
$ docker pull postgres:15-alpine
$ docker run \
	-v $PWD/data:/var/lib/postgresql/data \
	-v $PWD/init-db:/docker-entrypoint-initdb.d \
	-p 5432:5432 \
	-e POSTGRES_PASSWORD=123 \
	--restart always \
	postgres:15-alpine
$ ./connect-authx serve
```

### For docker environment

```
$ docker compose up
```
Note that you cannot alter the docker image in the container registry. Only the owner of the repository is allowed to do so.

If error happens in `core` service due to failure to connect to database, restart it:
```
$ docker compose restart core
```
The failure happens due to `db` service isn't ready when `core` attempts to connect to it.
