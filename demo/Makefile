MAKEFLAGS += --jobs --output-sync

ASSETS   := $(CURDIR)/assets
GIF      := $(CURDIR)/gif
ROOT     != git rev-parse --show-toplevel
TAPE     := $(CURDIR)/tape
TEMPLATE := $(ROOT)/template

ASSETS_LIST += $(ASSETS)/echo.png
ASSETS_LIST += $(ASSETS)/error.png
ASSETS_LIST += $(ASSETS)/warning.png
ASSETS_LIST += $(ASSETS)/info.png

all: $(ASSETS_LIST)

clean:
	@ $(RM) --recursive --verbose $(ASSETS)
	@ $(RM) --recursive --verbose $(GIF)
	@ $(RM) --recursive --verbose $(TAPE)

deploy: $(TEMPLATE)/scripts/deploy-gh-pages.sh $(ASSETS_LIST)
	bash $< $(ASSETS) assets

#####################
# Auxiliary Targets #
#####################

echo:
	echo "Hello, world!"

error:
	$(error ERROR MESSAGE)

warning:
	$(warning WARNING MESSAGE)

info:
	$(info INFO MESSAGE)

$(ASSETS) $(GIF) $(TAPE):
	@ mkdir --parents --verbose $@

$(ASSETS)/%.png: $(GIF)/%.gif | $(ASSETS)
	magick convert $< -coalesce -background transparent -flatten $@

$(GIF)/%.gif: $(TAPE)/%.tape | $(GIF)
	MAKELEVEL=0 vhs $< --output $@

$(TAPE)/%.tape: $(CURDIR)/template.tape | $(TAPE)
	sed --expression="s/TARGET/$*/g" $< > $@
