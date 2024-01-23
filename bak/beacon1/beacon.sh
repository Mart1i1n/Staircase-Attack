#! /bin/bash

./beacon-chain --datadir=beacondata --min-sync-peers=0 --genesis-state=../genesis.ssz --bootstrap-node= --chain-config-file=../config.yml --config-file=../config.yml --chain-id=32382 --execution-endpoint=http://localhost:8551 -accept-terms-of-use --jwt-secret=../gethdata/geth/jwtsecret
