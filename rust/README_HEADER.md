# pharmaobservatory (Rust)

Rust client for the [Pharma Observatory](https://pharmaobservatory.com) public REST API. Generated from the
[OpenAPI document](https://pharmaobservatory.com/api/v1/openapi.yaml); see the
[API reference](https://pharmaobservatory.com/api-reference) for what each endpoint returns. The client is async
(`reqwest`), so call it from a Tokio runtime.

## Install

```shell
cargo add pharmaobservatory
```

## A minimal authenticated call

The API takes a bearer token (see the `ApiBearerJwt` security scheme in the OpenAPI document for how to obtain one).

```rust
use pharmaobservatory::apis::{api_v1_base_api, configuration::Configuration};

#[tokio::main]
async fn main() {
    let mut config = Configuration::new();
    config.bearer_access_token = Some(std::env::var("PHARMAOBSERVATORY_TOKEN").expect("PHARMAOBSERVATORY_TOKEN"));

    match api_v1_base_api::api_v1_whoami_get(&config).await {
        Ok(who) => println!("{who:?}"),
        Err(e) => eprintln!("{e}"),
    }
}
```

This example needs `tokio = { version = "1", features = ["macros", "rt-multi-thread"] }` in your dependencies.
The default host is `https://pharmaobservatory.com`. The generated reference follows.

---

