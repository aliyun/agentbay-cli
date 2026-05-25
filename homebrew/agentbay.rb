class Agentbay < Formula
  desc "Secure infrastructure for running AI-generated code"
  homepage "https://github.com/aliyun/agentbay-cli"
  url "https://github.com/aliyun/agentbay-cli/archive/refs/tags/v0.3.2.tar.gz"
  sha256 "5a6e217873a9713c41d9951e3a3acacbe821c5ff95f03a05bd8a26e226a24452"
  license "Apache-2.0"
  head "https://github.com/aliyun/agentbay-cli.git", branch: "master"

  bottle do
    root_url "https://github.com/aliyun/agentbay-cli/releases/download/v0.3.2"
    sha256 cellar: :any_skip_relocation, arm64_sonoma: "6d7d4af4301cbb39ba96f88351e61b0c2dae6a8e11bcdeb55a0f02a850faed81"
    sha256 cellar: :any_skip_relocation, arm64_ventura: "6d7d4af4301cbb39ba96f88351e61b0c2dae6a8e11bcdeb55a0f02a850faed81"
    sha256 cellar: :any_skip_relocation, arm64_sequoia: "6d7d4af4301cbb39ba96f88351e61b0c2dae6a8e11bcdeb55a0f02a850faed81"
    sha256 cellar: :any_skip_relocation, sonoma:       "486c2af9a1a9de64f5768d95154b8114c10753f6a3efc98468e94840eab93ac4"
    sha256 cellar: :any_skip_relocation, ventura:      "486c2af9a1a9de64f5768d95154b8114c10753f6a3efc98468e94840eab93ac4"
    sha256 cellar: :any_skip_relocation, x86_64_linux: "d9783079122982a34e5847de7a49573319c5400bac71b4d08adb37db45994a3e"
    sha256 cellar: :any_skip_relocation, aarch64_linux: "f5936d2d405c74850a43ecb4756053fe56ae115c778c0e6584c545caf4968762"
  end

  depends_on "go" => :build

  def install
    version = self.version
    git_commit = "8c44249"
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
