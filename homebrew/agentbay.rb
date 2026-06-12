class Agentbay < Formula
  desc "Secure infrastructure for running AI-generated code"
  homepage "https://github.com/aliyun/agentbay-cli"
  url "https://github.com/aliyun/agentbay-cli/archive/refs/tags/v0.4.1.tar.gz"
  sha256 "eb63b68446f811893d3a6eefb107a8a0709e3b09eede3d126590572e0480587c"
  license "Apache-2.0"
  head "https://github.com/aliyun/agentbay-cli.git", branch: "master"

  bottle do
    root_url "https://github.com/aliyun/agentbay-cli/releases/download/v0.4.1"
    sha256 cellar: :any_skip_relocation, arm64_sonoma: "03dd76c0baa01ba9354b7ea1bd029c6f2a7ccb679dd00465119d545c2a4772e4"
    sha256 cellar: :any_skip_relocation, arm64_ventura: "03dd76c0baa01ba9354b7ea1bd029c6f2a7ccb679dd00465119d545c2a4772e4"
    sha256 cellar: :any_skip_relocation, arm64_sequoia: "03dd76c0baa01ba9354b7ea1bd029c6f2a7ccb679dd00465119d545c2a4772e4"
    sha256 cellar: :any_skip_relocation, sonoma:       "dc150abc21d1a640ac16030b3605c5d0352445bf39b035c662ad1c93e343bed7"
    sha256 cellar: :any_skip_relocation, ventura:      "dc150abc21d1a640ac16030b3605c5d0352445bf39b035c662ad1c93e343bed7"
    sha256 cellar: :any_skip_relocation, x86_64_linux: "2d2e68d2fe1c2ca009671bf29cee5b15ec8d542bd1bd464ebfb822487bdd40b3"
    sha256 cellar: :any_skip_relocation, aarch64_linux: "b5f0a5433e028fefa69383bf23f8d6694cab84148f060294a8971adadf8e2dd1"
  end

  depends_on "go" => :build

  def install
    version = self.version
    git_commit = "e2afa10"
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
