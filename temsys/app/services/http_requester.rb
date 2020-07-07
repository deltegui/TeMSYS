class HttpRequester

  def initialize(base_url)
    @base_url = base_url
  end

  def get(endpoint)
    res = Net::HTTP::get(self.make_uri(endpoint))
    JSON[res]
  end

  def post(endpoint, data = nil)
    data ||= ""
    res = Net::HTTP::post(self.make_uri(endpoint), data, 'Content-Type' => 'application/json')
    JSON[res]
  end

  def delete(endpoint)
    res = Net::HTTP::delete(self.make_uri(endpoint))
    JSON[res]
  end

  private

  def make_uri(endpoint)
    URI.parse "#{@base_url}#{endpoint}"
  end
end