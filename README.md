# gp - top-down parser eDSL in Golang


``gp`` does not support BNF, EBNF, WSN, or any other meta-languages for grammar specification, instead ``gp`` offers embedded DSL that helps you to create a parser as a composition of smaller parsing functions.
## Wirth syntax notation (WSN)

```
=      // production
.      // production terminator (full stop / period)
{x}    // repetition of x
[a]b   // optionality, matches both "ab" and "b"
a|b    // variability, matches "a" or "b"
(a|b)c // grouping, matches "ac" or "ab"
```

Any fragment that is not self-terminable (has known length) is terminated by the following one.


namespace {
  
}
