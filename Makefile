.PHONY: gen
gen: dcl-attend-frontend/src/services/models.ts

# Depends on any files in model
dcl-attend-frontend/src/services/models.ts: $(wildcard internal/model/*) scripts/typesgen/main.go
	# This tmp dir would not work on windows ¯\_(ツ)_/¯
	go run scripts/typesgen/main.go > /tmp/models.ts && \
	mv /tmp/models.ts $@
	# TODO: Run TS formatter on the file here