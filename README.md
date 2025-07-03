# Comparable Type Constraint in Go

## Was ist der `comparable` type constraint
>	[The equality operators == and != apply to operands of comparable types.](https://go.dev/ref/spec#Comparison_operators)

`comparable` hat zwei Dimensionen:
1.  umfasst typen die `comparable` implementieren und dadurch vergleichbar sind,
2.  umfasst typen die `comparable` nicht implementieren und trotzdem vergleichbar sein können.

### Erster Punkt
Einige sind relativ einfach, hier muss nichts weiter beachtet werden:
-  Numerische Werte:  vergleichbar wie gewohnt aus anderen Sprachen,
-  `bool`:              Gleichheit bei gleichem Wert,
-  `string` :           Vergleich der Bytes,
-  `pointer`:           Vergleich der Adresse,
-  `channel`:           Gleich, wenn der selbe Aufruf die beiden verglichenen channel erstellt hat.

### Zweiter Punkt
Folgende Typen sind nicht immer vergleichbar, nur wenn die zugehörige Bedingung erfüllt ist:
-  `array`:             vergleichbar, wenn Element-Typ vergleichbar ist,
-  `structs`:           vergleichbar, wenn alle Attribut-Typen vergleichbar sind,
-  Generische Typen:    vergleichbar, wenn alle enthaltenen Typen vergleichbar sind.

### Was ist nicht `comparable`?
`maps`, `slices` und Funktionen sind nicht vergleichbar im Sinn von Go `comparable`, auch wenn die zugrunde liegenden Typen `comparable` sind. 

## Vergleich in C++
In C++ kann man von Haus aus Objekte des Typs `std::map` vergleichen, dabei werden die Elemente auf Gleichheit verglichen, nach Regeln der zugrundeliegenden Typen.

Das Äquivalent eines `slices` aus Go ist der `std::vector`, diesen kann man ebenfalls über den bereits existierenden `operator==` vergleichen welcher die Elemente der Reihe nach vergleicht wie bei der map.

Ein Objekt des Typs `std::function` kann man zwar vergleichen, aber nur mit einem `nullptr`, will man den Speicherort der Funktion vergleichen gibt es einen Trick indem man die Adressen aus den Objekten holt und diese vergleicht:
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
~~~
struct MySet : std::vector<int> {
  bool operator==(MySet const& other) const {
    return std::is_permutation(begin(), end(), other.begin());
  }
};
~~~

### Vergleich in Java
Die `Object` Klasse, von der jede andere Klasse erbt, gibt eine `equals()`-Methode vor, welche per default `==` benutzt, also die Referenz vergleicht.

Java kennt verschiedene Arten von maps, z.B. `HashMap`, `TreeMap`, alle maps haben aber die `equals()`-Methode überschrieben und vergleichen die maps auf Inhalt.

Für dynamische Arrays gibt es in Java die `ArrayList`, die gleich wie die maps die `equals()`-Methode angepasst hat um den Inhalt zu prüfen.

Bei Funktionen macht Java keinen Unterschied zwischen `==` und `.equals()`, allerdings erstellt jeder Aufruf ein neues Objekt und somit eine neue Referenz sodass der Vergleich wieder etwas schwieriger ist.
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

## Kann ich nicht `comparable` Typen in Go trotzdem vergleichen?
Für den Schlüssel in einer map in Go reicht constraint `any` nicht aus weil explizit beim Zugriff der Operator `==` verwendet wird, und in Go kann man den `==`-Operator nicht überschreiben, deshalb ist es nie möglich maps, slices oder funktionen als Schlüssel zu benutzen. structs mit maps, slices oder funktionen als Attributen, Arrays mit maps, slices oder funktionen als Elemente oder generische Typen welche maps, slices oder funktionen enthalten können zwar erstellt werden, aber werden diese nicht `comparable` Typen als Schlüssel festgestellt kommt es zur Laufzeit zum panic und das Programm kann abstürzen wenn dieser nicht aufgefangen wird.

### Eigene Funktion schreiben
Natürlich kann man auch eigene Hilfsfunktionen sich ausdenken um maps und slices zu vergleichen, wenn eine eigene komplexe Datenstruktur verglichen werden soll kann es sogar am performantesten sein eine Methode zu schreiben welche auf die Daten angepasst ist, aber es gibt auch Methoden bzw. Funktionen aus Bibliotheken welcher man sich bedienen kann:
-    `map` und `slice` bieten eine eigene `equal`-Methode an,
~~~
import "golang.org/x/exp/maps"
       "golang.org/x/exp/slices"

func main() {
	map1 := map[string]string{
		"Anna": "anna@gmail.com",
		"Bob":  "bob@gmail.com",
	}
	map2 := map[string]string{
		"Anna": "anna@gmail.com",
		"Bob":  "bob2@gmail.com",
	}
	maps.Equal(map1, map2)

	slice1 := []int{1, 2, 3}
	slice2 := []int{1, 2, 3}
	slices.Equal(slice1, slice2)
}
~~~
-    das `cmp`-Package bietet die Funktion [cmp.Equal](https://pkg.go.dev/github.com/google/go-cmp/cmp#Equal) welche interface{} annimmt und anhand des Inhalts vergleicht, dabei müssen allerdings ein paar Einschränkungen beachtet werden, so müssen z.B. alle Felder eines `structs` exportiert werden aber man gewinnt auch mit ein paar fortschrittliche Anpassungs Optionen,
-    aus dem `reflect`-Package gibt es die [reflect.DeepEqual](https://pkg.go.dev/reflect#DeepEqual)-Methode welche bei den meisten Daten korrekte Ergebnisse liefert, aber auch inkonsistent sein kann und unerwartete Ergebnisse erzeugt.

## Bewertung
Grundsätzlich ist es gut, dass man auch eigene `struct`s (mit Einschränkung) vergleichen kann ohne dabei eine eigene Methode schreiben zu müssen. Und solange man keine ausgefallenen oder komplexe Datenstrukturen hat kann man sich auch auf Vergleichsfunktionen aus Hilfsbibliotheken stützen.

Die einzige wirkliche Einschränkung ist die der `map`-Schlüssel. Auch wenn hier mit einem `string` oder Hash der das Objekt representieren soll oder `pointer` gearbeitet werden kann, ist dies nicht immer möglich oder gewünscht. 

Andererseits beugt Go so einige Bugs vor die entstehen können wenn nicht-vergleichbare Typen als Key erlaubt wären, dies führt in anderen Sprachen manchmal zu Bugs die nur schwer erkennbar sind. Go geht einen Mittelweg wenn es zur Laufzeit-panic kommt, einige Sprachen sagen deutlich dass manche Operationen zu undefiniertem Verhalten führen und geben dem Programmierer die Verantwortung das zu beachten, aber andere Sprachen lassen erst gar nicht zu, dass kritischer Code kompiliert wird. Go verhindert undefiniertes Verhalten in diesem Bereich, aber keine Abstürze. Ich finde man könnte hier konsequenter sein und versuchen das früher zu verhindern sodass dem Benutzer das Programm nicht um die Ohren fliegt.
