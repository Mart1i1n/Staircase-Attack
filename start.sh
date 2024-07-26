#!/bin/bash

if [ $# -ne 1 ]; then
    echo "Usage: $0 <number>"
    exit 1
fi

input=$1
RESULTS_DIR="results"
mkdir -p "$RESULTS_DIR"

if ! [[ "$input" =~ ^[0-9]+$ ]]; then
    echo "Error: Input must be a natural number."
    exit 1
fi

if [ "$input" -eq 0 ]; then
    echo "Input is zero, executing basic test..."
    ./scripts/run.sh testcase/docker-compose-basic.yml
    sleep 1800
    ./scripts/stop.sh
    mv data/attacker/reward.csv "$RESULTS_DIR/basic.csv"
    echo "Basic test executed"
else
    echo "Input is not zero, executing experiment E1..."
    COMMAND_BASE="./scripts/run.sh"
    BASE_CONFIG_PATH="testcase/docker-compose-"
    INDEXES=(296 310 320 333)
    for INDEX in "${INDEXES[@]}"; do
        CONFIG_FILE="${BASE_CONFIG_PATH}${INDEX}.yml"
        echo "Running configuration: $CONFIG_FILE"
        $COMMAND_BASE "$CONFIG_FILE"
        sleep $((input * 3600))
        ./scripts/stop.sh
        mv data/attacker/reward.csv "${RESULTS_DIR}/${INDEX}.csv"
        echo "Moved reward.csv to ${RESULTS_DIR}/${INDEX}.csv"
    done
    echo "Experiment 1 executed"
fi