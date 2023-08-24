NAME := color-make

BUILD := $(CURDIR)/build
DIST  := $(CURDIR)/dist

SYSTEM      != python -c 'import platform; print(platform.system().lower())'
MACHINE     != python -c 'import platform; print(platform.machine().lower())'
EXE         := $(if $(filter windows,$(SYSTEM)),.exe,)
DIST_TARGET := $(DIST)/$(NAME)-$(SYSTEM)-$(MACHINE)$(EXE)

all:

clean:
	@ $(RM) --recursive --verbose $(BUILD)
	@ $(RM) --recursive --verbose $(DIST)

dist: $(DIST_TARGET)

install:
	pipx install --force $(CURDIR)

pretty: black prettier

setup:
	conda install --yes libpython-static
	poetry install

#####################
# Auxiliary Targets #
#####################

black:
	isort --profile=black $(CURDIR)
	black $(CURDIR)

prettier:
	prettier --write $(CURDIR)

$(DIST_TARGET): $(CURDIR)/main.py
ifneq ($(SYSTEM),windows)
	python -m nuitka --standalone --onefile --output-filename=$(@F) --output-dir=$(@D) --remove-output $<
else
	pyinstaller --distpath=$(DIST) --workpath=$(BUILD) --onefile --name=$(NAME)-$(SYSTEM)-$(MACHINE) $<
endif
