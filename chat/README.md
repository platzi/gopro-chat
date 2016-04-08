# Bienvenido a Revel

### Iniciar el servidor web:

    revel run myapp

   Ejecutar <tt>--help</tt> para ver las opciones.

### Ingresa a http://localhost:9000/ y podrás ver:

"It works"

### Descripción de contenido

La estructura por defecto de una aplicación generada con revel:

    myapp               Raíz del proyecto
      app               Contenido de la aplicación
        controllers     Controladores de la aplicación
          init.go       Registro de interceptors
        models          Modelos de la aplicación
        routes          Rutas (código generado)
        views           Vistas
      tests             Directorio de pruebas
      conf              Archivos de configuración
        app.conf        Archivo principal de configuración
        routes          Definición de las rutas
      messages          Archivos de mensajes
      public            Recursos publicos
        css             Archivos CSS
        js              Archivos javascript
        images          Archivos de imágenes

app

    El directorio app contiene el código fuente y plantillas para la aplicación.

conf

    El directorio conf contiene los archivos de configuración. Hay dos archivos principales:

    * app.conf, principal archivo de configuración de la aplicación, el cual contiene los parámetros de la configuración estándar. 
    * routes, archivo de la definición de rutas.


messages

    Contiene todos los archivos de mensajes -localizados- .

public

    Almacena los recursos estáticos que serán servidos directamente por el servidor Web. Tipicamente se dividen en tres directorios: images, css y js.

    El nombre de esos directorios pueden ser cualquier cosa; el desarrollador necesita actualizar las rutas.

test

    Directorio para las pruebas. Revel brinda un framework de pruebas que hace más fácil escribir y ejecutar pruebas funcionales de la aplicación.

### Una guía para empezar a desarrollar:

* [Introducción a Revel](http://revel.github.io/tutorial/index.html).
* [El manual de Revel](http://revel.github.io/manual/index.html).
* [Aplicaciones de ejemplo con Revel](http://revel.github.io/samples/index.html).
* [La documentación del API](http://revel.github.io/docs/godoc/index.html).

## Colabora

Te invitamos a contribuir a Revel, por favor hecha un vistazo a [Guía para colaborar con Revel](https://github.com/nubleer/revel/blob/master/CONTRIBUTING.md) para conocer las normas de colaboración [Join us](https://groups.google.com/forum/#!forum/revel-framework)!
