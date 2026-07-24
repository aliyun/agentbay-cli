class Agentbay < Formula
  desc "Secure infrastructure for running AI-generated code"
  homepage "https://github.com/aliyun/agentbay-cli"
  url "https://github.com/aliyun/agentbay-cli/archive/refs/tags/v0.4.4.tar.gz"
  sha256 "906e2083fdf93c9b7ab3061aad6b7b5bba9266956684fc2548d62f4064dc61ca"
  license "Apache-2.0"
  head "https://github.com/aliyun/agentbay-cli.git", branch: "master"

  bottle do
    root_url "https://github.com/aliyun/agentbay-cli/releases/download/v0.4.4"
    sha256 cellar: :any_skip_relocation, arm64_sonoma: "61152d544d7841f55e7a4593cf968b59c68883e90e1d5262a56ab06ebb2fb6c8"
    sha256 cellar: :any_skip_relocation, arm64_ventura: "61152d544d7841f55e7a4593cf968b59c68883e90e1d5262a56ab06ebb2fb6c8"
    sha256 cellar: :any_skip_relocation, arm64_sequoia: "61152d544d7841f55e7a4593cf968b59c68883e90e1d5262a56ab06ebb2fb6c8"
    sha256 cellar: :any_skip_relocation, sonoma:       "b90666f18763f0796858b570bc0b92df6662c6715d5f58beff5f39a6cf816edf"
    sha256 cellar: :any_skip_relocation, ventura:      "b90666f18763f0796858b570bc0b92df6662c6715d5f58beff5f39a6cf816edf"
    sha256 cellar: :any_skip_relocation, x86_64_linux: "49b98c98179accc5a62baf367b158ee8c0e1e0dc68ee3975e7580973476f727d"
    sha256 cellar: :any_skip_relocation, aarch64_linux: "4a07371e0c619bfcde54dc27b32674c9778e38f9aab90372f0b7091922766144"
  end

  depends_on "go" => :build

  def install
    version = self.version
    git_commit = "0619699"
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
