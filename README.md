# Ethernaut

My incomplete solutions for the Ethernaut wargame.

## Contract Development

[Solidity][solidity] has a lot of foot-guns, and the compiler creators have
not developed the necessary safety (linting) mechanisms.

[Vyper][vyper] looks better (safer by default), probably the ideal happy-medium between
preventing foot-guns and letting you do what you want.

[GLOW DSL][glow] looks interesting, but isn't very mature (though, what of this is?).
One benefit it has as a DSL is that you don't have to target a specific chain,
if that matters. EVM seems to be the winner by marketshare for now, so this isn't
much of a boon now.

YOLO is where I ignore all logic and reason, and build my own toy compiler.
Nothing fancy, and honestly a light abstraction over the EVM instruction set.
I'm going to be the only person using it, so...

## Contract Management

Contract submission and wallet managed with Go. Light tooling to handle the
primary account's wallet and handle contract execution.

## Contract Reverse Engineering

[Ghidra-evm][ghidra-evm] is the best choice here, especially for those with
Ghidra/IDA experience.

I was going to create my own (Solidity) disassembler [at one point][soldis], but
even the go-ethereum SDK chokes on unknown opcodes when using the official
disassembler, lol. Probably weird Solidity compiler optimizations?

Ghidra-evm is finnicky to set up, but it's fine.

[solidity]: https://docs.soliditylang.org/en/latest/
[vyper]: https://vyper.readthedocs.io/en/latest/
[glow]: https://glow-lang.org
[ghidra-evm]: https://github.com/adelapie/ghidra-evm
