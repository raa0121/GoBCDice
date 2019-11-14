# GoBCDice

[![Build Status](https://travis-ci.org/raa0121/GoBCDice.svg?branch=master)](https://travis-ci.org/raa0121/GoBCDice)
[![Build status](https://ci.appveyor.com/api/projects/status/4gl47493rao9t4b8/branch/master?svg=true)](https://ci.appveyor.com/project/raa0121/gobcdice/branch/master)
[![codecov](https://codecov.io/gh/raa0121/GoBCDice/branch/master/graph/badge.svg)](https://codecov.io/gh/raa0121/GoBCDice)

GoBCDice is a Go implementation of [BCDice](https://github.com/bcdice/BCDice),
a dice bot for tabletop RPGs supporting many Japanese game systems.
It will consist of the core dice roller (dice notation parser and evaluator) and
many game-system-specific dice bots.

## Usage

Prerequisite: [Go](https://golang.org/dl/) &ge; 1.12

Currently, only the REPL of GoBCDice can be built and run.

```bash
cd cmd/GoBCDiceREPL

# Build REPL
go build

# Run REPL
./GoBCDiceREPL
```

To modify and build the dice notation parser (pkg/core/parser/parser.go),
you also need to install [github.com/mna/pigeon](https://github.com/mna/pigeon)
first:

```bash
GO111MODULE=off go get -u github.com/mna/pigeon
```

And then build it with the following commands:

```bash
cd pkg/core/parser
make
```

## Core dice roller

The core dice roller of GoBCDice provides the common dice rolling feature.

BCDice supports the following dice notations (the detailed description can be found
at [bcdice/BCDice/docs/README.txt](https://github.com/bcdice/BCDice/tree/master/docs)
on GitHub).
The notations currently supported by GoBCDice are checked.

* [x] Sum roll (加算ロール, D): `xDn`
    * [x] With success check: `xDn>=y` etc.
* [x] Basic roll (バラバラロール, B): `nBx`
    * [x] With success check: `xBn>=y` etc.
* [x] Exploding roll (個数振り足しロール, R): `xRn>=y` etc.
* [x] Compounding roll (上方無限ロール, U): `xUn[t]`
    * [x] With success check: `xUn[t]>=y` etc.

x: number of dice, n: sides of die, y: target number, t: threshold for rerolling dice.

The optional syntaxes are as follows:

* [x] Embedding random number: `[min...max]`
* [x] Secret roll: `SxDn` etc.

The core dice roller also supports the following commands:

* [x] Calculation (arithmetic operation, C): `C(1+2-3*4/5)` etc.
* [x] Random sampling (choice): `CHOICE[A,B,C]` etc.

### Operators

Operators available in dice rolling and calculation are listed.

#### Arithmetic operators

In arithmetic operations, a numerical value is treated as an integer.

* Unary operators
    * Unary plus (noop) `+`
    * Unary minus (sign inversion) `-`
* Binary operators
    * Add `+`
    * Subtract `-`
    * Multiply `*`
    * Divide
        * Divide with rounding down `/`
        * Divide with rounding `/R`
        * Divide with rounding up `/U`

#### Comparison operators

Comparison operators are used for a success check.

* Equal to `=`
* Not equal to `<>`
* Less than `<`
* Less than or equal to `<=`
* Greater than `>`
* Greater than or equal to `>=`

## Author

[raa0121](https://twitter.com/raa0121)

Original BCDice authors are [Faceless](https://twitter.com/Faceless192x) and
[Taitai Takeru (たいたい竹流)](https://twitter.com/torgtaitai).
