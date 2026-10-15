# pharmaobservatory (Ruby)

Ruby client for the [Pharma Observatory](https://pharmaobservatory.com) public REST API. Generated from the
[OpenAPI document](https://pharmaobservatory.com/api/v1/openapi.yaml); see the
[API reference](https://pharmaobservatory.com/api-reference) for what each endpoint returns. Requires Ruby 3.2 or newer.

## Install

```shell
gem install pharmaobservatory
```

or in a Gemfile: `gem "pharmaobservatory"`. (`pharma-observatory` is an alias gem that depends on this one.)

## A minimal authenticated call

The API takes a bearer token (see the `ApiBearerJwt` security scheme in the OpenAPI document for how to obtain one).

```ruby
require "pharmaobservatory"

PharmaObservatory.configure do |config|
  config.access_token = ENV.fetch("PHARMAOBSERVATORY_TOKEN")
end

puts PharmaObservatory::ApiV1BaseApi.new.api_v1_whoami_get.inspect
```

The default host is `https://pharmaobservatory.com`. The generated reference follows.

---

