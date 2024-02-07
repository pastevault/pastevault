# Run Go server
./main &

# Run SvelteKit server
node ui/build/index.js &

# Wait for any process to exit
wait -n

# Exit with status of process that exited first
exit $?