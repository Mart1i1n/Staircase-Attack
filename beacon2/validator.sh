#! /bin/bash

./validator   --datadir=validatordata   --accept-terms-of-use   --interop-start-index=64   --interop-num-validators=64   --force-clear-db   --beacon-rpc-provider=127.0.0.1:4001   --grpc-gateway-port=3501   --chain-config-file=../config.yml   --config-file=../config.yml
