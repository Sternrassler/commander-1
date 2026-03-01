class MinCommander < Formula
  desc "Keyboard-driven terminal file manager inspired by Norton Commander"
  homepage "https://github.com/sternrassler/commander-1"
  url "https://github.com/sternrassler/commander-1/archive/refs/tags/v2.1.1.tar.gz"
  sha256 "f09bddfceb46f12beb6b526c90cca392ea84ae6dfd76eeb7430000d80e49e5b8"
  license "MIT"
  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args(ldflags: "-s -w"), "."
  end

  test do
    assert_predicate bin/"min-commander", :executable?
  end
end
