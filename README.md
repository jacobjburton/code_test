# Code Test

## NodeJS Version

navigate to the js folder: /code_test/js

Use the following to run test from cmd line (using jest)
```shell
npm install

npm run test
```

## Go Version

navigate to the go folder: /code_test/go

Use the following to run test from cmd line
```shell
go test -v -count=1 ./...
```

### Local Setup

#### Prerequisites

##### You will need these things installed on your host:

- go (1.16)
    - this is all you need to run the go test from the command line
    - go to https://golang.org/dl/ to install
- dnsmasq
- nginx

If you are on a MacBook, you can use [homebrew](https://brew.sh/) to accomplish this.
Once you have homebrew installed, (if you don't already), run `brew install dnsmasq nginx`. This command will probably take some time.
Homebrew will tell you that dnsmaq needs to run as root - this is okay and you should allow this.
Once you have those installed through homebrew it is recommended you run `brew services start dnsmasq`
and `brew services start nginx` to ensure that those services automatically start when you restart you computer.
##### You will also need:

- docker

On a MacBook, Docker Desktop for Mac is recommended.

### Setup

Create a file at the root go level (/code_test/go) name `.env`.

In your `.env` file, make sure the `APP_NAME` variables is set. If it isn't the application won't work.

The last part of the setup requires adding `127.0.0.1 bracketizer.test` to your `/etc/hosts/` file. On macOS this file is
read only, so you'll need to use `sudo`.
```shell
sudo sh -c "echo '127.0.0.1 bracketizer.test' >> /etc/hosts"
```

### Run
Use the `Makefile` to control your containers and services

Build containers and start services:
```shell
make local
```
Stop services and remove containers:
```shell
make down
```
Rebuild app container and restart service:
```shell
make reset
```

You can also view the other make commands available:
```shell
make help
```


#### Debugging
Docker has a built in way to let you view the logs from your service containers:
```shell
docker compose logs -f
```

##### In Postman
Health Check
- GET Request
- url: bracketizer.test/bracketizer/health
- expected response: "Health Check: Pass"

Check Brackets
- POST Request
- url: bracketizer.test/bracketizer/sweep
- body:
    - json
    - example:
    ```shell
    {
        "we_might_have_brackets": "{ I can has {}chzbrgr{} now }, { yis }"
    }
    ```