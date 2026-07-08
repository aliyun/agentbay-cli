class Agentbay < Formula
  desc "Secure infrastructure for running AI-generated code"
  homepage "https://github.com/aliyun/agentbay-cli"
  url "https://github.com/aliyun/agentbay-cli/archive/refs/tags/v0.4.2.tar.gz"
  sha256 "ed1f7b750f5342e631fc1aefad83c25784a23e9c6f6268b688e4b0a9ce40b29a"
  license "Apache-2.0"
  head "https://github.com/aliyun/agentbay-cli.git", branch: "master"

  bottle do
    root_url "https://github.com/aliyun/agentbay-cli/releases/download/v0.4.2"
    sha256 cellar: :any_skip_relocation, arm64_sonoma: "0ebd951e065c8d02233ead1b245b7af0a64d899ef3f1f0fa9594cde11f7991f2"
    sha256 cellar: :any_skip_relocation, arm64_ventura: "0ebd951e065c8d02233ead1b245b7af0a64d899ef3f1f0fa9594cde11f7991f2"
    sha256 cellar: :any_skip_relocation, arm64_sequoia: "0ebd951e065c8d02233ead1b245b7af0a64d899ef3f1f0fa9594cde11f7991f2"
    sha256 cellar: :any_skip_relocation, sonoma:       "24b85da6bab5b2abe96dbaaa950e6f1bf1cfb825a09a24b6b925135821f52a82"
    sha256 cellar: :any_skip_relocation, ventura:      "24b85da6bab5b2abe96dbaaa950e6f1bf1cfb825a09a24b6b925135821f52a82"
    sha256 cellar: :any_skip_relocation, x86_64_linux: "3dfe65686d880781d0fc4345da34e64b4cfca85523fa61dcbbc17e1fe351e65d"
    sha256 cellar: :any_skip_relocation, aarch64_linux: "b648fcd67bc2cf9e14e20a6d7a54406f4aaebf365a5e2eca3c48ef8a7d62abe2"
  end

  depends_on "go" => :build

  def install
    version = self.version
    git_commit = "9237300"
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
