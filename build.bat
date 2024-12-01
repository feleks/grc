rm assets_contents.go && \
rm grc.exe && \
go run assets/store.go && \
go build -ldflags="-H windowsgui"
