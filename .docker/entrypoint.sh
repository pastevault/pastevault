#!/bin/sh

# Run Go server
./main &

# Run SvelteKit server
cd ./ui/build && node index.js &

# Wait for any process to exit
wait -n

# Exit with status of process that exited first
exit $?