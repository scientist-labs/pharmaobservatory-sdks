# Hand-written (the generator never touches it): a dataset response that carries a column the spec
# does not list still deserializes, and the known columns are intact.
require 'json'
require 'spec_helper'

describe 'dataset response with an unrecorded column' do
  it 'deserializes' do
    raw = File.read(File.expand_path('../../testdata/biomarkers_extra_column.json', __dir__))
    resp = PharmaObservatory::ApiClient.new.convert_to_type(JSON.parse(raw), 'BiomarkersResponse')
    expect(resp.rows.first.route_key).to eq('braf-v600e')
  end
end
