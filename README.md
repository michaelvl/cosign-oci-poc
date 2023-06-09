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

The GitHub workflows in this repo adds an annotation to the signature
with the SemVer of the artifact. However, it is important to remember,
that such versioning can be changed (e.g. tags can be moved), and the
annotation only guarantees, that e.g. a SemVer tag was associated with
the artifact version at build time.

### Base Images are Verified

For this specific example project:

```
cosign verify gcr.io/distroless/static-debian11:latest --certificate-oidc-issuer https://accounts.google.com  --certificate-identity keyless@distroless.iam.gserviceaccount.com
```

### Additional Documents are Signed

- SBOMS
- Attestations

## Verifying Artifacts

Container:

```
export IMAGE_DIGEST=sha256:62cfb67608e6b5665379409220c1f340e91392c4a419449085fefbff09241da2
export IMAGE_SEMVER_EXPECTED=0.5.0
cosign verify --certificate-identity-regexp https://github.com/michaelvl/cosign-oci-poc/.github/workflows/build.yaml@refs/.* --certificate-oidc-issuer https://token.actions.githubusercontent.com -a "imageRef=refs/tags/$IMAGE_SEMVER_EXPECTED"  ghcr.io/michaelvl/cosign-oci-poc@$IMAGE_DIGEST | jq .
```

Helm Chart:

```
export CHART_DIGEST=sha256:a0f685b1df374ae4d4e5d032c36fd64aada28bf1cf9f614591fef4a50c90cec6
export CHART_SEMVER_EXPECTED=0.1.0
cosign verify --certificate-identity-regexp https://github.com/michaelvl/cosign-oci-poc/.github/workflows/helm-release.yaml@refs/.* --certificate-oidc-issuer https://token.actions.githubusercontent.com -a "chartVersion=$CHART_SEMVER_EXPECTED" ghcr.io/michaelvl/cosign-oci-poc-helm@$CHART_DIGEST | jq .
```
