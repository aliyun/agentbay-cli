class Agentbay < Formula
  desc "Secure infrastructure for running AI-generated code"
  homepage "https://github.com/aliyun/agentbay-cli"
  url "https://github.com/aliyun/agentbay-cli/archive/refs/tags/v0.3.1.tar.gz"
  sha256 "be2d235a2033f14d907083ea7d0e4254b8ae1d9ee1253a80645ae932e047fee0"
  license "Apache-2.0"
  head "https://github.com/aliyun/agentbay-cli.git", branch: "master"

  bottle do
    root_url "https://github.com/aliyun/agentbay-cli/releases/download/v0.3.1"
    sha256 cellar: :any_skip_relocation, arm64_sonoma: "0a874be48ab1bc6c2c67f65b5419223d6e01e829eea1c0a2f31d5164b249693a"
    sha256 cellar: :any_skip_relocation, arm64_ventura: "0a874be48ab1bc6c2c67f65b5419223d6e01e829eea1c0a2f31d5164b249693a"
    sha256 cellar: :any_skip_relocation, arm64_sequoia: "0a874be48ab1bc6c2c67f65b5419223d6e01e829eea1c0a2f31d5164b249693a"
    sha256 cellar: :any_skip_relocation, sonoma:       "a282c4db9fe77bd7f5b5e0ba72e6a36b83afe79be3af85cf83f1e71d5250ba8a"
    sha256 cellar: :any_skip_relocation, ventura:      "a282c4db9fe77bd7f5b5e0ba72e6a36b83afe79be3af85cf83f1e71d5250ba8a"
    sha256 cellar: :any_skip_relocation, x86_64_linux: "5f03fd04c499b1b4d1a2a601eac944495ae81088c9abaddcd8f1d5533b19c01c"
    sha256 cellar: :any_skip_relocation, aarch64_linux: "8c8f86a45f99cdc8e1c5571b74360b7a8ef6efdc1aff8b8c6a2fa566cffd3ecf"
  end

  depends_on "go" => :build

  def install
    version = self.version
    git_commit = "7b7e29e"
    build_date = Time.now.utc.strftime("%Y-%m-%dT%H:%M:%SZ")

    ENV["GOPROXY"] = "https://proxy.golang.org,https://goproxy.io,direct"
    ENV["GONOSUMCHECK"] = "*"
    ENV["GOFLAGS"] = "-mod=mod"
    ENV["GO111MODULE"] = "on"

    ldflags = %W[
      -s
      -w
      -X github.com/agentbay/agentbay-cli/cmd.Version=#{version}
      -X github.com/agentbay/agentbay-cli/cmd.GitCommit=#{git_commit}
      -X github.com/agentbay/agentbay-cli/cmd.BuildDate=#{build_date}
    ]

    system "go", "build", *std_go_args(ldflags: ldflags), "."
  end

  test do
    assert_predicate bin/"agentbay", :executable?

    version_output = shell_output("#{bin}/agentbay version 2>&1")
    assert_match version.to_s, version_output

    help_output = shell_output("#{bin}/agentbay --help")
    assert_match "agentbay", help_output
    assert_match "help", help_output
  end
end
