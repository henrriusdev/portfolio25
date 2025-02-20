# Detectar el sistema operativo y arquitectura para TailwindCSS
UNAME_S := $(shell uname -s 2>/dev/null || echo Windows)
UNAME_M := $(shell uname -m 2>/dev/null || echo x64)

ifeq ($(UNAME_S), Windows)
    TAILWINDCSS_OS_ARCH := windows-x64.exe
    EXEC_TAILWIND := tailwindcss.exe
    RENAME_CMD := if exist tailwindcss-$(TAILWINDCSS_OS_ARCH) ren tailwindcss-$(TAILWINDCSS_OS_ARCH) tailwindcss.exe
else
    ifeq ($(UNAME_S), Darwin)
        ifeq ($(UNAME_M), arm64)
            TAILWINDCSS_OS_ARCH := macos-arm64
        else
            TAILWINDCSS_OS_ARCH := macos-x64
        endif
    else
        ifeq ($(UNAME_M), aarch64)
            TAILWINDCSS_OS_ARCH := linux-arm64
        else
            TAILWINDCSS_OS_ARCH := linux-x64
        endif
    endif
    EXEC_TAILWIND := ./tailwindcss
    RENAME_CMD := mv tailwindcss-$(TAILWINDCSS_OS_ARCH) tailwindcss
endif

# Rutas
TAILWIND_SRC := src/assets/tailwind.css
TAILWIND_OUT := src/public/styles/app.css

.PHONY: build-css
build-css: tailwindcss
	$(EXEC_TAILWIND) -i $(TAILWIND_SRC) -o $(TAILWIND_OUT) --minify

.PHONY: cover
cover:
	go tool cover -html=cover.out

.PHONY: start
start: build-css
	go run ./cmd/app

.PHONY: tailwindcss
tailwindcss:
	@if not exist $(EXEC_TAILWIND) ( \
		curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/latest/download/tailwindcss-$(TAILWINDCSS_OS_ARCH) && \
		$(RENAME_CMD) \
	)

.PHONY: test
test:
	go test -coverprofile=cover.out -shuffle on ./...

.PHONY: watch-css
watch-css: tailwindcss
	$(EXEC_TAILWIND) -i $(TAILWIND_SRC) -o $(TAILWIND_OUT) --watch