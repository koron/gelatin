PACKAGES = \
	./ahocorasick \
	./omap \
	./trie

test:
	go test $(PACKAGES)
