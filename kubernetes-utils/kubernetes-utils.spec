%global debug_package %{nil}

%global _version          1.2.1
%global _release          1
%global _ktop_version     0.5.3
%global _kubectx_version  0.11.0
%global _kubens_version   0.11.0
%global _kubetail_version 1.6.22
%global _stern_version    1.33.1

Name         : kubernetes-utils
Version      : %{_version}
Release      : %{_release}%{?dist}
Summary      : CLI Tools for Kubernetes Deployment Management
Group        : Unspecified
License      : Apache License V2
Provides     : kubectx            = %{_kubectx_version}%{?dist}
Provides     : kubectx(%{_arch})  = %{_kubectx_version}%{?dist}
Provides     : kubens             = %{_kubens_version}%{?dist}
Provides     : kubens(%{_arch})   = %{_kubens_version}%{?dist}
Provides     : kubetail           = %{_kubetail_version}%{?dist}
Provides     : kubetail(%{_arch}) = %{_kubetail_version}%{?dist}
Provides     : stern              = %{_stern_version}%{?dist}
Provides     : stern(%{_arch})    = %{_stern_version}%{?dist}
BuildRequires: rpm-build 
Requires     : filesystem
Requires     : bash
Source0      : https://github.com/ahmetb/kubectx/releases/download/v%{_kubectx_version}/kubectx_v%{_kubectx_version}_linux_%{_arch}.tar.gz
Source1      : https://github.com/ahmetb/kubectx/releases/download/v%{_kubens_version}/kubens_v%{_kubens_version}_linux_%{_arch}.tar.gz
Source2      : https://github.com/johanhaleby/kubetail/archive/refs/tags/%{_kubetail_version}.tar.gz
Source3      : https://github.com/stern/stern/releases/download/v%{_stern_version}/stern_%{_stern_version}_linux_amd64.tar.gz
Source4      : https://github.com/vladimirvivien/ktop/releases/download/v%{_ktop_version}/ktop_v%{_ktop_version}_linux_amd64.tar.gz

%Description
CLI Tools for Kubernetes Deployment Management:
* kubectx  : tool to switch between contexts (clusters) on kubectl faster.
* kubens   : tool to switch between Kubernetes namespaces (and configure them for kubectl) easily.
* kubettail: bash script that enables you to aggregate (tail/follow) logs from multiple pods into one stream.
* stern    : tail multiple pods on Kubernetes and multiple containers within the pod.

%prep
%setup -c -T -a 0
mv LICENSE LICENSE.kubectx
%setup -D -T -a 1
mv LICENSE LICENSE.kubens
%setup -D -T -a 2
mv kubetail-%{_kubetail_version}/LICENSE kubetail-%{_kubetail_version}/LICENSE.kubetail
%setup -D -T -a 3
mv LICENSE LICENSE.stern
%setup -D -T -a 4
mv LICENSE LICENSE.ktop

%build

%install
rm -rf %{buildroot}
mkdir -p %{buildroot}%{_bindir}
install -m 755 ktop %{buildroot}%{_bindir}/
install -m 755 kubectx %{buildroot}%{_bindir}/
install -m 755 kubens  %{buildroot}%{_bindir}/
install -m 755 stern   %{buildroot}%{_bindir}/
install -m 755 kubetail-%{_kubetail_version}/kubetail  %{buildroot}%{_bindir}/kubettail
install -d %{buildroot}%{_sysconfdir}/bash_completion.d
install -m 755  kubetail-%{_kubetail_version}/completion/kubetail.bash %{buildroot}%{_sysconfdir}/bash_completion.d/


%files
%attr(0755,root,root) %{_bindir}/ktop
%attr(0755,root,root) %{_bindir}/kubectx
%attr(0755,root,root) %{_bindir}/kubens
%attr(0755,root,root) %{_bindir}/kubettail
%attr(0755,root,root) %{_bindir}/stern
%attr(0644,root,root) %{_sysconfdir}/bash_completion.d/kubetail.bash
%license LICENSE.kubectx LICENSE.kubens kubetail-%{_kubetail_version}/LICENSE.kubetail LICENSE.stern LICENSE.ktop

%changelog
* Thu Apr 02 2026 Paulo Sousa <sousa.paulo.1975@gmail.com> - 1.2.1-1
- Ktop 0.5.3
- Kubectx 0.11.0
- Kubens 0.11.0

* Fri Nov 21 2025 Paulo Sousa <sousa.paulo.1975@gmail.com> - 1.1.1-1
- Kubetail 1.6.22

* Wed Nov 12 2025 Paulo Sousa <sousa.paulo.1975@gmail.com> - 1.1.0-1
- Stern v1.33.1

* Thu Nov 06 2025 Paulo Sousa <sousa.paulo.1975@gmail.com> - 1.0.0-3
- Change kubetail name to kubettail, because conflit with kubetail-cli package

* Mon Oct 27 2025 Paulo Sousa <sousa.paulo.1975@gmail.com> - 1.0.0-2
- Stern description and provides

* Wed Oct 22 2025 Paulo Sousa <sousa.paulo.1975@gmail.com> - 1.0.0-1
- First release

