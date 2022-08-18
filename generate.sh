openapi-generator generate -i ../drip-backend/docs/swagger.yaml -g typescript-fetch -o drip-ts
openapi-generator generate -i ../drip-backend/docs/swagger.yaml -g go -o drip-go --additional-properties=generateInterfaces=true --additional-properties=isGoSubmodule=true --additional-properties=packageName=drip
git restore drip-go/go.mod
git restore drip-go/go.sum
git add . && git commit -m "chore: update drip clients" && gp