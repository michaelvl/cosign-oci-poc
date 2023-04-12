# PoC for using Cosign to sign artifacts

This repository contain a PoC of secure handling of container and Helm
chart artifacts.

The PoC revolves around a simple HTTP service, which is packaged into
a container and made deployable to Kubernetes through a Helm
chart. **Both the container and Helm chart artifacts are signed using
[Cosign](https://github.com/sigstore/cosign)**.

## First Principles

### Keep the Repository Pristine

The repository is used for *source*, not build artifacts.

This means, that only *developers* commit to the repository and build
artifacts are not commited back to the repository. This rules out
tools like [Helm chart
releaser](https://github.com/helm/chart-releaser-action) and pipelines
running [Helm docs](https://github.com/norwoodj/helm-docs).

### Artifacts are Signed

[Cosign](https://github.com/sigstore/cosign) will be used to implement Keyless signing.

Developers signing commits are out-of-scope. See
e.g. [git-signature-checker](https://github.com/michaelvl/git-signature-checker).

### All Artifacts are Stored in an OCI Registry

To simplify tooling and key management, a single tool will be used for
storing artifacts. This also allows us to use a single tool for
signing and verifying signatures on artifacts.

### Artifacts are Only Referenced Using Digests, Never SemVer Tags

See [Why We Should Use `latest` Tag on Container Images](https://medium.com/@michael.vittrup.larsen/why-we-should-use-latest-tag-on-container-images-fc0266877ab5)

## Verifying Artifacts

Container:

```
export IMAGE_DIGEST=sha256:fc84c7d1b142d12a484377a828d42e360a81d40eeff7d3dfaa539877fb4c74d0
cosign verify --certificate-identity https://github.com/michaelvl/cosign-oci-poc/.github/workflows/build.yaml@refs/heads/main --certificate-oidc-issuer https://token.actions.githubusercontent.com ghcr.io/michaelvl/cosign-oci-poc@$IMAGE_DIGEST
```

Helm Chart:

```
export CHART_DIGEST=sha256:8d7648e56a2d12c4cda2645f2e63b3c820f2f8789b46f887f1a1ebdcdf9ebbf7
cosign verify --certificate-identity https://github.com/michaelvl/cosign-oci-poc/.github/workflows/helm-release.yaml@refs/heads/main --certificate-oidc-issuer https://token.actions.githubusercontent.com ghcr.io/michaelvl/cosign-oci-poc-helm@$CHART_DIGEST
```
