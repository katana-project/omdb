package omdb

//go:generate curl -s -o openapi.json https://converter.swagger.io/api/convert?url=https://www.omdbapi.com/swagger.json
//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest --config models.cfg.yaml openapi.json
//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest --config client.cfg.yaml openapi.json
