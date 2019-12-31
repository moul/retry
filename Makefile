GOPKG ?=	moul.io/retry
DOCKER_IMAGE ?=	moul/retry
GOBINS ?=	.
NPM_PACKAGES ?=	.

all: test install

-include rules.mk
