import dao

DESC = "desc"
PARAMS = "params"
FUNC = "func"
WAIT = "needWait"


ticketWait = {}


def goToHell(user, userSender, chat_id):
    if user == '':
        return 'que raro que tu... lee el manual!', COMMANDS["/hell"][WAIT]

    dao.Update(user,dao.HELL)
    return dao.GetAnswer(dao.HELL), COMMANDS["/hell"][WAIT]


def goToHeaven(user, userSender, chat_id):
    if user == '':
        return 'que raro que tu... lee el manual!', COMMANDS["/heaven"][WAIT]

    dao.Update(user, dao.HEAVEN)
    return dao.GetAnswer(dao.HEAVEN), COMMANDS["/heaven"][WAIT]


def getStats(user, userSender, chat_id):
    stats = dao.GetStats(userSender)

    if stats != []:
        hell = stats[0]['hell']
        heaven = stats[0]['heaven']
        emoji = u'\U0001f608'

        if heaven > hell:
            emoji = u'\u271d\ufe0f'

        return 'Heaven: {}, Hell: {} ... {}'.format(heaven, hell, emoji), False
   
    return '{} la estadisticas no importan, vas al infierno de cualquier manera.'.format(userSender), COMMANDS["/stats"][WAIT]

def getAllStats(user, userSender, chat_id):
    if userSender != 'Tecnologer':
        return "solo dios tiene ese poder", COMMANDS["/all"][WAIT]

    stats = dao.GetAllStats()

    if len(stats) == 0:
        return "no hay nada", False

    response = ""
    for val in stats:
        heaven = val['heaven']
        hell = val["hell"]

        emoji = u'\U0001f608'

        if heaven > hell:
            emoji = u'\u271d\ufe0f'

        response += '- {} -> Heaven: {}, Hell: {} ... {}\n'.format(
            val["user"],  heaven, hell, emoji)
    
    return response, COMMANDS["/all"][WAIT]

def cancel(user, userSender, chat_id):
    if not chat_id in ticketWait:
        return "que vas a cancelar, no tienes nada", COMMANDS["/cancel"][WAIT]
        
    del ticketWait[chat_id]
    return "che rajon!", COMMANDS["/cancel"][WAIT]


def showHelp(user, userSender, chat_id):
    res = ""
    for k, v in COMMANDS.iteritems():
        res += "- {} {}{}=> {}\n".format(k, v[PARAMS], "" if v[PARAMS]=="" else " ", v[DESC])
    
    return res, COMMANDS["/help"][WAIT]

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
        PARAMS: ""
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
    }
}

alias = {
    "/infierno": "/hell",
    "/cielo": "/heaven",
    "/puntos": "/stats",
    "/s": "/stats",
    "/man": "/help",
    "/ayuda": "/help",
    "/?": "/help"
}
