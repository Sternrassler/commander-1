class MinCommander < Formula
  desc "Keyboard-driven terminal file manager inspired by Norton Commander"
  homepage "https://github.com/sternrassler/commander-1"
  url "https://github.com/sternrassler/commander-1/archive/refs/tags/v2.1.2.tar.gz"
  sha256 "31b001586e8182ac2b899169bdb5405c41e1b5ec42b829ac1d79ac99d39312a7"
  license "MIT"
  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w"), "."
  end

  test do
    assert_predicate bin/"min-commander", :executable?
  end
end
