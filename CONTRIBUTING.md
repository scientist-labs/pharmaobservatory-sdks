# Contributing

## The SDK code is generated. Do not edit it.

Everything under `ruby/`, `go/` and `rust/` (except the small hand-maintained files listed in each `.openapi-generator-ignore`) is overwritten by the next regeneration, so an edit to a generated file is lost, and a pull request that edits one will not be merged.

If something in an SDK is wrong or missing, **open an issue** describing the call you tried and what you expected. Most problems are fixed in one of three places:

- the API's OpenAPI document (the fix lands in the API, and the daily regeneration picks it up),
- the generator configuration in `<language>/config/openapi-config.json`, or `bin/generate`,
- a supported customization, see [CUSTOMIZATION.md](CUSTOMIZATION.md).

## What a pull request may change

`bin/`, `.github/`, `<language>/config/`, `<language>/README_HEADER.md`, `<language>/.openapi-generator-ignore`, the root documents, and, once they exist, the files [CUSTOMIZATION.md](CUSTOMIZATION.md) describes.

## Regenerating locally

```shell
bin/generate            # all three, from the live document
bin/generate ruby       # one language
SPEC=path/to/openapi.yaml bin/generate go
```

Needs docker (the generator runs in a pinned Java image) or `OPENAPI_GENERATOR_JAR` pointing at the pinned jar with `java` on the path. Commit the regenerated output together with the change that caused it.

## Checks (what CI runs)

- Ruby: `cd ruby && bundle exec rspec`, then `bin/check-gems` (builds both gems, installs them into a throwaway `GEM_HOME`, requires the alias).
- Go: `cd go && go build ./... && go vet ./... && go test ./...`.
- Rust: `cd rust && cargo build && cargo test && cargo clippy`.
- `bin/check-no-codename`: no file name or file content may mention the project's internal codename. There is no exceptions list.

## Versions

Each package carries its own version in its generator config (`gemVersion` for Ruby, `packageVersion` for Go and Rust). The regeneration workflow bumps the patch version of each package whose generated output changed. `bin/bump-patch <language>` does the same by hand; minor and major bumps are made by editing the config in a pull request. The `pharma-observatory` alias gem takes its version from the Ruby gem.

## Releases

Merging to `main` publishes any package whose version is not yet on its registry, once the owner has enabled publishing. See `.github/workflows/release.yml`.
