class Agentbay < Formula
  desc "Secure infrastructure for running AI-generated code"
  homepage "https://github.com/aliyun/agentbay-cli"
  url "https://github.com/aliyun/agentbay-cli/archive/refs/tags/v0.5.0.tar.gz"
  sha256 "ac2f4ba09cd826ded9f4a905734978e365f76170357f4793918c86b0c6f1d3b5"
  license "Apache-2.0"
  head "https://github.com/aliyun/agentbay-cli.git", branch: "master"

  bottle do
    root_url "https://github.com/aliyun/agentbay-cli/releases/download/v0.5.0"
    sha256 cellar: :any_skip_relocation, arm64_sonoma: "e0ab6d7045474a63bc458ae4b940eebd0c2684127d1f77d9eef66b58dc9afb48"
    sha256 cellar: :any_skip_relocation, arm64_ventura: "e0ab6d7045474a63bc458ae4b940eebd0c2684127d1f77d9eef66b58dc9afb48"
    sha256 cellar: :any_skip_relocation, arm64_sequoia: "e0ab6d7045474a63bc458ae4b940eebd0c2684127d1f77d9eef66b58dc9afb48"
    sha256 cellar: :any_skip_relocation, sonoma:       "654fc04f447f94a39e9f2169132de63d942bafa1ec94a558fd36e8bdc6a8e2fd"
    sha256 cellar: :any_skip_relocation, ventura:      "654fc04f447f94a39e9f2169132de63d942bafa1ec94a558fd36e8bdc6a8e2fd"
    sha256 cellar: :any_skip_relocation, x86_64_linux: "81ce37ac2d926226a2fffad6da6f6a38aad0cbf269318de2901a9b90e1f04958"
    sha256 cellar: :any_skip_relocation, aarch64_linux: "4cce2fe0939342ecd01a25afefe694aed439cdda60064d24b30557aaa1354ee0"
  end

  depends_on "go" => :build

  def install
    version = self.version
    git_commit = "ef11620"
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
