all: vault-update_linux_amd64.zip \
	vault-update_linux_386.zip \
	vault-update_linux_arm.zip \
	vault-update_linux_arm64.zip \
	vault-update_darwin_amd64.zip \
	vault-update_darwin_386.zip

vault-update_%.zip: main.go
	$(eval parts := $(subst _, ,$*))
	GOOS=$(word 1, $(parts)) GOARCH=$(word 2, $(parts)) CGO_ENABLED=0 go build -a -installsuffix cgo
	zip $@ vault-update
	rm -f vault-update
