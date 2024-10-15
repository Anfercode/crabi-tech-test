
# USer Managment System - Crabi 🚘🦀🚘

This is the technical test implementation make with 💙 by [Andres Campero](https://github.com/Anfercode)




# Dependencies.

- Docker version 26.0.0 🐋
- GNU Make 4.2.1 💻



## Project Setup

After downloading the zip file, run the `make start` Command

```bash
  make start
```

this will install all the project dependencies


## See more utilities

To see more about the functionalities of the project run the following command.
```bash
  make help
```
# User Management API Reference

## Endpoints

### 1. CreateUser

Create a new user with the details provided.

#### POST /users/

**Request:**

| Parameter  | Type     | Description                        |
| :--------- | :------- | :--------------------------------- |
| `id`       | `string` | **Required**. The unique ID of the user. |
| `name`     | `string` | **Required**. The first name of the user. |
| `lastName` | `string` | **Required**. The last name of the user. |
| `email`    | `string` | **Required**. The email address of the user. Must be unique and valid. |
| `password` | `string` | **Required**. The password for the user account. |

**Responses:**

- `200 OK`: Returns a JSON representation of the created user.
- `400 Bad Request`: If any of the required fields are missing or the email is invalid.
- `500 Internal Server Error`: If there was an error processing the request.

### 2. Login

Authenticates a user based on their email and password.

#### POST /login/

**Request:**

| Parameter  | Type     | Description                        |
| :--------- | :------- | :--------------------------------- |
| `email`    | `string` | **Required**. The email address of the user. |
| `password` | `string` | **Required**. The password of the user account. |

**Responses:**

- `200 OK`: Returns a JSON representation of the authenticated user.
- `400 Bad Request`: If any of the required fields are missing.
- `401 Unauthorized`: If the email or password does not match.

### 3. GetUser

Retrieves detailed information about a user based on their email.

#### GET /users/{email}

**Request Parameters:**

| Parameter | Type     | Description                           |
| :-------- | :------- | :------------------------------------ |
| `email`   | `string` | **Required**. The email to retrieve. |

**Responses:**

- `200 OK`: Returns a JSON representation of the user.
- `404 Not Found`: If no user with the specified email exists.


## Second step solution

### Registro y autenticacion

Yo lo que haria en este caso es, implementar un social auth, ya sea con google o con otra herramienta, para no forzar a usuario a hacer "login", haria que llene un formulario desde la web donde nos provea informacion, dentro de esa informacion el correo, ya con el correo crearia un usuario "anonimo" en la plataforma, esto para no tener data duplicada o trivial, despues se le haria una solicitud al backend, donde valide si el correo existe o no, y hacer posteriormente un upsert del usuario proveido, y por ultimo se haria una integracion con calendly donde el usuario pueda tener la facilidad de agendar la videollamada correspondiente, esto excusivamente para los clientes, para los doctores se se manejaria de manera mas estricta la informacion de registro de doctores ya que ellos son usuario interno

### Recipe payment

En este caso haria que la consulta sea una maquina de estados, esto para poder manejar eventos especificos dependiendo del caso, ya con esto puede haber un status tanto de pending for payment y de payed, ya con esto cuando el status pase a payed, en el evento de dominio del cambio de status a payed ejecutaria un trigger que envie al correo del cliente o usuario el recipe proveido por el doctor, este manejo de eventos para el mvp lo implementaria con un event bus en memoria para validar que todo funcione correctamente, ya despues implementaria algo como un rabbitmq o un manejador de colas para el caso de que el negocio necesite crecer o escalar

### Tarjeta de presentacion

En este caso en especifico, haria un listing de doctores, con sus especialidades, en dado caso que el paciente necesite un trato mas especifico, los doctores podran estar taggeados por especializacion para facilitar la busqueda del paciente, para el caso de una consulta "general", haria una funcionalidad a parte donde sea un boton, donde el usuario llene un formulario con que patologia o dolencia tiene y despues hacer un sistema de colas para asignar los doctores dependiendo de la carga operativa de cada uno, por ejemplo que un doctor no pueda tener mas de *n* pacientes a la semana, asi me aseguro que cada doctor no tenga una carga mas que los demas, ahora para la interfaz de calendar management, lo que haria, es que el paciente propone un horario, dependiendo de la disponibilidad del profesional escogido o asignado, ahi al momento de que el usuario escoja una fecha, se le avisara al doctor, por correo, el horario escogido por el cliente, ya con esto el doctor confirmaria o reagendaria la cita

### Conclusion

Esto es un happy path propuesto por mi para el dado caso de que todo salga como se propuso, y el enfoque es full orientado a la parte de como se resolveria de una manera mas tecnica, por que un buen mvp puede ser simplemente hacer algo con no-code o con la suite de google para validar la idea, espero que la solucion planteada sea del agrado del equipo.

## Authors

- Andres Campero [@anfercode](https://www.github.com/anfercode)

