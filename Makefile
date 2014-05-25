PACKAGES = \
	./ahocorasick \
	./trie

test:
	go test $(PACKAGES)
