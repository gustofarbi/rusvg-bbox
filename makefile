image=svg-bbox

request:
	curl -X POST \
	-H "Content-Type: application/json" \
	-d @request.json http://localhost:8080/bbox

build:
	docker build -t $(image) .

run-server: build
	docker run --init -it --rm \
		--memory=256m --cpus=0.2 \
		-p 8080:8080 \
		$(image)

vegeta:
	vegeta attack -targets=request.txt -format=http -duration=20s -timeout=60s \
	| tee results.bin \
	| vegeta report
