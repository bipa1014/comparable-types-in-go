# comparable-types-in-go

## Was ist der `comparable` type constraint
`comparable` hat zwei Dimensionen:
1.  umfasst typen die `comparable` implementieren und dadurch vergleichbar sind,
2.  umfasst typen die `comparable` nicht implementieren und trotzdem vergleichbar sein können.

### Erster Punkt
Einige sind relativ einfach, hier muss nichts weiter beachtet werden:
-  Numerische Werte:  vergleichbar wie gewoht aus anderen Sprachen,
-  `bool`:              Gleichheit bei gleichem Wert,
-  `string` :           Vergleich der Bytes,
-  `pointer`:           Vergleich der Adresse,
-  `channel`:           Vergleich auf Basis des Aufrufs welcher den channel erstellt hat.

### Zweiter Punkt
Folgende Typen sind nicht immer vergleichbar, nur wenn die zugehöhrige Bedingung erfüllt ist:
-  `array`:             vergleichbar wenn alle Elemente des arrays vergleichbar sind,
-  `structs`:           vergleichbar wenn alle Attribute vergleichbar sind,
-  Typ Parameter:     vergleichbar wenn alle entaltenen Typen vergleichbar sind.

### Was ist nicht `comparable`?
`maps`, `slices` und Funktionen sind nie vergleichbar im Sinn von Go `comparable`, auch wenn die zugrunde liegenden Typen `comparable` sind. 

## 

## was kann man mit comparable erreichen?

vergleich der werte, sortierung etc..

## wie hat sich comparable in go über die zeit verändert

comparable als interface -> 1.18

implementierungs kriterium aufgeweicht um mehr comparable typen zu ermöglichen -> 1.20

## macht go das gut (im vergeich zu anderen sprachen)?

go musste spezial lösung gehen und extra für vergleich ausnahme des interface systems machen weil sonst nicht alle typen vergleichbar währen, aber: diese lösung kann zu abstürzen zur runtime führen
