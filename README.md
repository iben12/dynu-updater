# Dynu Dynamic DNS IP updater

## Minimum Viable Product

A simple Go program to update domain IP address on [Dynu](https://dynu.com). It uses [MyIP API](https://api.myip.com) to detect the current external IP address of the host and sets it on the defined Dynu domain.

It accepts plain password (as `PASSWORD`) or password MD5/SHA256 hashes (as `PASSWORD_HASH`) from env, prefers hash if both are set.

> **Note**
> The program sends the password to the APi in hashed format even it is provided as plain text.

## Usage

### Binary

- Execute `make build`
- Set environment variables:
  
  ```shell
  export USERNAME=testuser
  export PASSWORD=testpass # or
  export PASSWORD_HASH=testhash
  export DOMAIN=example.com
  export PERIOD_HOURS=1 # Update period in hours
  ```

- Run the binary created in `./build/dynu-updater`

### Docker

- Execute `make d-build`
- Copy the `.env_example` file to `.env` and replace test config
- Execute `make d-run`

Or simply run the public Docker image (it has `x86/amd64` and `arm64` versions) with:

```shell
docker run ghcr.io/iben12/dynu-updater \
  -e USERNAME=testuser \
  -e PASSWORD=testpass \ # or
  -e PASSWORD_HASH=testhash \
  -e DOMAIN=example.com \
  -e PERIOD_HOURS=1
```
