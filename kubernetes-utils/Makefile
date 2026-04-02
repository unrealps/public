_ktop_version     := 0.5.3
_kubectx_version  := 0.11.0
_kubens_version   := 0.11.0
_kubetail_version := 1.6.22
_stern_version    := 1.33.1

all: prepare download build

prepare:
	mkdir -p ~/rpmbuild/SOURCES || /bin/true

download: download_ktop download_kubectx download_kubens download_kubetail download_stern

download_ktop:
	wget -c -P ~/rpmbuild/SOURCES https://github.com/vladimirvivien/ktop/releases/download/v$(_ktop_version)/ktop_v$(_ktop_version)_linux_amd64.tar.gz

download_kubectx:
	wget -c -P ~/rpmbuild/SOURCES https://github.com/ahmetb/kubectx/releases/download/v$(_kubectx_version)/kubectx_v$(_kubectx_version)_linux_x86_64.tar.gz

download_kubens:
	wget -c -P  ~/rpmbuild/SOURCES https://github.com/ahmetb/kubectx/releases/download/v$(_kubens_version)/kubens_v$(_kubens_version)_linux_x86_64.tar.gz

download_kubetail:
	wget -c -P  ~/rpmbuild/SOURCES https://github.com/johanhaleby/kubetail/archive/refs/tags/$(_kubetail_version).tar.gz

download_stern:
	wget -c -P  ~/rpmbuild/SOURCES https://github.com/stern/stern/releases/download/v$(_stern_version)/stern_$(_stern_version)_linux_amd64.tar.gz

build:
	rpmbuild -ba \
		 kubernetes-utils.spec