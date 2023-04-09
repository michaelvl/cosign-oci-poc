# PoC for using Cosign to sign artifacts

## Verifying Artifacts

```
cosign verify --certificate-identity https://github.com/michaelvl/cosign-oci-poc/.github/workflows/build.yaml@refs/heads/main --certificate-oidc-issuer https://token.actions.githubusercontent.com ghcr.io/michaelvl/cosign-oci-poc@sha256:fc84c7d1b142d12a484377a828d42e360a81d40eeff7d3dfaa539877fb4c74d0
```
