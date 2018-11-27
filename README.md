# Userdotd

## Docker Testing Images

Run and Install base testing image

```bash
# build
$ docker build -t userdotd/base -f ./docker/base/Dockerfile .
# run with shell
$ docker run --rm -it userdotd/base
```

Run and Install fish testing image w/ the fish shell

```bash
# build
$ docker build -t userdotd/fish -f ./docker/fish/Dockerfile .
# run with shell
$ docker run --rm -it userdotd/fish
```
