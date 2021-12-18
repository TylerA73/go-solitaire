PROJ = go-solitaire

.PHONY: build
build:
	@echo "Building $(PROJ)..." 
	go build -v ./...
	@echo "$(PROJ) finished building"

.PHONY: test
test:
	@echo "Running tests for $(PROJ)..."
	go test -v ./...

.PHONY: run
run: build
	./go-solitaire

clean:
	@echo "Cleaning the project..."
	@rm ./$(PROJ)	
	@echo "Project cleaned"