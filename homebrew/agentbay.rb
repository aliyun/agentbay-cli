class Agentbay < Formula
  desc "Secure infrastructure for running AI-generated code"
  homepage "https://github.com/aliyun/agentbay-cli"
  url "https://github.com/aliyun/agentbay-cli/archive/refs/tags/v0.4.3.tar.gz"
  sha256 "d0e6c270d8aa4673249bf282d44beb7245baf06bbfd82be1b5c1e94eab1b8e5d"
  license "Apache-2.0"
  head "https://github.com/aliyun/agentbay-cli.git", branch: "master"

  bottle do
    root_url "https://github.com/aliyun/agentbay-cli/releases/download/v0.4.3"
    sha256 cellar: :any_skip_relocation, arm64_sonoma: "5d6ef31fa716522c62717e600d3c5908520db22ef256bf238e35370d195c688e"
    sha256 cellar: :any_skip_relocation, arm64_ventura: "5d6ef31fa716522c62717e600d3c5908520db22ef256bf238e35370d195c688e"
    sha256 cellar: :any_skip_relocation, arm64_sequoia: "5d6ef31fa716522c62717e600d3c5908520db22ef256bf238e35370d195c688e"
    sha256 cellar: :any_skip_relocation, sonoma:       "6d6abf338bacfd7e8a2f9ec4a5581ce4221d9d6232c450811789412f0a2087dd"
    sha256 cellar: :any_skip_relocation, ventura:      "6d6abf338bacfd7e8a2f9ec4a5581ce4221d9d6232c450811789412f0a2087dd"
    sha256 cellar: :any_skip_relocation, x86_64_linux: "9da567ccb026c38e87b4752fd6421505d31976afc4f8656831116b7fb2a7ac7a"
    sha256 cellar: :any_skip_relocation, aarch64_linux: "a4724999c10999ea3e59221a10b461f2238aa4a2e24fdf109f6be9828520ce9e"
  end

  depends_on "go" => :build

  def install
    version = self.version
    git_commit = "d94c649"
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
