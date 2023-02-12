include Makefile.common

##############################################
# Multi related targets and variables
##############################################
ALL_MODULES := $(shell find . -name go.mod -type f -exec dirname {} \; | sort)

.PHONY: $(ALL_MODULES)
$(ALL_MODULES): 
	@$(MAKE) -C $@ info $(TARGET)

.PHONY : for-all-target
for-all-target: $(ALL_MODULES)

.PHONY: all-common
all-common:
	@$(MAKE) for-all-target TARGET="common"

.PHONY: all-tidy
all-tidy:
	@$(MAKE) for-all-target TARGET="tidy"

.PHONY: all-download
all-download:
	@$(MAKE) for-all-target TARGET="download"

.PHONY: all-test
all-test:
	@$(MAKE) for-all-target TARGET="test"

.PHONY: all-porto
all-porto:
	@$(MAKE) for-all-target TARGET="porto"

.PHONY: all-update-deps
all-update-deps:
	@$(MAKE) for-all-target TARGET="update-deps"