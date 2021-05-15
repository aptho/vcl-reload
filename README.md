# Varnish VCL reload

Simple go script that watches a directory for changes and reloads vcl

### How it works
[Go script](./main.go) watches files in `/etc/varnish/default.vcl`, and calls varnishreload if changes are detected

[Dockerfile](./Dockerfile) extends `varnish:latest` to include the go script

[docker-compose.yml](./docker-compose.yml) runs the docker container and mounts the [varnish](./varnish) directory into `/etc/varnish`
