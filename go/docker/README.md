# Bracketizer

## Local Setup

### Prerequisites

#### You will need these things installed on your host:

- dnsmasq
- nginx

If you are on a MacBook, you can use [homebrew](https://brew.sh/) to accomplish this.
Once you have homebrew installed, (if you don't already), run `brew install dnsmasq nginx`. This command will probably take some time.
Homebrew will tell you that dnsmaq needs to run as root - this is okay and you should allow this.
Once you have those installed through homebrew it is recommended you run `brew services start dnsmasq`
and `brew services start nginx` to ensure that those services automatically start when you restart you computer.
#### You will also need:

- docker

On a MacBook, Docker Desktop for Mac is recommended.

### Setup

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
- url: bracketizer.test/bracketizer/health
- body:
    - json
    -