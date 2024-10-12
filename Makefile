F_BUILD=gitlab-downloader

EXPORT_P=""

DIR_BUILD=build
DIR_CONFIG=~/.config/gitlab-downloader
DIR_BIN=~/.local/bin

#formating
NC=\033[0m
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[0;33m

BOLD=\033[1m
UNDERLINE=\033[4m


# user commands
all: build

build:
	@go build -o $(DIR_BUILD)/$(F_BUILD) ./cmd/cli/main.go

install: build
	@mkdir -p $(DIR_BIN)
	@mkdir -p $(DIR_CONFIG)
	@if [ -f $(DIR_CONFIG)/settings.json ]; then \
		echo "Хотите перезаписать существующий конфиг в $(BOLD)$(DIR_CONFIG)$(NC)? [y/N] "; \
		while true; do \
			read -s -n 1 response; \
			if [ "$$response" = "y" ] || [ "$$response" = "Y" ]; then \
				cp -f config/* $(DIR_CONFIG); \
				echo "$(YELLOW)✔ Конфиг перезаписан$(NC)"; \
				break; \
			elif [ "$$response" = "n" ] || [ "$$response" = "N" ]; then \
				echo "$(YELLOW)✔ Конфиг не перезаписан$(NC)"; \
				break; \
			else \
				echo "$(RED)Некорректный ввод. Пожалуйста, введите [y/N] $(NC)"; \
			fi; \
		done; \
	else \
	  	cp -f config/* $(DIR_CONFIG); \
	  	echo "$(YELLOW)✔ Конфиг создан в $(BOLD)$(DIR_CONFIG)$(NC)"; \
	fi
	@if ! grep -q 'export PATH="$$HOME/.local/bin:$$PATH"' ~/.zprofile; then \
  		echo 'export PATH="$$HOME/.local/bin:$$PATH"' >> ~/.zprofile; \
  	fi
	@cp $(DIR_BUILD)/$(F_BUILD) $(DIR_BIN)
	@chmod +x $(DIR_BIN)/$(F_BUILD)
	@echo "$(GREEN)✔ Программа $(F_BUILD) установлена в директорию $(DIR_BIN)$(NC)"
	@echo "Конфигурация программы находится в $(BOLD)$(DIR_CONFIG)$(NC)"
	@echo "Для применения изменений введите $(BOLD)source ~/.zprofile$(NC) или перезапустите терминал"
	@echo "Для запуска программы введите $(BOLD)$(F_BUILD)$(NC)"


uninstall:
	@rm -f $(DIR_BIN)/$(F_BUILD)
	@echo "Удалить конфиг из $(BOLD)$(DIR_CONFIG)$(NC)? [y/N] "; \
	while true; do \
  		read -s -n 1 response; \
		if [ "$$response" = "y" ] || [ "$$response" = "Y" ]; then \
			rm -rf $(DIR_CONFIG); \
			echo "$(YELLOW)✔ Конфиг удален$(NC)"; \
			break; \
		elif [ "$$response" = "n" ] || [ "$$response" = "N" ]; then \
			echo "$(YELLOW)✔ Конфиг сохранен в директории$(DIR_CONFIG)$(NC)"; \
			break; \
		else \
			echo "$(RED)Некорректный ввод. Пожалуйста, введите [y/N] $(NC)"; \
		fi; \
	done
	@rm -rf $(DIR_BUILD)
	@echo "$(GREEN)✔ Программа $(F_BUILD) удалена из директории $(DIR_BIN)$(NC)"

# developer tools
run:
	$(DIR_BUILD)/$(F_BUILD)

clean:
	rm -rf $(DIR_BUILD)

debug_run:
	go run ./cmd/cli/main.go