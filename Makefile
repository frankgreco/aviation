SERVICE_NAME 	?= aviation-service
AWS_REGION 		?= us-east-2
AWS_BUCKET_NAME	?= aviation-cloudformation
AWS_STACK_NAME 	?= $(SERVICE_NAME)-stack
AWS_TEMPLATE   	?= template.yaml

ALL_SRC		= $(shell find . -name "*.go" | grep -v -e vendor)
PACKAGES   	= $(shell go list ./... | grep -v -E 'vendor')
RACE       	= -race
GOTEST     	= go test -v $(RACE)
GOFMT      	= gofmt
PASS      	= $(shell printf "\033[32mPASS\033[0m")
FAIL      	= $(shell printf "\033[31mFAIL\033[0m")
COLORIZE   	= sed ''/PASS/s//$(PASS)/'' | sed ''/FAIL/s//$(FAIL)/''

.PHONY: deploy
deploy: 
	sam deploy \
		--template $(AWS_TEMPLATE) \
		--stack-name $(AWS_STACK_NAME) \
		--s3-bucket $(AWS_BUCKET_NAME) \
		--capabilities CAPABILITY_NAMED_IAM \
		--parameter-overrides "RdsUsername=$(RDS_USERNAME)" "RdsPassword=$(RDS_PASSWORD)" "RdsEndpoint=$(RDS_ENDPOINT)"

.PHONY: describe
describe:
	AWS_PAGER="" aws cloudformation describe-stacks \
			--region $(AWS_REGION) \
			--stack-name $(AWS_STACK_NAME)

.PHONY: up
up: binary
	sam local -t $(AWS_TEMPLATE)

.PHONY: invoke
invoke: binary
	sam local generate-event cloudwatch scheduled-event | sam local invoke -t $(AWS_TEMPLATE)

.PHONY: test
test:
	@bash -c "set -e; set -o pipefail; $(GOTEST) $(PACKAGES) | $(COLORIZE)"

.PHONY: install
install:
	GO111MODULE=on go mod download

.PHONY: fmt
fmt:
	@$(GOFMT) -e -s -l -w $(ALL_SRC)

.PHONY: binary
binary:
	@GOOS=linux GO111MODULE=on go build -o bin/$(GO_BINARY) ./$(GO_BINARY)

.PHONY: ui
ui:
	aws s3 cp ./ui/dist/ s3://aviation-website/ --recursive