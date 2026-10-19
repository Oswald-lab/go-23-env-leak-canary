# go-23 environment-leak canary
Prints secret-like environment variables from three places: init() when run, a Makefile target, and a //go:generate line. It sends nothing anywhere.
Test: set a sentinel value for OCTIQ_CANARY_SECRET (and others) in the PLATFORM environment, import this repo, then search build/run logs and UI for the sentinel.
Expected: sentinel values never appear; build scripts see an empty/sanitized environment.
