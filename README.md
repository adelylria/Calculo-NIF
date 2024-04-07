## Cálculo del dígito de control del NIF/NIE
- Para verificar el NIF de españoles residentes mayores de edad, el algoritmo de cálculo del dígito de control es el siguiente:

Se divide el número entre 23 y el resto se sustituye por una letra que se determina por inspección mediante la siguiente tabla:

| `RESTO` | 0  | 1  | 2  | 4  | 5  | 6  | 7  | 8  | 9  | 10 | 11 |
| :-----: |:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|
| `LETRA` | T  | R  | W  | A  | G  | M  | Y  | F  | P  | D  | X  | B  |

| `RESTO` | 12 | 13 | 14 | 15 | 16 | 17 | 18 | 19 | 20 | 21 | 22 |
| :-----: |:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|
| `LETRA` | N  | J  | Z  | S  | Q  | V  | H  | L  | C  | K  | E  |

- Los NIE's de extranjeros residentes en España tienen una letra (X, Y, Z), 7 números y dígito de control.

Para el cálculo del dígito de control se sustituye:
```
 X → 0
 Y → 1
 Z → 2
```
y se aplica el mismo algoritmo que para el NIF.