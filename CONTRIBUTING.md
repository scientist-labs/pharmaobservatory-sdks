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

## Owner setup

- **Regeneration PR.** The daily workflow opens its pull request with the default token, which needs "Allow GitHub Actions to create and approve pull requests" turned on: first for the `scientist-labs` organization (Settings, Actions, General), then for this repository (the repository setting is refused while the organization one is off). Until then the job pushes the branch `automation/regenerate` and fails with a message naming the setting; the PR can be opened from that branch by hand. When the in-run CI passes the workflow squash-merges the PR itself, deletes the branch and dispatches the Release workflow; if CI fails the PR stays open.
- **Publishing gates.** Set the repository variables `PUBLISH_RUBYGEMS`, `PUBLISH_GO` and `PUBLISH_CRATES` to `true` (Settings, Secrets and variables, Actions, Variables) to enable each registry separately. Gems are pushed with the org secret `RUBYGEMS_API_KEY`; Go publishes by tag.
- **crates.io.** The crates job publishes with the org Actions secret `CARGO_REGISTRY_TOKEN`, so no manual first publish is needed. Set `PUBLISH_CRATES=true` to enable it.

## Checks (what CI runs)

- Ruby: `cd ruby && bundle exec rspec`, then `bin/check-gems` (builds both gems, installs them into a throwaway `GEM_HOME`, requires the alias).
- Go: `cd go && go build ./... && go vet ./... && go test ./...`.
- Rust: `cd rust && cargo build && cargo test && cargo clippy`.
- `bin/check-no-codename`: no file name or file content may mention the project's internal codename. There is no exceptions list.

## Versions

Each package carries its own version in its generator config (`gemVersion` for Ruby, `packageVersion` for Go and Rust). The regeneration workflow bumps the patch version of each package whose generated output changed. `bin/bump-patch <language>` does the same by hand; minor and major bumps are made by editing the config in a pull request. The `pharma-observatory` alias gem takes its version from the Ruby gem.

## Releases

Merging to `main` publishes any package whose version is newer than the latest on its registry, for each registry whose `PUBLISH_*` variable is `true`. See `.github/workflows/release.yml`.
