image=rustvg

request:
	curl -X POST \
	-H "Content-Type: application/json" \
	-d @request.json http://localhost:8080/bbox

build:
	docker build -t $(image) .

run-server: build
	docker run --init -it --rm \
		--memory=128m --cpus=0.1 \
		-p 8080:8080 \
		$(image)

vegeta:
	vegeta attack -targets=request.txt -format=http -duration=52s -timeout=20s -rate=200 \
	| tee results.bin \
	| vegeta report
