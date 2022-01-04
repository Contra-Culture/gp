# gp - top-down parser eDSL in Golang


``gp`` does not support BNF, EBNF, WSN, or any other meta-languages for grammar specification, instead ``gp`` offers embedded DSL that helps you to create a parser as a composition of smaller parsing functions.
## Wirth syntax notation (WSN)

```
=      // production, name
.      // production terminator (full stop / period)
{x}    // repetition of x
[a]b   // optionality, matches both "ab" and "b"
a|b    // variability, matches "a" or "b"
(a|b)c // grouping, matches "ac" or "ab"
```

Any token must contain knowledge of its termination within itself. Token knows itself and warrants its finality.

Typed identity namespaces.

Symbol classes.

[]parsers


