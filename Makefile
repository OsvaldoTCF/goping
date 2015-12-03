GO=go

all: install

install:
	${GO} install

doc:
	@echo "Serving documentation on http://localhost:6060..."
	@godoc -http=:6060 -index
