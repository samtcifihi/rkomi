# README

rkomi is a utility for calculating appropriate reverse komi and/or handicap for Go (aka: weiqi/baduk).
Results are displayed in a single line and include the players' strengths, the color they play (adjusted if necessary from that given if handicap stones are used), the number of handicap stones used (if any), and the total value of komi, including both base komi and reverse komi, to be added or subtracted from White's final score on the board.

If the minimums and maximums allow it, the utility will always attempt to create as nearly as possible a full handicap for the skill gap based on the base komi and stone width provided, or there by default.
If it can do this spot on with only reverse komi, it will.
Failing that, if it can do this spot on with the minimum number of handicap stones which still allows the reverse komi to fully make up the gap, it will.
Failing that, if it can do this with the minimum number of handicap stones which still allows the reverse komi to fully make up at least as much of the gap as with any other strictly higher number of handicap stones, it will.
Failing that, it will choose an number of handicap stones which allow the reverse komi to make the gap as small as possible, with fewer handicap stones winning ties.

Uses stdin and stdout to support piping and passing arguments.
Input can be piped in from another command, or given as arguments when calling the utility.
Output can be piped into another command, or just outputed to the terminal.

## Usage and Optional Flags

TODO: add single letter options with the `flag` package if possible, otherwise replace all flags with single letters based on their curent first character.

Accepts either kyu-dan (mantissa optional) or OGS glicko-2 values (or a combination) for the strengths of two players, White and Black (default order, but flag `order` (default `order=white <- {white, black, uwate, shitate}`) may be passed to change this, along with the order players are displayed in the output).
Optional flags (all except strengths for both players) may be sent before, after, or between the player strengths.

`baseKomi`, the komi which would be used in an even game, is set by default to `7.0`, but may be changed with the flag `komi=<float64>`.
`stoneWidth`, the number of points each stone is worth, is set by default to `14`, but may be changed with the flag `stonewidth=<uint>` (`int` internally).
`preferHandicap`, which sets attempting to make up the skill gap with handicap stones before adding in reverse komi, is set by default to `false`, but may be changed with the flag `preferhandi=<bool>`.

`minRkomi`, the minimum absolute value of reverse komi before baseKomi is factored in which may be returned, is set by default to `0`.
`maxRkomi`, the maximum absolute value of reverse komi before baseKomi is factored in which may be returned, is set by default to `354`.
`rkomi=<int>` changes both `minRkomi` and `maxRkomi` to the specified value.
`rkomi=<int>,<int>` changes both `minRkomi` and `maxRkomi` to the first and second values, respectively.
A negative or blank (negative internally) value removes the restriction.

`minHandicap`, the minimum number of handicap stones which will be used, is set by default to `0`.
`maxHandicap`, the maximum number of handicap stones which will be used, is set by default to `0`.
`handi=<int>` changes both `minHandicap` and `maxHandicap` to the specified value.
`handi=<int>,<int>` changes `minHandicap` and `maxHandicap` to the first and second values, respectively.
A negative or blank (negative internally) value removes the restriction.
