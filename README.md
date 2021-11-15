# Ed-Api
Esta es una API con los campos necesarios que debe tener una estructura curso y cumple lo necesario para un CRUD ( Create, Read, Update, Delete), el cual almacenamos los datos en memoria, exactamente en un `Map` con clave o llave de tipo `int` y cuerpo de tipo `Course` el cual es el nombre de la estructura de curso.
### Carpetas
`data` en esta carpeta implementamos la interfaz `Data` el cual se encargara de la  persistencia de los datos.<br>
`infrastructure` Hacemos toda la comunicación de la API, la ruta, la persistencia y el handler.