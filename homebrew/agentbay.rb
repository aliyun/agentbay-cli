class Agentbay < Formula
  desc "Secure infrastructure for running AI-generated code"
  homepage "https://github.com/aliyun/agentbay-cli"
  url "https://github.com/aliyun/agentbay-cli/archive/refs/tags/v0.4.0.tar.gz"
  sha256 "4b5ffc468dbadb0b406c547f8de5a1e0f8e47349dbecf80cb8e10b5aa8d56c42"
  license "Apache-2.0"
  head "https://github.com/aliyun/agentbay-cli.git", branch: "master"

  bottle do
    root_url "https://github.com/aliyun/agentbay-cli/releases/download/v0.4.0"
    sha256 cellar: :any_skip_relocation, arm64_sonoma: "f306e616831df9614dea57b45dee899475408cd76a4c14a79830d6d29225eb9f"
    sha256 cellar: :any_skip_relocation, arm64_ventura: "f306e616831df9614dea57b45dee899475408cd76a4c14a79830d6d29225eb9f"
    sha256 cellar: :any_skip_relocation, arm64_sequoia: "f306e616831df9614dea57b45dee899475408cd76a4c14a79830d6d29225eb9f"
    sha256 cellar: :any_skip_relocation, sonoma:       "9b77a71b32e67238eb14e1edd839bd1bd17f9f34f2fe89ba9a57a04bf6890e0d"
    sha256 cellar: :any_skip_relocation, ventura:      "9b77a71b32e67238eb14e1edd839bd1bd17f9f34f2fe89ba9a57a04bf6890e0d"
    sha256 cellar: :any_skip_relocation, x86_64_linux: "4efd4b7a346d26d0cdc2d0cd212260b9a80d7c54737ba64ab0ec66bbb67a3d8c"
    sha256 cellar: :any_skip_relocation, aarch64_linux: "d6e618972a0358ff66df68145ebbbb75a88c385834fd582daa316542ab366bb4"
  end

  depends_on "go" => :build

  def install
    version = self.version
    git_commit = "3038675"
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
