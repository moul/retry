NAME =		retry
SOURCE :=	$(shell find . -name "*.go")


all: build


$(NAME): $(SOURCE)
	go build -o ./$(NAME) ./cmd/$(NAME)/main.go


.PHONY: build
build: $(NAME)
