git stash
git checkout main
git pull

cd ../drip-backend && git stash && git checkout main && git pull && cd ../drip-client

openapi-generator generate -i ../drip-backend/docs/swagger.yaml -g typescript-fetch -o drip-ts
openapi-generator generate -i ../drip-backend/docs/swagger.yaml -g go -o drip-go --additional-properties=generateInterfaces=true --additional-properties=isGoSubmodule=true --additional-properties=packageName=drip
git restore drip-go/go.mod
git restore drip-go/go.sum

cd ../drip-cloud-functions && git stash && git checkout main && git pull && cd ./functions && npm run build

openapi-generator generate -i ../drip-cloud-functions/functions/swagger.yaml -g go -o drip-extension-go --additional-properties=generateInterfaces=true --additional-properties=isGoSubmodule=true --additional-properties=packageName=dripextension
git restore drip-extension-go/go.mod
git restore drip-extension-go/go.sum

git add . && git commit -m "chore: update drip clients" && git push