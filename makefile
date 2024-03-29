image=rustvg

request:
	curl -X POST \
	-H "Content-Type: application/json" \
	-d @request.json http://localhost:8080/bbox

build:
	docker build -t $(image) .

run-server: build
	docker run --init -it --rm \
		--memory=64m --cpus=0.05 \
		-p 8080:8080 \
		$(image)

vegeta:
	vegeta attack -targets=request.txt -format=http -duration=18s -timeout=20s -rate=60 \
	| tee results.bin \
	| vegeta report

vegeta-max:
	vegeta attack -targets=request.txt -format=http -duration=50s -timeout=20s -rate=0 -max-workers=200 \
	| tee results.bin \
	| vegeta report
