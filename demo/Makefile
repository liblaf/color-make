MAKEFLAGS += --jobs

ASSETS   := $(CURDIR)/assets
FRAMES   := $(CURDIR)/frames
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
	$(RM) --recursive $(ASSETS)
	$(RM) --recursive $(FRAMES)
	$(RM) --recursive $(GIF)
	$(RM) --recursive $(TAPE)

deploy: $(TEMPLATE)/scripts/deploy-gh-pages.sh $(ASSETS_LIST)
	bash $< $(ASSETS) assets

echo:
	echo "Hello, world!"

error:
	$(error ERROR MESSAGE)

warning:
	$(warning WARNING MESSAGE)

info:
	$(info INFO MESSAGE)

$(GIF) $(TAPE):
	@ mkdir --parents --verbose $@

$(ASSETS)/%.png: $(FRAMES)/%/frame-100.png
	@ install -D --mode="u=rw,go=r" --no-target-directory --verbose $< $@

$(FRAMES)/%/frame-100.png: $(GIF)/%.gif
	@ mkdir --parents --verbose $(@D)
	magick convert -coalesce $< $(@D)/frame-%03d.png

$(GIF)/%.gif: $(TAPE)/%.tape
	vhs $< --output $@

$(TAPE)/%.tape: $(CURDIR)/template.tape | $(TAPE)
	sed --expression 's/TARGET/$*/g' $< > $@
