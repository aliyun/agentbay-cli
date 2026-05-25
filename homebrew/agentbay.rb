class Agentbay < Formula
  desc "Secure infrastructure for running AI-generated code"
  homepage "https://github.com/aliyun/agentbay-cli"
  url "https://github.com/aliyun/agentbay-cli/archive/refs/tags/v0.3.0.tar.gz"
  sha256 "b1717ee4b3bcd13a4b1c9cbf8f8e004269b92f0fdf6940cf2095b8f89f663dfa"
  license "Apache-2.0"
  head "https://github.com/aliyun/agentbay-cli.git", branch: "master"

  bottle do
    root_url "https://github.com/aliyun/agentbay-cli/releases/download/v#{version}"
    sha256 cellar: :any_skip_relocation, arm64_sonoma: "PLACEHOLDER_ARM64_SONOMA"
    sha256 cellar: :any_skip_relocation, sonoma:       "PLACEHOLDER_SONOMA"
    sha256 cellar: :any_skip_relocation, x86_64_linux: "PLACEHOLDER_X86_64_LINUX"
  end

  depends_on "go" => :build

  def install
    version = self.version
    git_commit = "16dda9f"
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
