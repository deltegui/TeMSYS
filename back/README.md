# Como ejecutar el proyecto
Primero instala las dependencias y software necesario:

* Necesitas tener instalado Golang y bien configurado
* Necesitas tener MySQL 8 o MariaDB 10

Hecho esto tienes que copiar el fichero 'dev.example.json' a 'dev.json' y configurarlo como necesites. Si estas en un entorno de produccion puede que quieras que el servidor use HTTPS. para ello debes habilitarlo en la configuración y generar las claves usando el script create_keys.sh que se encuentra en la raiz del proyecto. También es recomendable en el caso de estar en producción, en usar variables de entorno o pasar los parámetros por terminal en vez de usar 'dev.json', que esta pensado más bien para desarrollo.

Hecho esto debes configurar la base de datos. Para ello carga el fichero sql llamado 'db.sql' que se encuentra en la raiz del proyecto en MySQL.

Ahora puedes crear el usuario administrador ejecutando el programa que se encuentra en cmd/admincreate/

Y ya puedes ejecutar el sistema, que se encuentra en cmd/temsys/

# Recomendaciones para desarrollar
Suele venir bien tener hot reloading. Para ello se recomienda usar reflex (https://github.com/cespare/reflex). Por ejemplo, se podría hacer:

```
reflex -r '\.go$' -s -- go run ./cmd/temsys/main.go
```

Con esto, cualquier cambio en un fichero de golang reiniciará el servidor.

# Estructura del proyecto

El proyecto sigue la Clean Architecture de Uncle Bob.

En la raiz del proyecto se encuentran algunos ficheros de configuración, además del código de las capas de aplicación y dominio. Este código de dominio esta dividido en ficheros, uno por cada "porción" de dominio (por ejemplo, todos los objetos de dominio y casos de uso relacionados con usuarios esta en el fichero users.go).

Todas las abstracciones definidas en estos ficheros se encuentran implementados por el código que hay dentro de las carpetas:

```
/api                        Todo lo relacionado con la API (middlewares, CORS, handlers...)
/builders                   Solo contiene un builder para la clase Sensor.
/cmd                        Comandos. Contiene el codigo fuente de los programas finales.
/configuration              Código que maneja la configuración del sistema.
/connectors                 Implementaciones de diferentes conectores de sensores.
/cronscheluder              Scheluder encargado de recoger información de los sensores.
/hash                       Implementaciones de hashers para contraseñas.
/mysql                      Implementaciones de repositorios en MySQL.
/testu                      Utilidades de testing.
/token                      Implementaciones de tokenes (por ejemplo JWT).
/validator                  Implementaciones de validadores de estructuras.
```
