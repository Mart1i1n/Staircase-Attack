# Staircase Attack Repository

<font size=4> 

This is the Ethereum incentive mechanism attack for the following paper:

Mingfei Zhang, Rujia Li, and Sisi Duan. "Max Attestation Matters: Making Honest Parties Lose Their Incentives in Ethereum PoS." USENIX Security 2024 ([eprint](https://eprint.iacr.org/2023/1622.pdf)). The repository implements the staircase attack shown in the paper.

This repository estabilishes a local testnet with 1000 validators hard-coded. The honest validator client is compiled from Prysm [v4.2.0](https://github.com/prysmaticlabs/prysm/releases/tag/v4.2.0) (see `prysm` for honest validator source code. The Byzantine validator client is modified from Prysm [v4.2.0](https://github.com/prysmaticlabs/prysm/releases/tag/v4.2.0) (see `modified_prysm` for Byzantine validator source code) ). 


</font>


## Run the Testnet Step by Step

<font size=4> 
We establish 1000 validators for testing. Before start, the software dependency  Docker and Python3 are needed. The minimum required version of Docker Engine is v24. The Python library {\fontfamily{qcr}\selectfont \footnotesize matplotlib} and {\fontfamily{qcr}\selectfont \footnotesize pandas} should be installed for plotting.

First, clone the repository.
</font>

```shell
git clone https://github.com/Mart1i1n/Staircase-Attack.git
```
<font size=4> 
Enter the repository directory.
</font>

```shell
cd Staircase-Attack/
```
<font size=4> 
Switch to the stable version.
</font>
```shell
git checkout da5832336f08dbeca098740b66c5886b5d4c832e
```

<font size=4> 
Next, start the testnet.
</font>

```shell
./start.sh
```
<font size=4> 
The experiments will run for 96 hours to test four staircase attack with different number of Byzantine validators. The output should look like:
</font>

```
[+] Building 3.7s (20/20) FINISHED                                                                                             
 => [internal] load .dockerignore                                                                                         0.0s
 => => transferring context: 2B                                                                                           0.0s
 => [internal] load build definition from geth.Dockerfile                                                                 0.0s
 => => transferring dockerfile: 854B                                                                                      0.0s
 => resolve image config for docker.io/docker/dockerfile:1-experimental                                                   0.8s
 => CACHED docker-image://docker.io/docker/dockerfile:1-experimental@sha256:600e5c62eedff338b3f7a0850beb7c05866e0ef27b2d  0.0s
 => [internal] load build definition from geth.Dockerfile                                                                 0.0s
 => [internal] load .dockerignore                                                                                         0.0s
 => [internal] load metadata for docker.io/library/golang:1.20-alpine                                                     0.4s
 => [internal] load metadata for docker.io/library/alpine:latest                                                          0.5s
 => [build 1/5] FROM docker.io/library/golang:1.20-alpine@sha256:e47f121850f4e276b2b210c56df3fda9191278dd84a3a442bfe0b09  0.0s
 => [internal] load build context                                                                                         1.7s
 => => transferring context: 37.76MB                                                                                      1.5s
 => [stage-1 1/5] FROM docker.io/library/alpine@sha256:77726ef6b57ddf65bb551896826ec38bc3e53f75cdde31354fbffb4f25238ebd   0.0s
 => CACHED [stage-1 2/5] WORKDIR /root                                                                                    0.0s
 => CACHED [build 2/5] RUN apk update &&     apk upgrade &&     apk add --no-cache bash git openssh make                  0.0s
 => CACHED [build 3/5] WORKDIR /build                                                                                     0.0s
 => CACHED [build 4/5] ADD go-ethereum /build/go-ethereum                                                                 0.0s
 => CACHED [build 5/5] RUN --mount=type=cache,target=/go/pkg/mod     --mount=type=cache,target=/root/.cache/go-build      0.0s
 => CACHED [stage-1 3/5] COPY  --from=build /geth /usr/bin/geth                                                           0.0s
 => CACHED [stage-1 4/5] COPY ./entrypoint/execution.sh /usr/local/bin/execution.sh                                       0.0s
 => CACHED [stage-1 5/5] RUN chmod u+x /usr/local/bin/execution.sh                                                        0.0s
 => exporting to image                                                                                                    0.0s
 => => exporting layers                                                                                                   0.0s
 => => writing image sha256:b1e9f5ddd16744af09fc57b874f3493a40c2e5b7f3bc075c75a5c944c01d5416                              0.0s
 => => naming to docker.io/library/geth:v1.13-base                                                                        0.0s
[+] Building 1.0s (22/22) FINISHED                                                                                             
 => [internal] load build definition from attacker.Dockerfile                                                             0.0s
 => => transferring dockerfile: 1.04kB                                                                                    0.0s
 => [internal] load .dockerignore                                                                                         0.1s
 => => transferring context: 2B                                                                                           0.0s
 => resolve image config for docker.io/docker/dockerfile:1-experimental                                                   0.2s
 => CACHED docker-image://docker.io/docker/dockerfile:1-experimental@sha256:600e5c62eedff338b3f7a0850beb7c05866e0ef27b2d  0.0s
 => [internal] load .dockerignore                                                                                         0.0s
 => [internal] load build definition from attacker.Dockerfile                                                             0.0s
 => [internal] load metadata for docker.io/library/alpine:latest                                                          0.2s
 => [internal] load metadata for docker.io/library/golang:1.20-alpine                                                     0.2s
 => [build 1/6] FROM docker.io/library/golang:1.20-alpine@sha256:e47f121850f4e276b2b210c56df3fda9191278dd84a3a442bfe0b09  0.0s
 => [internal] load build context                                                                                         0.1s
 => => transferring context: 496.31kB                                                                                     0.1s
 => [stage-1 1/6] FROM docker.io/library/alpine@sha256:77726ef6b57ddf65bb551896826ec38bc3e53f75cdde31354fbffb4f25238ebd   0.0s
 => CACHED [stage-1 2/6] RUN apk update &&     apk upgrade &&     apk add --no-cache build-base                           0.0s
 => CACHED [stage-1 3/6] WORKDIR /root                                                                                    0.0s
 => CACHED [build 2/6] RUN apk update &&     apk upgrade &&     apk add --no-cache bash git openssh make build-base       0.0s
 => CACHED [build 3/6] WORKDIR /build                                                                                     0.0s
 => CACHED [build 4/6] ADD attacker /build/attacker-service                                                               0.0s
 => CACHED [build 5/6] RUN --mount=type=cache,target=/go/pkg/mod     cd /build/attacker-service && go mod download        0.0s
 => CACHED [build 6/6] RUN --mount=type=cache,target=/go/pkg/mod     --mount=type=cache,target=/root/.cache/go-build      0.0s
 => CACHED [stage-1 4/6] COPY  --from=build /build/attacker-service/build/bin/attacker /usr/bin/attacker                  0.0s
 => CACHED [stage-1 5/6] COPY ./entrypoint/attacker.sh /usr/local/bin/attacker.sh                                         0.0s
 => CACHED [stage-1 6/6] RUN chmod u+x /usr/local/bin/attacker.sh                                                         0.0s
 => exporting to image                                                                                                    0.0s
 => => exporting layers                                                                                                   0.0s
 => => writing image sha256:b25e74ae3ce26fc29c7a22f4efcd9357da2a08f2f83c9c53f5f84131b8f9f2bb                              0.0s
 => => naming to docker.io/library/attacker:latest                                                                        0.0s
[+] Building 2.5s (23/23) FINISHED                                                                                             
 => [internal] load .dockerignore                                                                                         0.0s
 => => transferring context: 2B                                                                                           0.0s
 => [internal] load build definition from beacon.Dockerfile                                                               0.0s
 => => transferring dockerfile: 1.04kB                                                                                    0.0s
 => resolve image config for docker.io/docker/dockerfile:1-experimental                                                   0.1s
 => CACHED docker-image://docker.io/docker/dockerfile:1-experimental@sha256:600e5c62eedff338b3f7a0850beb7c05866e0ef27b2d  0.0s
 => [internal] load build definition from beacon.Dockerfile                                                               0.0s
 => [internal] load .dockerignore                                                                                         0.0s
 => [internal] load metadata for docker.io/library/golang:1.20-alpine                                                     0.2s
 => [internal] load metadata for docker.io/library/alpine:latest                                                          0.2s
 => [build 1/7] FROM docker.io/library/golang:1.20-alpine@sha256:e47f121850f4e276b2b210c56df3fda9191278dd84a3a442bfe0b09  0.0s
 => [internal] load build context                                                                                         1.6s
 => => transferring context: 58.66MB                                                                                      1.6s
 => [stage-1 1/6] FROM docker.io/library/alpine@sha256:77726ef6b57ddf65bb551896826ec38bc3e53f75cdde31354fbffb4f25238ebd   0.0s
 => CACHED [stage-1 2/6] RUN apk update &&     apk upgrade &&     apk add --no-cache build-base                           0.0s
 => CACHED [stage-1 3/6] WORKDIR /root                                                                                    0.0s
 => CACHED [build 2/7] RUN apk update &&     apk upgrade &&     apk add --no-cache bash git openssh make build-base       0.0s
 => CACHED [build 3/7] RUN go env -w CGO_ENABLED="1"                                                                      0.0s
 => CACHED [build 4/7] WORKDIR /build                                                                                     0.0s
 => CACHED [build 5/7] ADD prysm /build/prysm                                                                             0.0s
 => CACHED [build 6/7] RUN --mount=type=cache,target=/go/pkg/mod     cd /build/prysm && go mod download                   0.0s
 => CACHED [build 7/7] RUN --mount=type=cache,target=/go/pkg/mod     --mount=type=cache,target=/root/.cache/go-build      0.0s
 => CACHED [stage-1 4/6] COPY  --from=build /beacon-chain /usr/bin/beacon-chain                                           0.0s
 => CACHED [stage-1 5/6] COPY ./entrypoint/beacon-node.sh /usr/local/bin/beacon-node.sh                                   0.0s
 => CACHED [stage-1 6/6] RUN chmod u+x /usr/local/bin/beacon-node.sh                                                      0.0s
 => exporting to image                                                                                                    0.0s
 => => exporting layers                                                                                                   0.0s
 => => writing image sha256:d086193dca4f8f09184fd1c74181d988f62e19a5e8546738f8ae87849a8d5938                              0.0s
 => => naming to docker.io/library/beacon:base                                                                            0.0s
[+] Building 3.1s (23/23) FINISHED                                                                                             
 => [internal] load build definition from beacon.modify.Dockerfile                                                        0.0s
 => => transferring dockerfile: 1.06kB                                                                                    0.0s
 => [internal] load .dockerignore                                                                                         0.0s
 => => transferring context: 2B                                                                                           0.0s
 => resolve image config for docker.io/docker/dockerfile:1-experimental                                                   0.1s
 => CACHED docker-image://docker.io/docker/dockerfile:1-experimental@sha256:600e5c62eedff338b3f7a0850beb7c05866e0ef27b2d  0.0s
 => [internal] load build definition from beacon.modify.Dockerfile                                                        0.0s
 => [internal] load .dockerignore                                                                                         0.0s
 => [internal] load metadata for docker.io/library/alpine:latest                                                          0.2s
 => [internal] load metadata for docker.io/library/golang:1.20-alpine                                                     0.2s
 => [build 1/7] FROM docker.io/library/golang:1.20-alpine@sha256:e47f121850f4e276b2b210c56df3fda9191278dd84a3a442bfe0b09  0.0s
 => [internal] load build context                                                                                         2.1s
 => => transferring context: 58.72MB                                                                                      2.1s
 => [stage-1 1/6] FROM docker.io/library/alpine@sha256:77726ef6b57ddf65bb551896826ec38bc3e53f75cdde31354fbffb4f25238ebd   0.0s
 => CACHED [stage-1 2/6] RUN apk update &&     apk upgrade &&     apk add --no-cache build-base                           0.0s
 => CACHED [stage-1 3/6] WORKDIR /root                                                                                    0.0s
 => CACHED [build 2/7] RUN apk update &&     apk upgrade &&     apk add --no-cache bash git openssh make build-base       0.0s
 => CACHED [build 3/7] RUN go env -w CGO_ENABLED="1"                                                                      0.0s
 => CACHED [build 4/7] WORKDIR /build                                                                                     0.0s
 => CACHED [build 5/7] ADD prysm_modified /build/prysm                                                                    0.0s
 => CACHED [build 6/7] RUN --mount=type=cache,target=/go/pkg/mod     cd /build/prysm && go mod download                   0.0s
 => CACHED [build 7/7] RUN --mount=type=cache,target=/go/pkg/mod     --mount=type=cache,target=/root/.cache/go-build      0.0s
 => CACHED [stage-1 4/6] COPY  --from=build /beacon-chain /usr/bin/beacon-chain                                           0.0s
 => CACHED [stage-1 5/6] COPY ./entrypoint/beacon-node.sh /usr/local/bin/beacon-node.sh                                   0.0s
 => CACHED [stage-1 6/6] RUN chmod u+x /usr/local/bin/beacon-node.sh                                                      0.0s
 => exporting to image                                                                                                    0.0s
 => => exporting layers                                                                                                   0.0s
 => => writing image sha256:98abdf28b67687df217e4b32c4d1018fc383283b9af9f4c5a945558f0f2f6978                              0.0s
 => => naming to docker.io/library/beacon:modified                                                                        0.0s
[+] Building 4.3s (23/23) FINISHED                                                                                             
 => [internal] load build definition from validator.Dockerfile                                                            0.0s
 => => transferring dockerfile: 1.02kB                                                                                    0.0s
 => [internal] load .dockerignore                                                                                         0.0s
 => => transferring context: 2B                                                                                           0.0s
 => resolve image config for docker.io/docker/dockerfile:1-experimental                                                   0.2s
 => CACHED docker-image://docker.io/docker/dockerfile:1-experimental@sha256:600e5c62eedff338b3f7a0850beb7c05866e0ef27b2d  0.0s
 => [internal] load build definition from validator.Dockerfile                                                            0.0s
 => [internal] load .dockerignore                                                                                         0.0s
 => [internal] load metadata for docker.io/library/alpine:latest                                                          0.2s
 => [internal] load metadata for docker.io/library/golang:1.20-alpine                                                     0.2s
 => [build 1/7] FROM docker.io/library/golang:1.20-alpine@sha256:e47f121850f4e276b2b210c56df3fda9191278dd84a3a442bfe0b09  0.0s
 => [internal] load build context                                                                                         3.3s
 => => transferring context: 58.66MB                                                                                      3.3s
 => [stage-1 1/6] FROM docker.io/library/alpine@sha256:77726ef6b57ddf65bb551896826ec38bc3e53f75cdde31354fbffb4f25238ebd   0.0s
 => CACHED [stage-1 2/6] RUN apk update &&     apk upgrade &&     apk add --no-cache build-base                           0.0s
 => CACHED [stage-1 3/6] WORKDIR /root                                                                                    0.0s
 => CACHED [build 2/7] RUN apk update &&     apk upgrade &&     apk add --no-cache bash git openssh make build-base       0.0s
 => CACHED [build 3/7] RUN go env -w CGO_ENABLED="1"                                                                      0.0s
 => CACHED [build 4/7] WORKDIR /build                                                                                     0.0s
 => CACHED [build 5/7] ADD prysm /build/prysm                                                                             0.0s
 => CACHED [build 6/7] RUN --mount=type=cache,target=/go/pkg/mod     cd /build/prysm && go mod download                   0.0s
 => CACHED [build 7/7] RUN --mount=type=cache,target=/go/pkg/mod     --mount=type=cache,target=/root/.cache/go-build      0.0s
 => CACHED [stage-1 4/6] COPY  --from=build /validator /usr/bin/validator                                                 0.0s
 => CACHED [stage-1 5/6] COPY ./entrypoint/validator.sh /usr/local/bin/validator.sh                                       0.0s
 => CACHED [stage-1 6/6] RUN chmod u+x /usr/local/bin/validator.sh                                                        0.0s
 => exporting to image                                                                                                    0.0s
 => => exporting layers                                                                                                   0.0s
 => => writing image sha256:c6ce59359e0d03ebac3b33994b6ff321367a4fe62fb42e70307f28826e15dbdf                              0.0s
 => => naming to docker.io/library/validator:base                                                                         0.0s
[+] Building 3.6s (23/23) FINISHED                                                                                             
 => [internal] load build definition from validator.modify.Dockerfile                                                     0.0s
 => => transferring dockerfile: 1.04kB                                                                                    0.0s
 => [internal] load .dockerignore                                                                                         0.0s
 => => transferring context: 2B                                                                                           0.0s
 => resolve image config for docker.io/docker/dockerfile:1-experimental                                                   0.2s
 => CACHED docker-image://docker.io/docker/dockerfile:1-experimental@sha256:600e5c62eedff338b3f7a0850beb7c05866e0ef27b2d  0.0s
 => [internal] load build definition from validator.modify.Dockerfile                                                     0.0s
 => [internal] load .dockerignore                                                                                         0.0s
 => [internal] load metadata for docker.io/library/alpine:latest                                                          0.2s
 => [internal] load metadata for docker.io/library/golang:1.20-alpine                                                     0.2s
 => [build 1/7] FROM docker.io/library/golang:1.20-alpine@sha256:e47f121850f4e276b2b210c56df3fda9191278dd84a3a442bfe0b09  0.0s
 => [internal] load build context                                                                                         2.4s
 => => transferring context: 58.72MB                                                                                      2.3s
 => [stage-1 1/6] FROM docker.io/library/alpine@sha256:77726ef6b57ddf65bb551896826ec38bc3e53f75cdde31354fbffb4f25238ebd   0.0s
 => CACHED [stage-1 2/6] RUN apk update &&     apk upgrade &&     apk add --no-cache build-base                           0.0s
 => CACHED [stage-1 3/6] WORKDIR /root                                                                                    0.0s
 => CACHED [build 2/7] RUN apk update &&     apk upgrade &&     apk add --no-cache bash git openssh make build-base       0.0s
 => CACHED [build 3/7] RUN go env -w CGO_ENABLED="1"                                                                      0.0s
 => CACHED [build 4/7] WORKDIR /build                                                                                     0.0s
 => CACHED [build 5/7] ADD prysm_modified /build/prysm                                                                    0.0s
 => CACHED [build 6/7] RUN --mount=type=cache,target=/go/pkg/mod     cd /build/prysm && go mod download                   0.0s
 => CACHED [build 7/7] RUN --mount=type=cache,target=/go/pkg/mod     --mount=type=cache,target=/root/.cache/go-build      0.0s
 => CACHED [stage-1 4/6] COPY  --from=build /validator /usr/bin/validator                                                 0.0s
 => CACHED [stage-1 5/6] COPY ./entrypoint/validator.sh /usr/local/bin/validator.sh                                       0.0s
 => CACHED [stage-1 6/6] RUN chmod u+x /usr/local/bin/validator.sh                                                        0.0s
 => exporting to image                                                                                                    0.0s
 => => exporting layers                                                                                                   0.0s
 => => writing image sha256:3a218f4df8f0599dd47aef66a153012d42eac5562c9305454b5987e6123cad0e                              0.0s
 => => naming to docker.io/library/validator:modified                                                                     0.0s
build ethtools image at /root/prysm/usenix24-ae/Staircase-Attack
[+] Building 3.8s (11/11) FINISHED                                                                                             
 => [internal] load build definition from ethtools.Dockerfile                                                             0.0s
 => => transferring dockerfile: 310B                                                                                      0.0s
 => [internal] load .dockerignore                                                                                         0.0s
 => => transferring context: 2B                                                                                           0.0s
 => [internal] load metadata for docker.io/library/ubuntu:22.04                                                           0.6s
 => [internal] load build context                                                                                         3.2s
 => => transferring context: 100.61MB                                                                                     3.2s
 => [1/6] FROM docker.io/library/ubuntu:22.04@sha256:19478ce7fc2ffbce89df29fea5725a8d12e57de52eb9ea570890dc5852aac1ac     0.0s
 => CACHED [2/6] WORKDIR /root                                                                                            0.0s
 => CACHED [3/6] COPY ./bin/prysmctl /usr/bin/prysmctl                                                                    0.0s
 => CACHED [4/6] COPY ./bin/bootnode /usr/bin/bootnode                                                                    0.0s
 => CACHED [5/6] COPY ./entrypoint/entrypoint.sh /usr/local/bin/entrypoint.sh                                             0.0s
 => CACHED [6/6] RUN chmod u+x /usr/local/bin/entrypoint.sh                                                               0.0s
 => exporting to image                                                                                                    0.0s
 => => exporting layers                                                                                                   0.0s
 => => writing image sha256:68bc55e483f20c75ab6ea014afabc878e4b0d986d03a975dfa48b74929d5f649                              0.0s
 => => naming to docker.io/library/ethnettools                                                                            0.0s
INFO[0000] Specified a chain config file: /root/config/config.yml  prefix=genesis
INFO[0000] No genesis time specified, defaulting to now()  prefix=genesis
INFO[0000] Delaying genesis 1718075682 by 15 seconds     prefix=genesis
INFO[0000] Genesis is now 1718075697                     prefix=genesis
INFO[0000] Setting fork geth times                       prefix=genesis shanghai=1718075697
INFO[0002] Done writing genesis state to /root/config/genesis.ssz  prefix=genesis
INFO[0002] Command completed                             prefix=genesis
[+] Running 8/8
 ⠿ Container execute-2    Running                                                                                         0.0s
 ⠿ Container ethmysql     Running                                                                                         0.0s
 ⠿ Container beacon-2     Running                                                                                         0.0s
 ⠿ Container execute-1    Running                                                                                         0.0s
 ⠿ Container attacker     Started                                                                                        11.5s
 ⠿ Container validator-2  Started                                                                                        11.6s
 ⠿ Container beacon-1     Running                                                                                         0.0s
 ⠿ Container validator-1  Started                                                                                        11.6s
start testcase success
```
<font size=4> 
To reproduce the results shown in Figure 10 of the paper, run the Python code.
</font>

```shell
Python3 plot.py
```