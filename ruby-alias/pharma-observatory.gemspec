# Alias gem. rubygems.org has no redirect between gem names, so `gem install pharma-observatory`
# installs this gem, whose only job is to depend on the real one (`pharmaobservatory`, built from
# ../ruby) at the SAME version and require it. The version is read from the real gem's version.rb
# so the two cannot drift.
require_relative "../ruby/lib/pharmaobservatory/version"

Gem::Specification.new do |s|
  s.name        = "pharma-observatory"
  s.version     = PharmaObservatory::VERSION
  s.platform    = Gem::Platform::RUBY
  s.authors     = ["Scientist.com"]
  s.email       = ["support@scientist.com"]
  s.homepage    = "https://pharmaobservatory.com"
  s.summary     = "Alias for the pharmaobservatory gem"
  s.description = "This gem only depends on and requires `pharmaobservatory`, the Ruby client for the Pharma Observatory public REST API. Install `pharmaobservatory` directly."
  s.license     = "MIT"
  s.required_ruby_version = ">= 3.2"
  s.metadata    = {
    "allowed_push_host" => "https://rubygems.org",
    "homepage_uri" => "https://pharmaobservatory.com",
    "documentation_uri" => "https://pharmaobservatory.com/api-reference"
  }

  s.add_runtime_dependency "pharmaobservatory", "= #{PharmaObservatory::VERSION}"

  s.files         = ["README.md", "lib/pharma-observatory.rb"]
  s.require_paths = ["lib"]
end
