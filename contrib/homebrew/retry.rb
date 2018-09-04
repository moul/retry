require "language/go"

class Retry < Formula
  desc "retry: repeat shell commands"
  homepage "https://github.com/moul/retry"
  url "https://github.com/moul/retry/archive/v0.4.0.tar.gz"
  sha256 "1feb8586da7e50227dc13badee37db2e8b981592f00f28b69cf3b8260a384168"

  head "https://github.com/moul/retry.git"

  depends_on "go" => :build

  def install
    ENV["GOPATH"] = buildpath
    ENV["GOBIN"] = buildpath
    (buildpath/"src/moul.io/retry").install Dir["*"]

    system "go", "build", "-o", "#{bin}/retry", "-v", "moul.io/retry"
  end
end
