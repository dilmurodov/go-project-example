CURRENT_DIR = $(shell pwd)

swag-init:
	swag init -g api/api.go -o api/docs

gen-proto:
	./scripts/gen-proto.sh ${CURRENT_DIR}

copy-proto-module:
	rm -rf ${CURRENT_DIR}/protos
	rsync -rv --exclude={'/.git','LICENSE','README.md'} ${CURRENT_DIR}/pb_protos/* ${CURRENT_DIR}/protos
