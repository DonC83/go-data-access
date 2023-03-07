# GO Data Access

This repo piggybacks off the back of the Golang tutorials found [here](https://go.dev/doc/tutorial/)

It builds off the simple database tutorial and uses MySql and Postgres in docker
to start off with.

Getting colima and docker installed

```
brew install docker
brew install colima
brew install docker-compose
```

Install the docker credential helper to allow docker to make use of the osxkeychain
```
brew install docker-credential-helper
```

Check the `~/.docker/config.json` and change the `credsStore` to `oskkeychain`

```
{
        "auths": {
                "https://index.docker.io/v1/": {}
        },
        "credsStore": "osxkeychain",
        "currentContext": "colima"
}
```

Start colima 
```
colima start
```

Start database
```
make devstack.start
```





Get on to the docker image bash
```
docker exec -it c8a3220b7c5e bash

```