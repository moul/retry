require "language/go"

class Retry < Formula
  desc "retry: repeat shell commands"
  homepage "https://github.com/moul/retry"
  url "https://github.com/moul/retry/archive/v0.1.0.tar.gz"
  sha256 "b69819119d958f5dec197092b12c687b356e5bc12340b7ce03aeccd9c8854bb2"

  head "https://github.com/moul/retry.git"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    ENV["GOBIN"] = buildpath
    (buildpath/"src/github.com/moul/retry").install Dir["*"]

    system "go", "build", "-o", "#{bin}/retry", "-v", "github.com/moul/retry/cmd/retry/"
  end
end
