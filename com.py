#!/usr/bin/python
# -*- coding: utf-8 -*-

import dao

DESC = "desc"
PARAMS = "params"
FUNC = "func"
WAIT = "needWait"

emLike = u'\U0001f44d'
emDislike = u'\U0001f44e'

class Answerype():
    TEXT = 1
    STICKER = 2
    GIF = 3
    PHOTO = 4

ticketWait = {}
proposalVoting = {}

def getMsgLeeMan():
    return {"a": 'que raro que tu... lee el manual!', "at": Answerype.TEXT}


def goToHell(user, userSender, chat_id):
    response = {
        "r": {
            "a": None,
            "at": Answerype.TEXT
        },
        "needWait": COMMANDS["/hell"][WAIT]
    }
    if user == '':
        response["r"] = getMsgLeeMan()
        response["needWait"] = False
        return response 

    if user != "test":
        dao.Update(user, dao.HELL)
    response["r"] = dao.GetAnswer(dao.HELL)
    return response


def goToHeaven(user, userSender, chat_id):
    response = {
        "r": {
            "a": None,
            "at": Answerype.TEXT
        },
        "needWait": COMMANDS["/heaven"][WAIT]
    }

    if user == '':
        response["r"] = getMsgLeeMan()
        response["needWait"] = False
        return response


    if user != "test":
        dao.Update(user, dao.HEAVEN)
    response["r"] = dao.GetAnswer(dao.HEAVEN)
    return response

def getStats(user, userSender, chat_id):
    response = {
        "r": {
            "a": None,
            "at": Answerype.TEXT
        },
        "needWait": COMMANDS["/stats"][WAIT]
    }
    stats = dao.GetStats(userSender)

    if stats != []:
        hell = stats[0]['hell']
        heaven = stats[0]['heaven']
        emoji = u'\U0001f608'

        if heaven > hell:
            emoji = u'\u271d\ufe0f'
        response["r"]["a"] = 'Heaven: {}, Hell: {} ... {}'.format(heaven, hell, emoji)
        return response

    response["r"]["a"] = '{} la estadisticas no importan, vas al infierno de cualquier manera.'.format(userSender)
    return response


def getAllStats(user, userSender, chat_id):
    response = {
        "r": {
            "a": None,
            "at": Answerype.TEXT
        },
        "needWait":  COMMANDS["/all"][WAIT]
    }
    if userSender != 'Tecnologer':
        response["r"]["a"] = "solo dios tiene ese poder"
        return response

    stats = dao.GetAllStats()

    if len(stats) == 0:
        response["r"]["a"] = "no hay nada"
        response["needWait"]= False
        return response

    res = ""
    for val in stats:
        heaven = val['heaven']
        hell = val["hell"]

        emoji = u'\U0001f608'

        if heaven > hell:
            emoji = u'\u271d\ufe0f'

        res += '- {} -> Heaven: {}, Hell: {} ... {}\n'.format(
            val["user"],  heaven, hell, emoji)
    
    response["r"]["a"] = res
    return response


def cancel(user, userSender, chat_id):
    response = {
        "r": {
            "a": None,
            "at": Answerype.TEXT
        },
        "needWait":  COMMANDS["/cancel"][WAIT]
    }
    if not chat_id in ticketWait:
        response["r"]["a"] = "que vas a cancelar, no tienes nada"
        return response
        
    del ticketWait[chat_id]
    response["r"] = dao.GetAnswer(dao.CANCEL)
    return response


def showHelp(user, userSender, chat_id):
    response = {
        "r": {
            "a": None,
            "at": Answerype.TEXT
        },
        "needWait":  COMMANDS["/help"][WAIT]
    }
    res = "Bot para telegram que registra las acciones buenas y malas de los usuarios.\n\n"
    for k, v in COMMANDS.iteritems():
        res += "- {} {}{}=> {}\n".format(k, v[PARAMS], "" if v[PARAMS]=="" else " ", v[DESC])
    
    response["r"]["a"] = res
    return response


def resetData(user, userSender, chat_id):
    response = {
        "r": {
            "a": None,
            "at": Answerype.TEXT
        },
        "needWait":  COMMANDS["/reset"][WAIT]
    }

    response["r"] = dao.GetAnswer(dao.RESET)

    return response

def VerifyAlias(cmd):
    if cmd in alias:
        return alias[cmd]
    
    return ""


def IsWaiting(chat_id, user_id):
    return chat_id in ticketWait and user_id in ticketWait[chat_id]

def GetWaitingCmd(chat_id, user_id):
    return "/hell" if ticketWait[chat_id][user_id] == dao.HELL else "/heaven"

def Wait(chat_id, user_id, type):
    ticketWait[chat_id] = {user_id: type}


def showAlias(comando, userSender, chat_id):
    response = {
        "r": {
            "a": None,
            "at": Answerype.TEXT
        },
        "needWait":  COMMANDS["/alias"][WAIT]
    }
    if comando == "":
        response["r"]["a"] = u'que raro que tu... lee el manual! (¬_¬)'
        return response

    res = "Alias para el comando {}\n\n".format(comando)
    for k, v in alias.iteritems():
        if v == comando:
            res += "- {}\n".format(k)


    response["r"]["a"] = res
    return response


def stop(comando, userSender, chat_id):
    return {
        "r": dao.GetAnswer(dao.STOP),
        "needWait":  COMMANDS["/stop"][WAIT]
    }

def proposalStartVoting(msg, *args):
    response = {
        "r": {
            "a": None,
            "at": Answerype.TEXT
        },
        "needWait":  COMMANDS["/voteanswer"][WAIT]
    }

    user_id = msg["from"]["id"]

    if user_id in proposalVoting:
        prop = proposalVoting[user_id]
    else:
        prop = dao.GetRandomProposal(user_id)
        if prop is None:
            response["r"]["a"] = "no hay nada para votar"
            return response

    _for = ""

    if prop["proposal"]["t"] == dao.HEAVEN:
        _for = "/heaven"
    elif prop["proposal"]["t"] == dao.HELL:
        _for = "/hell"
    elif prop["proposal"]["t"] == dao.CANCEL:
        _for = "/cancel"
    
    help = "Usa {} para darle un punto a favor, o {} para darle un punto en contra. Si llega a {} puntos a favor se usara como respuesta.".format(
        emLike, emDislike, dao.MAXVOTES)
    r = {"a": "", "at": Answerype.TEXT}
    if prop["proposal"]["at"] == Answerype.TEXT:
        r["a"] = 'Respondera "{}" despues de ejecutar el comando {}.\n{}'.format(
            prop["proposal"]["a"], _for, help)
    elif prop["proposal"]["at"] == Answerype.STICKER:
        r["a"] = 'Respondera el siguiente sticker despues de ejecutar el comando {}.\n{}'.format(
            _for, help)
        r["file_id"] = prop["proposal"]["a"]
        r["file_t"] = Answerype.STICKER
    elif prop["proposal"]["at"] == Answerype.GIF:
        r["a"] = 'Respondera el siguiente gif despues de ejecutar el comando {}.\n{}'.format(
            _for, help)
        r["file_id"] = prop["proposal"]["a"]
        r["file_t"] = Answerype.GIF
    elif prop["proposal"]["at"] == Answerype.PHOTO:
        r["a"] = 'Respondera la siguiente imagen despues de ejecutar el comando {}.\n{}'.format(
            _for, help)
        r["file_id"] = prop["proposal"]["a"]
        r["file_t"] = Answerype.PHOTO

    proposalVoting[user_id] = prop
    response["r"] = r
    return response

# definicion de comandos
COMMANDS = {
    "/hell":{
        FUNC: goToHell,
        DESC: "Se agrega al usuario un boleto al infierno",
        PARAMS: "<username>",
        WAIT: True
    }, 
    "/heaven": {
        FUNC: goToHeaven,
        DESC: "Se agrega al usuario un boleto al cielo",
        PARAMS: "<username>",
        WAIT: True
    },
    "/stats": {
        FUNC: getStats,
        DESC: "Muestra tus estadisticas",
        PARAMS: "",
        WAIT: True
    },
    "/all": {
        FUNC: getAllStats,
        DESC: "Modo Dios: Muestra todas las estadisticas",
        PARAMS: "",
        WAIT: True
    },
    "/cancel": {
        FUNC: cancel,
        DESC: "Cancela la peticion actual",
        PARAMS: "",
        WAIT: False
    },
    "/help": {
        FUNC: showHelp,
        DESC: "Muestra la informacion de los comandos",
        PARAMS: "",
        WAIT: False
    },
    "/reset": {
        FUNC: resetData,
        DESC: "Restablece tus estadisticas",
        PARAMS: "",
        WAIT: False
    },
    "/alias":{
        FUNC: showAlias,
        DESC: "Muestra el alias para el comando elegido",
        PARAMS: "</comando>",
        WAIT: False
    },
    "/stop": {
        FUNC: stop,
        DESC: '"Detiene" el bot. Evitaria que siguiera leyendo mensajes.',
        PARAMS: "",
        WAIT: False
    },
    "/addanswer": {
        # FUNC: addAnwser,
        DESC: "Añade una respuesta para un tipo de comando. Donde tipo puede tomar valor de:\n1.- Hell\n2.- Heaven\n3.- Cancel",
        PARAMS: "<tipo> [mensaje texto|sticker_id]",
        WAIT: False
    },
    "/voteanswer": {
        FUNC: proposalStartVoting,
        DESC: u"Te mostrara una propuesta de respuesta y esperara tu votacion usando: {} o {}".format(emLike, emDislike),
        PARAMS: "",
        WAIT: False
    }
}

alias = {
    u"/infierno": u"/hell",
    u"/cielo":    u"/heaven",
    u"/puntos":   u"/stats",
    u"/s":        u"/stats",
    u"/man":      u"/help",
    u"/ayuda":    u"/help",
    u"/?":        u"/help",
    u"/no":       u"/cancel",
    u"/cancela":  u"/cancel",
    u"/add":      u"/addanswer",
    u"/proposal": u"/addanswer",
    u"/propuesta": u"/addanswer",
    u"/vote":     u"/voteanswer",
    u"/votar":    u"/voteanswer",
}
