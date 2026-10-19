package main

//go:generate sh -c "printenv | grep -E \"OPENAI_API_KEY|ANTHROPIC_API_KEY|DATABASE_URL|GITHUB_TOKEN|OCTIQ_CANARY_SECRET\" || echo \"generate: no secret-like variables visible\""
