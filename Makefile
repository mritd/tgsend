BUILD_VERSION   := $(shell cat version)
BUILD_DATE      := $(shell date "+%F %T")
COMMIT_SHA1     := $(shell git rev-parse HEAD)

all:
	./cross_compile.sh

release: clean all
	ghr -u mritd -t ${GITHUB_TOKEN} -replace -recreate -name "Bump ${BUILD_VERSION}" --debug ${BUILD_VERSION} dist

clean:
	rm -rf dist

install:
	go install

.PHONY : all release clean install

.EXPORT_ALL_VARIABLES:

GO111MODULE = on
GOPROXY = https://goproxy.cn