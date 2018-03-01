
.PHONY: clean build dockerize

build:
	$(MAKE) -C message-transport	build
	$(MAKE) -C kubernetes-crds		build
	$(MAKE) -C function-controller	build
	$(MAKE) -C function-sidecar		build
	$(MAKE) -C http-gateway			build
	$(MAKE) -C topic-controller		build
	$(MAKE) -C riff-cli				build

dockerize:
	$(MAKE) -C function-controller	dockerize
	$(MAKE) -C function-sidecar		dockerize
	$(MAKE) -C http-gateway			dockerize
	$(MAKE) -C topic-controller		dockerize


vendor: glide.lock
	glide install -v --force

glide.lock: glide.yaml
	glide up -v --force

clean:
	$(MAKE) -C function-controller	clean
	$(MAKE) -C function-sidecar		clean
	$(MAKE) -C http-gateway			clean
	$(MAKE) -C topic-controller		clean
	$(MAKE) -C riff-cli				clean

