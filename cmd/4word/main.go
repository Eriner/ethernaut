package main

// A small, toy EVM compiler
// Will never be as feature-complete as Solidity or Vyper.
/*
  # Safety

  All smart contracts have a few issues to consider:
  * race conditions
  * function/contract recursion
  * oracle subversion

  The more things change, the more they stay the same.

  Many of these issues are solveable by the compiler itself, but Solidity has
  too many foot-guns to make me feel good about using it. Vyper seems to be a
  more-sane alternative, but still not to my liking.

  While remix offers some warnings about the safety of external calls,
  this is insufficient. The obligation of ensuring safety (and rejecting
  contents which can be deemed with certainty to be unsafe) is that of
  the compiler.

*/
func main() {
	// TODO: AST
}
