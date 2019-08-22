# HellOrHeavenBot

Bot multifuncion para telegram, soporta una variedad de [comandos](#comandos).

La tarea principal para que fue creado es llevar un contador de boletos al infierno (`/hell`) o al cielo (`/heaven`), estos "boletos" son registrados por otro usuario. Cada usuario puede revisar sus estadisticas con el comando `/stats`.

Adicional a esto se agrego un sistema para que los usuarios pudieran agregar respuestas (`/addanswer`) que el bot usara cuando se asigne un nuevo boleto. Estas respuestas son guardadas en una base de datos tipo JSON ([tinydb][2]), cuando se requiera el bot buscara y eligira una respuesta al azar.

Separado del tema principal, puede reaccionar a palabras que se le envien, para agregar reacciones se usa el comando `/customanswer` el cual necesita una expresion regular, dicha expresion sera evaluada y si el mensaje cumple con ella, enviara la respuesta que se le asigno. Ejemplos:

```
# si un usuario escribe "hola", el bot respondera con un "hola"
> /customanswer hola
> hola

# si le preguntan "como te llamas\?", el responde "me llamo luci"
> /customanswer como te llamas\?
> me llamo luci
```

**Nota:** Cada simbolo `>` indica que es un mensaje diferente.

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
- `/customanswer <ReGex>` => Agregara una respuesta personalizada. Cuando se cumpla la expresion regular respondera con lo que se le indique mensaje, sticker o gif.
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

- [Crear un bot en telegram](#telegram-bot)
- [Python 3.7.x][4]
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

# Telegram bot

- Envia el comando `/newbot` a [BotFather][5]
- Te solicitara el nombre de tu bot. Ejemplo: MiPrimerBot
- Despues es necesario asignarle un nombre usuario. Ejemplo: MiPrimerBot (este ya estara en uso, sera necesario seleccionar otro)
- Te enviara un mensaje como este:

  ```
  Done! Congratulations on your new bot. You will find it at t.me/MiPrimerBot. You can now add a description, about section and profile picture for your bot, see /help for a list of commands. By the way, when you've finished creating your cool bot, ping our Bot Support if you want a better username for it. Just make sure the bot is fully operational before you do this.

    Use this token to access the HTTP API:
    <BOT_TOKEN>
    Keep your token secure and store it safely, it can be used by anyone to control your bot.

    For a description of the Bot API, see this page: https://core.telegram.org/bots/api
  ```

- `<BOT_TOKEN>` sera el valor a reemplazar en el archivo `key.py`

Pruebalo: [t.me/hellorheavenbot][1]

[1]: https://t.me/hellorheavenbot
[2]: https://tinydb.readthedocs.io/en/latest/getting-started.html#installing-tinydb
[3]: https://telepot.readthedocs.io/en/latest/
[4]: https://www.python.org/downloads/
[5]: https://t.me/botfather
