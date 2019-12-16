### Leeme

Ejemplo de implementacion de plugins en go utilizando rpc con la libreria https://github.com/hashicorp/go-plugin 

1. Compila el programa principal.

```sh
go build -o plugin-test
```

2. Ejecuta el binario y notaras que el evento original es igual al filtrado

```sh
./plugin-test
```

3. Compila cualquier plugin y guardalo en el directorio plugin con el nombre ```filterevent.plug```. En ese caso se usa el inc10

```sh
go build -o ./plugin/filterevent.plug ./plugins/inc10/main.go
```

4. Ejecuta de nuevo el binario y notaras diferencias entre el evento original y el filtrado

```sh
./plugin-test
```

5. Puedes hacer el mismo proceso con el plugin inc20 o bien puedes escribir tu propio plugin