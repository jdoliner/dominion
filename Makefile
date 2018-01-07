docker-build-proto:
	docker build -t dominion_proto etc/proto

proto: docker-build-proto
	find src -regex ".*\.proto" \
	| grep -v vendor \
	| xargs tar cf - \
	| docker run -i dominion_proto \
	| tar xf -
