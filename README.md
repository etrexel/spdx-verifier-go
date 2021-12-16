# SPDX Verifier Go
This is an implementation of a verifier that checks license data from an SPDX file.

## Build
```shell
go build -o ./bin/spdx_verifier ./cmd
```

## Generating an SPDX file
There are a handful of tools that can be used to generate SPDX files for docker images:
- [syft](https://github.com/anchore/syft)
- [tern](https://github.com/tern-tools/tern)

First build a container to generate an SPDX file for
```shell
docker build -t localhost:5000/net-monitor:v1 https://github.com/wabbit-networks/net-monitor.git#main
```

### Syft
```shell
syft localhost:5000/net-monitor:v1 -o spdx --file sbom.spdx
```

### Tern
```shell
tern report -f spdxtagvalue -o sbom.spdx -i localhost:5000/net-monitor:v1
```

## Listing Licenses
```shell
./bin/spdx_verifier list sbom.spdx

Document: localhost-5000/net-monitor-v1
  Package:                        busybox, License:                   GPL-2.0-only
  Package:         ca-certificates-bundle, License:                MPL-2.0 AND MIT
  Package:                        scanelf, License:                   GPL-2.0-only
  Package:                    alpine-keys, License:                            MIT
  Package:                   libcrypto1.1, License:                        OpenSSL
  Package:                     musl-utils, License:                            MIT
  Package:                      libssl1.1, License:                        OpenSSL
  Package:                           zlib, License:                           Zlib
  Package:              alpine-baselayout, License:                   GPL-2.0-only
  Package:                      apk-tools, License:                   GPL-2.0-only
  Package:                     libc-utils, License:  BSD-2-Clause AND BSD-3-Clause
  Package:                       libretls, License:                            ISC
  Package:                           musl, License:                            MIT
  Package:                     ssl_client, License:                   GPL-2.0-only

```

## Verifying Licenses
```shell
./bin/spdx_verifier verify -l allowed_licenses.txt sbom.spdx

Document: localhost-5000/net-monitor-v1
VERIFICATION SUCCESS: all packages have acceptable licenses
```

Remove some lines from allowed_licenses.txt and run again:
```shell
./bin/spdx_verifier verify -l allowed_licenses.txt sbom.spdx

Document: localhost-5000/net-monitor-v1
VERIFICATION FAILED: some packages contain a license that is not allowed
Problem packages:
  Package:                           zlib, License:                           Zlib
  Package:                    alpine-keys, License:                            MIT
  Package:                     musl-utils, License:                            MIT
  Package:                       libretls, License:                            ISC
  Package:                           musl, License:                            MIT
```
