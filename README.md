# HellOrHeavenBot

Bot para telegram que registra las acciones buenas y malas de los usuarios.

## Comandos

- `/addanswer <tipo> [mensaje texto|sticker_id]` => Añade una respuesta para un tipo de comando. Donde tipo puede tomar valor numerico de la siguiente lista:
  1. Hell
  2. Heaven
  3. Cancel
- `/help` => Muestra la informacion de los comandos
- `/allahmode` => Activa el modo Allah.
- `/heaven <username>` => Se agrega al usuario un boleto al cielo
- `/cancel` => Cancela la peticion actual
- `/voteanswer` => Te mostrara una propuesta de respuesta y esperara tu votacion usando: 👍 o 👎
- `/stats` => Muestra tus estadisticas
- `/customanswer <ReGex> [mensaje texto]` => Agregara una respuesta personalizada. Cuando se cumpla la expresion regular respondera con lo que se le indique mensaje, sticker o gif.
- `/reset` => Restablece tus estadisticas
- `/all` => Modo Dios: Muestra todas las estadisticas
- `/alias </comando>` => Muestra el alias para el comando elegido
- `/broadcast` => Envia un mensaje a todos los chats que se han comunicado con el bot.
- `/getchatid` => Retorna el id del chat en base a un nombre.
- `/stop` => "Detiene" el bot. Evitaria que siguiera leyendo mensajes.
- `/direct` => Envia un mensaje a directo a un chat.
- `/ping` => Retorna el tiempo que ha pasado desde que se ejecuto.
- `/hell <username>` => Se agrega al usuario un boleto al infierno

## Dependencias y requisitos

- [Python 2.7.x][4]
- [TinyDb][2]: `pip install tinydb`
- [telepot][3]: `pip install telepot`

### Ejecutar con Docker

Require docker instalado y configurado en variables de entorno.

1. `docker build .` => esto arrojara un id al finalizar el proceso

   ```
   > docker build .
   Sending build context to Docker daemon  414.2kB
   Step 1/5 : FROM python:2
   ---> d75b4eed9ada
   Step 2/5 : RUN pip install tinydb
   ---> Using cache
   ---> b4b987febb7a
   Step 3/5 : RUN pip install telepot
   ---> Using cache
   ---> cdeba640ee41
   Step 4/5 : ADD . /
   ---> faa785f24bcf
   Step 5/5 : CMD ["python","-u", "/main.py"]
   ---> Running in 97ae0e70909b
   Removing intermediate container 97ae0e70909b
   ---> c6f3704701a8
   Successfully built c6f3704701a8
   ```

2. `docker run <build_id>`

   ```
   > docker run c6f3704701a8
   Listening ...
   ```

Pruebalo: [t.me/hellorheavenbot][1]

[1]: https://t.me/hellorheavenbot
[2]: https://tinydb.readthedocs.io/en/latest/getting-started.html#installing-tinydb
[3]: https://telepot.readthedocs.io/en/latest/
[4]: https://www.python.org/downloads/release/python-278/#download
