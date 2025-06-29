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
-  Typ Parameter:       vergleichbar wenn alle entaltenen Typen vergleichbar sind.

### Was ist nicht `comparable`?
`maps`, `slices` und Funktionen sind nicht vergleichbar im Sinn von Go `comparable`, auch wenn die zugrunde liegenden Typen `comparable` sind. 

## Vergleich in C++
In C++ kann man von Haus aus Objekte des Typs `std::map` vergleichen, dabei werden die Elemente auf Gleichheit verglichen, nach Regeln der zugrundeliegenden Typen.

Das Äquivalent eines `slices` aus Go ist der `std::vector`, diesen kann man ebenfalls über den bereits existierenden `operator==` vergleichen welcher die Elemente der Reihe nach vergleicht wie bei der Map.

Ein Objekt des Typs `std::function` kann man zwar vergleichen, aber nur mit einem `nullptr`, will man den Speicherort der Funktion vergleichen gibt es einen Trick indem man die Adresse aus den Object holt und diese vergleicht:
~~~
template<typename T, typename... U>
size_t getAddress(std::function<T(U...)> f) {
    typedef T(fnType)(U...);
    fnType ** fnPointer = f.template target<fnType*>();
    return (size_t) *fnPointer;
}

if (getAddress(f) == getAddress(g)) {...}
~~~
Um eigene Datenstrukturen zu vergleichen kann man in C++ auch einfach eine `operator==` definieren und Implementieren, so kann man z.B. von std::vector erben und die Vergleichsmethode mit einer überschreiben, welche die Reihenfolge der Elemente ignoriert.

### Vergleich in Java
Die `Object` Klasse, von der jede andere Klasse erbt, gibt eine `equals()`-Methode vor, welche per default `==` benutzt, also die Referenz vergleicht.

Java kennt verschiedene Arten von Maps, z.B. `HashMap`, `TreeMap`, alle Maps haben aber die `equals()`-Methode überschrieben und vergleichen die Maps auf Inhalt.

Für dynmische Arrays gibt es in Java die `ArrayList`, die gleich wie die Maps die `equals()`-Methode angepasst hat um den Inhalt zu prüfen.

Bei Funktionen macht Java keinen Unterschied zwischen `==` und `.equals()`, allerdings erstellt jeder Aufruf ein neues Objekt und somit eine neue Referenz sodas der Vergleich wieder etwas schwieriger ist.
~~~
import java.util.function.Function;

public class Demo {
    public static void main(String[] args) {
        Function<String, Integer> f1 = Integer::parseInt;
        Function<String, Integer> f2 = Integer::parseInt;
        Function<String, Integer> f3 = f1;
        
        System.out.println(f1 == f2);      // false
        System.out.println(f1.equals(f2)); // false
        
        System.out.println(f1 == f3);      // true
        System.out.println(f1.equals(f3)); // true
    }
}
~~~
Ähnlich wie bei eigenen Klassen in C++ kann man bei seinen Klassen in Java die von Object geerbte Methode `equals()` überschreiben um den gewünschten Vergleich zur Auswahl zu haben.

## Kann man nicht `comparable` in Go umgehen?
Für den Schlüssel in einer Map in Go reicht constraint `any` nicht aus weil explizit beim Zugriff der Operator `==` verwendet wird, und in Go kann man denn `==` Operator nicht überschreiben, deshalb ist es nicht möglich Maps, slices oder funktionen, oder structs mit maps, slices oder funktionen als Attributen, Arrays mit maps, slices oder funktionen als Elemente oder Typ Parameter welche maps, slices oder funktionen enthalten als Schlüssel zu benutzen. Sollte dies trotzdem versucht und geschrieben werden kann der Code zwar kompiliert werden, aber wird ein nicht `comparable` Typ als Schlüssel festgestellt kommt es zur Laufzeit zum panic und das Programm kann abstürzen wenn dieser nicht aufgefangen wird.

### Eigene Methode schreiben
Natürlich kann man auch eigene Methoden sich ausdenken um maps und slices zu vergleichen, wenn eine eigene komplexe Datenstruktur verglichen werden soll kann es sogar am performantesten sein eine Methode zu schreiben welche auf die Daten angepasst ist, aber es gibt auch 
