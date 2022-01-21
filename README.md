# README

rkomi is a utility for calculating appropriate reverse komi and/or handicap for Go (aka: weiqi/baduk).
Results are displayed in a single line and include the players' strengths, the color they play (adjusted if necessary from that given if handicap stones are used), the number of handicap stones used (if any), and the total value of komi, including both base komi and reverse komi, to be added or subtracted from White's final score on the board.

TODO If the minimums and maximums allow it, the utility will always attempt to create as nearly as possible a full handicap for the skill gap based on the base komi and stone width provided, or there by default.
TODO If it can do this spot on with only reverse komi, it will.
TODO Failing that, if it can do this spot on with the minimum number of handicap stones which still allows the reverse komi to fully make up the gap, it will.
TODO Failing that, if it can do this with the minimum number of handicap stones which still allows the reverse komi to fully make up at least as much of the gap as with any other strictly higher number of handicap stones, it will.
TODO Failing that, it will choose an number of handicap stones which allow the reverse komi to make the gap as small as possible, with fewer handicap stones winning ties.

TODO Uses stdin and stdout to support piping and passing arguments.
TODO Input can be piped in from another command, or given as arguments when calling the utility.
TODO Output can be piped into another command, or just outputed to the terminal.

## Usage and Optional Flags

TODO: add single letter options with the `flag` package if possible, otherwise replace all flags with single letters based on their curent first character.

Accepts either kyu-dan (integer values) or OGS glicko-2 values (or a combination) for the strengths of two players, White and Black.
TODO Optional flags (all except strengths for both players) may be sent before, after, or between the player strengths.

TODO `baseKomi`, the komi which would be used in an even game, is set by default to `7.0`, but may be changed with the flag `komi=<float64>`.
TODO `stoneWidth`, the number of points each stone is worth, is set by default to `14`, but may be changed with the flag `stonewidth=<uint>` (`int` internally).
TODO `preferHandicap`, which sets attempting to make up the skill gap with handicap stones before adding in reverse komi, is set by default to `false`, but may be changed with the flag `preferhandi=<bool>`.

TODO `order`, the order in which players are displayed in the output, is set by default to `white`.
TODO `random`, set if the colors are randomized in 0 handicap stone games, is set by default to `true`.
TODO `order={white, black, uwate, shitate, none}` changes `order` to the specified value.
TODO `random=<bool>` changes `random` to the specified value.
TODO For `order={white, black}`, the first player passed as an argument is the specified color if possible if `random=false`.
TODO For `order={uwate, shitate}`, if `random=false` and if possible, if the first player passed as an argument is the specified strength, uwate plays white, otherwise uwate plays black.
TODO For `order=none`, `random` has no effect, the players are displayed in the order they were passed as arguments, and their color is always randomized if possible.

TODO `minRkomi`, the minimum absolute value of reverse komi before baseKomi is factored in which may be returned, is set by default to `0`.
TODO `maxRkomi`, the maximum absolute value of reverse komi before baseKomi is factored in which may be returned, is set by default to `354`.
TODO `rkomi=<int>` changes both `minRkomi` and `maxRkomi` to the specified value.
TODO `rkomi=<int>,<int>` changes both `minRkomi` and `maxRkomi` to the first and second values, respectively.
TODO A negative or blank (negative internally) value removes the restriction.

TODO `minHandicap`, the minimum number of handicap stones which will be used, is set by default to `0`.
TODO `maxHandicap`, the maximum number of handicap stones which will be used, is set by default to `0`.
TODO `handi=<int>` changes both `minHandicap` and `maxHandicap` to the specified value.
TODO `handi=<int>,<int>` changes `minHandicap` and `maxHandicap` to the first and second values, respectively.
TODO A negative or blank (negative internally) value removes the restriction.
