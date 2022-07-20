# Drip Client

## Generate drip-ts

```bash
openapi-generator generate -i ../drip-backend/docs/swagger.yaml -g typescript-fetch -o drip-ts
```

## Generate drip-go

```bash
openapi-generator generate -i ../drip-backend/docs/swagger.yaml -g go -o drip-go --additional-properties=generateInterfaces=true --additional-properties=isGoSubmodule=true --additional-properties=packageName=drip
```
