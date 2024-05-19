#!/bin/bash

dotenv() {
    if [ -f .env ]; then
        while IFS='=' read -r key value; do
            if [[ ! -z "$key" && "$key" != \#* ]]; then
                export "$key=$value"
            fi
        done < .env
    else
        echo ".env file not found."
    fi
}


dotenv