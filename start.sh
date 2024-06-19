#!/bin/bash

COMMAND_BASE="./scripts/run.sh"
BASE_CONFIG_PATH="testcase/docker-compose-"
INDEXES=(296 310 320 333)

RESULTS_DIR="results"
mkdir -p $RESULTS_DIR

for INDEX in "${INDEXES[@]}"; do

    CONFIG_FILE="${BASE_CONFIG_PATH}${INDEX}.yml"
    $COMMAND_BASE $CONFIG_FILE
    sleep 87000
    ./scripts/stop.sh
    mv data/attacker/reward.csv ${RESULTS_DIR}/${INDEX}.csv
    
done
