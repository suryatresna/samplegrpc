ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find proto -name *.proto -not -path '*/vendor/*'")
else
	API_PROTO_FILES=$(shell find proto -name *.proto -not -path '*/vendor/*')
endif

init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
generate:
	protoc -I . --proto_path=. \
 	       --go_out=paths=source_relative:. \
		   --go-grpc_out=paths=source_relative:. \
	       $(API_PROTO_FILES)