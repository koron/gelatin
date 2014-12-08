PACKAGES = \
	./ahocorasick \
	./args \
	./omap \
	./trie

test:
	go test $(PACKAGES)
