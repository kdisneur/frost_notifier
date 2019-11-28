CREATED_AT := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
RELEASE_NAME := $(shell basename $$(git describe --all))

GIT_OBJECT := HEAD
GIT_SHORT_SHA := $(shell git rev-parse --short $(GIT_OBJECT))

GIT_SHA := $(shell git rev-parse $(GIT_OBJECT))
GIT_REPOSITORY := https://github.com/$(shell git remote get-url origin | sed -e 's/.*://' -e 's/\\.git//')

GO_VERSION := $(shell awk '/^golang / {print $$2}' .tool-versions)
GO_FMT_BIN := gofmt

KUBECTL_BIN := kubectl

build-image:
	@sed -e "s/<GO_VERSION>/$(GO_VERSION)/" Dockerfile.in \
	  | docker build \
		  --build-arg GIT_REPOSITORY=$(GIT_REPOSITORY) \
		  --build-arg GIT_SHA=$(GIT_SHA) \
		  --build-arg CREATED_AT=$(CREATED_AT) \
		  --tag kdisneur/frost-notifier:$(RELEASE_NAME)-$(GIT_SHORT_SHA) \
		  --tag kdisneur/frost-notifier:$(RELEASE_NAME)-latest \
		  --file - \
		  .
	@docker push kdisneur/frost-notifier:$(RELEASE_NAME)-$(GIT_SHORT_SHA)
	@docker push kdisneur/frost-notifier:$(RELEASE_NAME)-latest

deploy:
	@sed -e 's/<IMAGE_TAG>/$(RELEASE_NAME)-$(GIT_SHORT_SHA)/' \
		-e "s/<OPENWEATHER_APIKEY>/$${OPENWEATHER_APIKEY}/" \
		-e "s/<TWILIO_ACCOUNTSID>/$${TWILIO_ACCOUNTSID}/"\
		-e "s/<TWILIO_TOKEN>/$${TWILIO_TOKEN}/" \
		-e "s/<TWILIO_VIRTUALPHONENUMBER>/$${TWILIO_VIRTUALPHONENUMBER}/" \
		-e "s/<LANGUAGE>/$${LANGUAGE}/" \
		-e "s/<COUNTRY>/$${COUNTRY}/" \
		-e "s/<PHONE>/$${PHONE}/" \
		-e "s/<POSTCODE>/$${POSTCODE}/" \
		k8s/kustomization.yaml.in > k8s/kustomization.yaml

	$(KUBECTL_BIN) apply -k ./k8s

test-unit:
	@go test -race ./...

test-format:
	@output=$$($(GO_FMT_BIN) -l main.go internal); \
	if [ -n "$${output}" ]; then \
		echo "$${output}"; \
		exit 1; \
	fi

regenerate-code:
	go generate ./...
