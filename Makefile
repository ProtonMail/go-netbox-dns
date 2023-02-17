all: generate

swagger.json: swagger_original.json scripts/swagger_modifier.py
	./scripts/swagger_modifier.py

go-swagger:
	go install github.com/go-swagger/go-swagger/cmd/swagger@v0.30.4

generate: go-swagger
	swagger generate client --target=./netbox_dns

clean:
	$(RM) -r netbox_dns/client netbox_dns/models

.PHONY: all generate go-swagger clean
