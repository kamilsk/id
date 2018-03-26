OPEN_BROWSER       =
SUPPORTED_VERSIONS = 1.9 1.10 latest

include makes/env.mk
include makes/local.mk
include makes/docker.mk
include env/cmd.mk
include env/docker.mk
include env/docker-compose.mk
include env/tools.mk

.PHONY: code-quality-check
code-quality-check: ARGS = \
	--exclude=".*_test\.go:.*error return value not checked.*\(errcheck\)$$" \
	--exclude="duplicate of.*_test.go.*\(dupl\)$$" \
	--exclude="static/bindata.go" \
	--vendor --deadline=5m ./... | sort
code-quality-check: docker-tool-gometalinter

.PHONY: code-quality-report
code-quality-report:
	time make code-quality-check | tail +7 | tee report.out | pbcopy


.PHONY: pull-github-tpl
pull-github-tpl:
	rm -rf .github
	git clone git@github.com:kamilsk/shared.git .github
	( \
	  cd .github && \
	  git checkout github-tpl-go && \
	  echo '- ' $$(cat README.md | head -n1 | awk '{print $$3}') 'at revision' $$(git rev-parse HEAD) \
	)
	rm -rf .github/.git .github/README.md

.PHONY: pull-makes
pull-makes:
	rm -rf makes
	git clone git@github.com:kamilsk/shared.git makes
	( \
	  cd makes && \
	  git checkout makefile-go && \
	  echo '- ' $$(cat README.md | head -n1 | awk '{print $$3}') 'at revision' $$(git rev-parse HEAD) \
	)
	rm -rf makes/.git

.PHONY: pull-template
pull-template:
	rm -rf template
	git clone git@bitbucket.org:octotpl/materialkit.git template
	( \
	  cd template && \
	  git checkout 2.x && \
	  git describe --tags \
	)
	( \
	  cp -n template/Template/assets/css/material-kit.min.css docs/assets/css/ && \
	  cp -n template/Template/assets/js/bootstrap-material-design.min.js docs/assets/js/ \
	  cp -n template/Template/assets/js/material-kit.min.js docs/assets/js/ \
	)
	rm -rf template/.git
