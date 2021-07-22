TAG ?= latest
IMAGE_NAME ?= http-kit

ifdef IMAGE_REPO
	IMAGE_URL := $(IMAGE_REPO)/$(IMAGE_NAME):$(TAG)
else
	IMAGE_URL := $(IMAGE_NAME):$(TAG)
endif

PULL = PULL
all:
	docker build -t $(IMAGE_URL) --network host .

tag:
	docker tag $(IMAGE_URL) $(IMAGE):$(TAG)

clean:
	docker rmi $(IMAGE_URL)

push:
	docker push $(IMAGE_URL)

run:
	docker run --name $(IMAGE_NAME) -it -p 80 $(IMAGE_URL)

stop-and-rm:
	docker stop $(IMAGE_NAME) || true && docker rm $(IMAGE_NAME) || true

force-run: stop-and-rm
	docker run --name $(IMAGE_NAME) -it -p 80 $(IMAGE_URL)

runshell:
	docker exec -it $(IMAGE_NAME) sh


.PHONY: all tag clean push run stop-and-rm force-run
