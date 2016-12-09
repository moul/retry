require "language/go"

class Retry < Formula
  desc "retry: repeat shell commands"
  homepage "https://github.com/moul/retry"
  url "https://github.com/moul/retry/archive/v0.2.0.tar.gz"
  sha256 "1e69a41a99a64c3b49fc5588c733bd90a404c9a7132717a89f7569dee0512c9f"

  head "https://github.com/moul/retry.git"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    ENV["GOBIN"] = buildpath
    (buildpath/"src/github.com/moul/retry").install Dir["*"]

    system "go", "build", "-o", "#{bin}/retry", "-v", "github.com/moul/retry/cmd/retry/"
  end
end
