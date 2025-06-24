# comparable-types-in-go
comparable types in go

### was ist der comparable type constraint

vergleichbare typen
```
==    equal
!=    not equal
<     less
<=    less or equal
>     greater
>=    greater or equal
```

### was kann man mit comparable erreichen?

vergleich der werte, sortierung etc..

### wie hat sich comparable in go über die zeit verändert

comparable als interface -> 1.18

implementierungs kriterium aufgeweicht um mehr comparable typen zu ermöglichen -> 1.20

### macht go das gut (im vergeich zu anderen sprachen)?

go musste spezial lösung gehen und extra für vergleich ausnahme des interface systems machen weil sonst nicht alle typen vergleichbar währen, aber: diese lösung kann zu abstürzen zur runtime führen
