# Pharma Observatory SDKs

Client libraries for the [Pharma Observatory](https://pharmaobservatory.com) REST API, generated from its public OpenAPI document: <https://pharmaobservatory.com/api/v1/openapi.yaml>.

| Language | Package | Directory | Install |
| --- | --- | --- | --- |
| Ruby | [`pharmaobservatory`](https://rubygems.org/gems/pharmaobservatory) (rubygems.org) | [`ruby/`](ruby/) | `gem install pharmaobservatory` |
| Ruby (alias) | [`pharma-observatory`](https://rubygems.org/gems/pharma-observatory) (rubygems.org), depends on the gem above at the same version | [`ruby-alias/`](ruby-alias/) | `gem install pharma-observatory` |
| Go | `github.com/scientist-labs/pharmaobservatory-sdks/go` | [`go/`](go/) | `go get github.com/scientist-labs/pharmaobservatory-sdks/go` |
| Rust | [`pharmaobservatory`](https://crates.io/crates/pharmaobservatory) (crates.io) | [`rust/`](rust/) | `cargo add pharmaobservatory` |

Each language directory has its own README with an install line and a minimal authenticated call.

**Status:** the code is generated and tested, but no package has been published yet (the first release waits on the registry setup described in `.github/workflows/release.yml`), so the install lines above do not work until then.

API reference: <https://pharmaobservatory.com/api-reference>

## How this repository works

Everything under `ruby/`, `go/` and `rust/` is produced by [openapi-generator](https://openapi-generator.tech) (version pinned in [`bin/generate`](bin/generate)) from the public OpenAPI document. A scheduled workflow regenerates daily and opens one pull request when the document changed. See [CONTRIBUTING.md](CONTRIBUTING.md) and, for how hand-written code will be added later, [CUSTOMIZATION.md](CUSTOMIZATION.md).

## License

MIT, see [LICENSE](LICENSE).
