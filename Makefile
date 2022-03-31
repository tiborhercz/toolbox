build:
			GOOS=js GOARCH=wasm go build -o static/public/main.wasm wasm/*.go && \
			cd static/public/ && npm run build && cd ../../ && \
			go build .