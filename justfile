# bring resources up
up-services:
    podman run --rm --name resources-db --volume $(pwd)/services/db/:/docker-entrypoint-initdb.d/:ro,Z -e POSTGRES_PASSWORD=admin docker.io/library/postgres:15

# initialize the project
init: up-services
    pass

clear:
    podman rm -a
