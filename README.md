# html-js-cv #

Generated html with:

```bash
cd ./tooling/
go run ./generator.go > ../index.html
```

# Install go (under fedora linux) #

```bash
dnf install -y golang-bin

```

# update website

See: [ansible/README.md](ansible/README.md)

# Server certificate by hand

```bash
cd /opt/
./certbot-auto renew --dry-run
./certbot-auto renew
curl -vvI https://olaf-radicke.de/
nginx -t
```