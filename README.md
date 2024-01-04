

This project is deprecate!
==========================

This is the old code of the website olaf-radicke.de.

## Old notes of html-js-cv

Generated html with:

```bash
cd ./tooling/
go run ./generator.go > ../index.html
```

## Install go (under fedora linux)

```bash
dnf install -y golang-bin

```

## update website

See: [ansible/README.md](ansible/README.md)

## Server certificate by hand

```bash
cd /opt/
./certbot-auto renew --dry-run
./certbot-auto renew
curl -vvI https://olaf-radicke.de/
nginx -t
```

## Create container image

Enter

```bash
podman build -t olaf-radicke-de:latest --no-cache=false .
```

Test run:

```
podman run --name olaf-radicke-de:latest -d -p 8080:80 --rm olaf-radicke-de
```

Push image:

```bash
podman login docker.io
podman tag  olaf-radicke-de:latest  olafradicke/olaf-radicke-de:1.0
podman push olafradicke/olaf-radicke-de:1.0
```