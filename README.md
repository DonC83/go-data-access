# GO Data Access

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