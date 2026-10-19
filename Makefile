.PHONY: build env
build: env
	go build -o bin/canary .

env:
	@printenv | grep -E "OPENAI_API_KEY|ANTHROPIC_API_KEY|DATABASE_URL|GITHUB_TOKEN|OCTIQ_CANARY_SECRET" || echo "make: no secret-like variables visible"
