# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: geuth android ios geuth-cross evm all test clean
.PHONY: geuth-linux geuth-linux-386 geuth-linux-amd64 geuth-linux-mips64 geuth-linux-mips64le
.PHONY: geuth-linux-arm geuth-linux-arm-5 geuth-linux-arm-6 geuth-linux-arm-7 geuth-linux-arm64
.PHONY: geuth-darwin geuth-darwin-386 geuth-darwin-amd64
.PHONY: geuth-windows geuth-windows-386 geuth-windows-amd64

GOBIN = ./build/bin
GO ?= latest
GORUN = env GO111MODULE=on go run

geuth:
	$(GORUN) build/ci.go install ./cmd/geuth
	@echo "Done building."
	@echo "Run \"$(GOBIN)/geuth\" to launch geuth."

all:
	$(GORUN) build/ci.go install

android:
	$(GORUN) build/ci.go aar --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/geuth.aar\" to use the library."
	@echo "Import \"$(GOBIN)/geuth-sources.jar\" to add javadocs"
	@echo "For more info see https://stackoverflow.com/questions/20994336/android-studio-how-to-attach-javadoc"

ios:
	$(GORUN) build/ci.go xcode --local
	@echo "Done building."
	@echo "Import \"$(GOBIN)/Geuth.framework\" to use the library."

test: all
	$(GORUN) build/ci.go test

lint: ## Run linters.
	$(GORUN) build/ci.go lint

clean:
	env GO111MODULE=on go clean -cache
	rm -fr build/_workspace/pkg/ $(GOBIN)/*

# The devtools target installs tools required for 'go generate'.
# You need to put $GOBIN (or $GOPATH/bin) in your PATH to use 'go generate'.

devtools:
	env GOBIN= go install golang.org/x/tools/cmd/stringer@latest
	env GOBIN= go install github.com/kevinburke/go-bindata/go-bindata@latest
	env GOBIN= go install github.com/fjl/gencodec@latest
	env GOBIN= go install github.com/golang/protobuf/protoc-gen-go@latest
	env GOBIN= go install ./cmd/abigen
	@type "solc" 2> /dev/null || echo 'Please install solc'
	@type "protoc" 2> /dev/null || echo 'Please install protoc'

# Cross Compilation Targets (xgo)

geuth-cross: geuth-linux geuth-darwin geuth-windows geuth-android geuth-ios
	@echo "Full cross compilation done:"
	@ls -ld $(GOBIN)/geuth-*

geuth-linux: geuth-linux-386 geuth-linux-amd64 geuth-linux-arm geuth-linux-mips64 geuth-linux-mips64le
	@echo "Linux cross compilation done:"
	@ls -ld $(GOBIN)/geuth-linux-*

geuth-linux-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/386 -v ./cmd/geuth
	@echo "Linux 386 cross compilation done:"
	@ls -ld $(GOBIN)/geuth-linux-* | grep 386

geuth-linux-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/amd64 -v ./cmd/geuth
	@echo "Linux amd64 cross compilation done:"
	@ls -ld $(GOBIN)/geuth-linux-* | grep amd64

geuth-linux-arm: geuth-linux-arm-5 geuth-linux-arm-6 geuth-linux-arm-7 geuth-linux-arm64
	@echo "Linux ARM cross compilation done:"
	@ls -ld $(GOBIN)/geuth-linux-* | grep arm

geuth-linux-arm-5:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-5 -v ./cmd/geuth
	@echo "Linux ARMv5 cross compilation done:"
	@ls -ld $(GOBIN)/geuth-linux-* | grep arm-5

geuth-linux-arm-6:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-6 -v ./cmd/geuth
	@echo "Linux ARMv6 cross compilation done:"
	@ls -ld $(GOBIN)/geuth-linux-* | grep arm-6

geuth-linux-arm-7:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm-7 -v ./cmd/geuth
	@echo "Linux ARMv7 cross compilation done:"
	@ls -ld $(GOBIN)/geuth-linux-* | grep arm-7

geuth-linux-arm64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/arm64 -v ./cmd/geuth
	@echo "Linux ARM64 cross compilation done:"
	@ls -ld $(GOBIN)/geuth-linux-* | grep arm64

geuth-linux-mips:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips --ldflags '-extldflags "-static"' -v ./cmd/geuth
	@echo "Linux MIPS cross compilation done:"
	@ls -ld $(GOBIN)/geuth-linux-* | grep mips

geuth-linux-mipsle:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mipsle --ldflags '-extldflags "-static"' -v ./cmd/geuth
	@echo "Linux MIPSle cross compilation done:"
	@ls -ld $(GOBIN)/geuth-linux-* | grep mipsle

geuth-linux-mips64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips64 --ldflags '-extldflags "-static"' -v ./cmd/geuth
	@echo "Linux MIPS64 cross compilation done:"
	@ls -ld $(GOBIN)/geuth-linux-* | grep mips64

geuth-linux-mips64le:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=linux/mips64le --ldflags '-extldflags "-static"' -v ./cmd/geuth
	@echo "Linux MIPS64le cross compilation done:"
	@ls -ld $(GOBIN)/geuth-linux-* | grep mips64le

geuth-darwin: geuth-darwin-386 geuth-darwin-amd64
	@echo "Darwin cross compilation done:"
	@ls -ld $(GOBIN)/geuth-darwin-*

geuth-darwin-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=darwin/386 -v ./cmd/geuth
	@echo "Darwin 386 cross compilation done:"
	@ls -ld $(GOBIN)/geuth-darwin-* | grep 386

geuth-darwin-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=darwin/amd64 -v ./cmd/geuth
	@echo "Darwin amd64 cross compilation done:"
	@ls -ld $(GOBIN)/geuth-darwin-* | grep amd64

geuth-windows: geuth-windows-386 geuth-windows-amd64
	@echo "Windows cross compilation done:"
	@ls -ld $(GOBIN)/geuth-windows-*

geuth-windows-386:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=windows/386 -v ./cmd/geuth
	@echo "Windows 386 cross compilation done:"
	@ls -ld $(GOBIN)/geuth-windows-* | grep 386

geuth-windows-amd64:
	$(GORUN) build/ci.go xgo -- --go=$(GO) --targets=windows/amd64 -v ./cmd/geuth
	@echo "Windows amd64 cross compilation done:"
	@ls -ld $(GOBIN)/geuth-windows-* | grep amd64
