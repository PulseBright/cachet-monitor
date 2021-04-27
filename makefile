
image_name=inputobject/cachet-monitor
image_tag=v3.0
image_version=${image_tag}

FULL_NAME=${image_name}:${image_tag}

build: 
	@docker build . --build-arg version=${image_version} -t ${FULL_NAME}

push: 
	@docker push ${FULL_NAME}

scan:
	@docker scan --accept-license ${FULL_NAME}

package:
	@./go-executable-build.sh cli/main.go

all: build scan push