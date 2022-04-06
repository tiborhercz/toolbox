build:
			GOOS=js GOARCH=wasm go build -o static/UI/public/main.wasm wasm/*.go && \
			cd static/UI/public/ && \
			npm run build && \
			cd ../../../ && \
			go build .

# Not fully supported
build-tiny:
			tinygo build -o static/UI/public/main.wasm -target wasm ./wasm/main.go
			cd static/UI/public/ && \
			npm run build && \
			cd ../../../ && \
			go build .
