<template>
<div class="container">
<h1>Documentacion de la API</h1>
<article>
<p>
  Aqui se muestran todos los endpoints de la API
  además de los valores que se requieren y de los
  valores devueltos. Todas las llamadas se harán
  a https://{url_host}/api, por lo que si se dice
  que se va a hacer una request GET a /users en realidad
  la URL es https://{url_host}/api/users.
</p>
<p>
  También tenga en cuenta que algunos endpoints requieren
  que la request tenga un token de autorización. Incluso,
  alugnos endpoints requiren de tokens pertencientes a
  usuarios administradores.
</p>
<p>
  Por último, si en una URL aparece entre llaves ({})
  significa que es un argumento de URL.
</p>
<section>
  <h2>Usuarios</h2>
  <h3>[POST] /user/login</h3>
  <p>
  Permite hacer login a un usuario con su nombre y contraseña.
  Devuelve un token que debe ser usado para otras requests.
  </p>
  <h5>Request</h5>
  <pre>
{
  "name": "username",
  "password": "userpass"
}</pre>
  <h5>Response</h5>
  <pre>
{
  "name": "diego",
  "role": "admin",
  "token": {
    "value": "some large token",
    "expires": "2021-03-02T00:25:26.655679435+01:00",
    "owner": "diego",
    "role": "admin"
  }
}</pre>
  <h5>Errores</h5>
  <ul>
    <li>[100] User Not Found</li>
    <li>[101] Invalid password</li>
  </ul>

  <h3>(ADMIN) [POST] /user/create</h3>
  <p>
  Permite crear a un usuario con su nombre y contraseña. El rol
  de los usuarios creados de esta manera son siempre "user". Requiere
  de un token de un usuario con el rol ADMIN.
  </p>
  <h5>Request</h5>
  <pre>
{
  "name": "username",
  "password": "userpass"
}</pre>
  <h5>Response</h5>
  <pre>
{
    "Name": "username",
    "Role": "user"
}</pre>
  <h5>Errores</h5>
  <ul>
    <li>[102] User Already Exists</li>
  </ul>

  <h3>(ADMIN) [DELETE] /user/{name}</h3>
  <p>
  Borra a un usuario con el nombre "name". Requiere
  de un token de un usuario con el rol ADMIN.
  </p>
  <h5>Request</h5>
  <pre>https://localhost/api/user/manolo</pre>
  <h5>Errores</h5>
  <ul>
    <li>[102] User Already Exists</li>
  </ul>

<h3>(ADMIN) [GET] /user/all</h3>
  <p>
  Obtiene todos los usuarios del sistema. Requiere
  de un token de un usuario con el rol ADMIN.
  </p>
  <h5>Response</h5>
  <pre>
[
    {
        "Name": "u1",
        "Role": "user"
    },
    {
        "Name": "u2",
        "Role": "user"
    },
    {
        "Name": "u3",
        "Role": "user"
    },
    {
        "Name": "u4",
        "Role": "user"
    }
]</pre>
</section>

<section>
  <h2>Sensores</h2>
  <h3>[GET] /sensor/{name}</h3>
  <p>
  Obtiene información de un sensor identificado por el nombre
  "name".
  </p>
  <h5>Request</h5>
  <pre>https://localhost/api/sensor/habitacion</pre>
  <h5>Response</h5>
  <pre>
{
    "name": "habitacion",
    "connection": {
        "type": "http",
        "value": "192.168.1.1"
    },
    "updateInterval": 60,
    "deleted": false,
    "supportedReports": [
        "humidity",
        "temperature"
    ]
}</pre>
  <h5>Errores</h5>
  <ul>
    <li>[200] Sensor Not Found</li>
  </ul>

  <h3>(AUTH) [GET] /sensor/{name}/now</h3>
  <p>
  Realiza una lectura de un sensor identificado por el nombre
  "name". Requiere de un token.
  </p>
  <h5>Request</h5>
  <pre>https://localhost/api/sensor/habitacion/now</pre>
  <h5>Response</h5>
  <pre>
[
    {
        "type": "humidity",
        "sensor": "habitacion",
        "date": "2021-03-01T19:45:36.577Z",
        "value": 61.5
    },
    {
        "type": "temperature",
        "sensor": "habitacion",
        "date": "2021-03-01T19:45:36.577Z",
        "value": 18.7
    }
]</pre>
  <h5>Errores</h5>
  <ul>
    <li>[200] Sensor Not Found</li>
    <li>[201] Sensor Does Not Respond</li>
  </ul>

  <h3>[GET] /sensor/{name}/reports</h3>
  <p>
  Obtiene todos los reportes filtrados de un sensor llamado
  "name".
  </p>
  <h5>Request</h5>
  <p>
    Se trata de una request GET a la que se le puede añadir los siguientes
    filtros en cualquier orden y opcionalmente
  </p>
  <ul>
    <li>
      <strong>from</strong>: Fecha de inicio. Debe tener el siguiente formato:
      2021-02-24T18:18:00.000Z. Requiere que también se incluya "to".
    </li>
    <li>
      <strong>to</strong>: Fecha de fin. Debe tener el siguiente formato:
      2021-02-24T18:18:00.000Z. Requiere que también se incluya "from".
    </li>
    <li>
      <strong>type</strong>: Tipo de reporte.
    </li>
    <li>
      <strong>trim</strong>: Número de reportes que quedarse por debajo.
    </li>
  </ul>
  <pre>http://localhost/api/sensor/habitacion/reports?from=2021-02-24T18:18:00.000Z&to=2021-02-25T18:18:00.000Z&type=temperature</pre>
  <h5>Response</h5>
  <pre>
[
    {
        "type": "humidity",
        "sensor": "habitacion",
        "date": "2021-03-01T19:45:36.577Z",
        "value": 61.5
    },
    {
        "type": "temperature",
        "sensor": "habitacion",
        "date": "2021-03-01T19:45:36.577Z",
        "value": 18.7
    }
]</pre>
  <h5>Errores</h5>
  <ul>
    <li>[200] Sensor Not Found</li>
  </ul>

<h3>(ADMIN) [POST] /sensor</h3>
  <p>
  Crea un sensor. Require de un token de tipo ADMIN.
  </p>
  <h5>Request</h5>
  <pre>
{
    "name": "nombre del sensor",
    "connection": {
        "type": "http",
        "value": "ip del sensor"
    },
    "updateInterval": 60,
    "supportedReports": [
        "humidity",
        "temperature"
    ]
}
</pre>
  <h5>Response</h5>
  <pre>
{
    "name": "nombre del sensor",
    "connection": {
        "type": "http",
        "value": "ip del sensor"
    },
    "updateInterval": 60,
    "deleted": false,
    "supportedReports": [
        "humidity",
        "temperature"
    ]
}</pre>
  <h5>Errores</h5>
  <ul>
    <li>[202] Sensor Already Exists</li>
  </ul>

<h3>(ADMIN) [PATCH] /sensor</h3>
  <p>
  Actualiza un sensor. Require de un token de tipo ADMIN.
  </p>
  <h5>Request</h5>
  <pre>
{
    "name": "nombre del sensor",
    "connection": {
        "type": "http",
        "value": "ip del sensor"
    },
    "updateInterval": 60,
    "supportedReports": [
        "humidity",
        "temperature"
    ]
}
</pre>
  <h5>Response</h5>
  <pre>
{
    "name": "nombre del sensor",
    "connection": {
        "type": "http",
        "value": "ip del sensor"
    },
    "updateInterval": 60,
    "deleted": false,
    "supportedReports": [
        "humidity",
        "temperature"
    ]
}</pre>
  <h5>Errores</h5>
  <ul>
    <li>[200] Sensor not found</li>
  </ul>

<h3>(ADMIN) [DELETE] /sensor/{name}</h3>
  <p>
  Borra un sensor. Require de un token de tipo ADMIN.
  </p>
  <h5>Request</h5>
  <pre>https://localhost/api/sensor/habitacion</pre>
  <h5>Response</h5>
  <pre>
{
    "name": "habitacion",
    "connection": {
        "type": "http",
        "value": "192.168.1.1"
    },
    "updateInterval": 60,
    "deleted": true,
    "supportedReports": [
        "humidity",
        "temperature"
    ]
}</pre>
  <h5>Errores</h5>
  <ul>
    <li>[200] Sensor not found</li>
  </ul>

<h3>[GET] /sensors</h3>
  <p>
  Obtiene todos los sensores.
  </p>
  <h5>Request</h5>
  Por defecto mostrará solo los sensores que no han sido borrados.
  Para mostrar los borrados se necesita pasar como parámetro "deleted" con
  el valor "true"
  <pre>https://localhost/api/sensors?deleted=true</pre>
  <h5>Response</h5>
  <pre>
[
    {
        "name": "cocina",
        "connection": {
            "type": "http",
            "value": "192.168.1.1"
        },
        "updateInterval": 60,
        "deleted": false,
        "supportedReports": [
            "humidity",
            "temperature"
        ]
    },
    {
        "name": "electric",
        "connection": {
            "type": "http",
            "value": "192.168.1.2"
        },
        "updateInterval": 30,
        "deleted": true,
        "supportedReports": [
            "watts"
        ]
    },
    {
        "name": "habitacion",
        "connection": {
            "type": "http",
            "value": "192.168.1.3"
        },
        "updateInterval": 60,
        "deleted": false,
        "supportedReports": [
            "humidity",
            "temperature"
        ]
    },
    {
        "name": "salon",
        "connection": {
            "type": "http",
            "value": "192.168.1.4"
        },
        "updateInterval": 60,
        "deleted": false,
        "supportedReports": [
            "humidity",
            "temperature"
        ]
    }
]</pre>

<h3>[GET] /sensors/now/average</h3>
  <p>
  Obtiene la media de las lecturas de todos los sensores.
  </p>
  <h5>Request</h5>
  <pre>https://localhost/api/sensor/habitacion</pre>
  <h5>Response</h5>
  <pre>
[
    {
        "type": "temperature",
        "sensor": "average",
        "date": "2021-03-01T20:13:13.944Z",
        "value": 22.4
    },
    {
        "type": "humidity",
        "sensor": "average",
        "date": "2021-03-01T20:13:13.944Z",
        "value": 53.23
    }
]</pre>
</section>

<section>
  <h2>Tipos de reportes</h2>
  <h3>[GET] /reports/types</h3>
  <p>
  Obtiene todos los tipos de reportes del sistema
  </p>
  <h5>Response</h5>
  <pre>
[
  "temperature",
  "humidity",
  "watts"
]</pre>

<h3>(ADMIN) [POST] /reports/types/{name}</h3>
  <p>
  Permite crear un nuevo tipo de reporte. Requiere de un
  token de tipo ADMIN.
  </p>
  <h5>Response</h5>
  <pre>
{
    "ReportType": "yourtype"
}</pre>
  <h5>Errores</h5>
  <ul>
    <li>[301] ReportType Already Exists</li>
  </ul>

</section>

</article>
</div>
</template>

<style scoped>
.container {
  bottom: 0px;
  width: 100vw;
  padding: 20px 20vw 20px 20vw;
  margin-top: 40px;

  display: flex;
  flex-direction: column;

  text-align: justify;
}

pre {
  --main-color: var(--fg-weak-color);
  --text-color: white;

  width: 100%;
  display: block;

  border-style: solid;
  border-width: 2px;
  border-color: #29292e;
  border-radius: 3px;

  padding: 10px;

  background-color: #29292e;
  color: white;
}
</style>
