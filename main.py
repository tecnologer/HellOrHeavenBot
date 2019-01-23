import telepot
import sys
import time
import random
import os
import key
import com
import re
import dao
import base64
import customanswer as ca
from pprint import pprint

dirname = os.path.dirname(__file__)

reload(sys)
sys.setdefaultencoding('utf-8')
bot = telepot.Bot(key.BOT_KEY)  # token

timeout = {}
answerTransactions = {}

#emojis
emLike = u'\U0001f44d'
emDislike = u'\U0001f44e'

#stickers
tranquiloviejo = u"CAADAQADNwADzxSlAAEpVbCJbOTMsAI"
awanta = u'CAADAQADqwADJaHuBMhw3ty2zbpjAg'
dejesedemamadas = u'CAADAQAD7wEAAs8UpQABdurS64LRGooC'
alapifu = u"CAADAQADkQADJaHuBAABSnzPxbzjJQI"
ticketHell = u'CAADAQADnQADJaHuBGvY1E43XYjJAg'
ticketHeaven = u'CAADAQADswADJaHuBEcjnhhUIqsPAg'
terco = u'CAADAQADqgADJaHuBEK37px2YeW-Ag'
amivalevrgtmb = u'CAADAQADtAADJaHuBJdeO7iayOyQAg'
amivalevrg = u'CAADAQADiAADJaHuBD7kz0JCJne4Ag'
atodosvalevrg = u'CAADAQADigADJaHuBEbW2qfTwX5XAg'
uypuesperdon = u'CAADAQADogADJaHuBLFm_SWQWCPDAg'
foca_gaaay = u'CAADBAADcAQAApv7sgABifFfdnNmjjsC'
ora_bergha = u'CAADAQAD1wEAAiRSnAABvoSTCsK5ylcC'
kheberga = u'CAADAQADiwADJaHuBCxFUkncLVKjAg'
oseakhe = u'CAADBQADfgMAAukKyAMythx0wTDJDAI'
gatolike = u'CAADAQADpgADJaHuBGgS8JEkEOvuAg'
kike = u'CAADAQAD5AEAAiRSnAABTkxC4DiGyCMC'


# regex
iscoraline = r"\s?(k|c)(a|o)r(a|o)line\s?"

# gifs
mcdinero_gif = u'CgADAQADAQADLm_4TFkwvxivN4ncAg'
hagaaay_gif = u'CgADAwADAQADhjxQTo1Kz-gOAQ_jAg'
ikillu_gif = u'CgADBAADFaAAAloXZAe9o2B4i9CciwI'
racists_gif = u'CgADBAADwKMAAlEXZAcPm6zqHWX1DAI'
trabajaperro_gif = u'CgADBAADeRcAAsUdZAefc7VUnBenbwI'
maradona_gif = u'CgADBAAD758AAvgaZAfNzwLnrluCJAI'
carlton_gif = u'CgADBAADcZ8AAmgXZAfrrj1C3Ln98gI'
vanndame_street = u'CgADAwADBwADpFcITFLjOjfn8IfZAg'
vanndame_dancing1 = u'CgADBAADxY8AAlsdZAcjIh--PzkbtwI'
vanndame_dancing2 = u'CgADAwADAQADmEhBTBW3PNcv7nfcAg'


def handle(msg):
    pprint(msg)

def getUserSender(msg):
    if 'username' in msg['from']:
        return msg['from']['username']
    else:
        return msg['from'][u'first_name']


def isBot(msg):
    return 'from' in msg and 'is_bot' in msg['from'] and msg['from']['is_bot']

def reply(msg, response, replyTo=True):
    chat_id = msg['chat']['id']
    msgId = msg['message_id']
    if not replyTo:
        msgId = None

    bot.sendMessage(chat_id=chat_id, text=response, reply_to_message_id=msgId )


def replyDocument(msg, docid, replyTo = True):
    chat_id = msg['chat']['id']
    msgId = replyTo if msg['message_id'] else None

    bot.sendDocument(chat_id=chat_id, document=docid, reply_to_message_id=msgId)


def responseDocument(msg, docid, caption=None):
    chat_id = msg['chat']['id']
    bot.sendDocument(chat_id=chat_id, document=docid, caption=caption)


def responseImage(msg, photo, caption=None):
    chat_id = msg['chat']['id']
    photo = os.path.join(dirname, photo)
    bot.sendPhoto(chat_id, open(photo, 'rb'), caption)

def replySticker(msg, sticker, reply=True):
    try:
        chat_id = msg['chat']['id']
        msgId = msg['message_id']
        if not reply:
            msgId = None
        bot.sendSticker(chat_id=chat_id, sticker=sticker, reply_to_message_id=msgId)
    except:
        print("error al enviar sticker: ", sticker)
    

def getAwnser(type):
    return dao.GetAnswer(type)


def newRecord(sender):
    timeout[sender] = {"time": time.time(), "count": 0}

def wait(msg, type):
    sender = getUserSender(msg)
    chat_id = msg['chat']['id']
    user_id = msg["from"]['id']

    reply(msg, "a quien le damos ese?")
    newRecord(sender)

    com.Wait(chat_id, user_id, type)

def validTimeout(msg, sender):
    if not sender in timeout:
        return True

    elapsed = int(time.time()-timeout[sender]["time"])

    if elapsed <= 30:
        timeout[sender]["count"]+=1

        if timeout[sender]["count"] == 1:
            replySticker(msg, tranquiloviejo)
        elif timeout[sender]["count"] == 5:
            replySticker(msg, awanta)
        elif timeout[sender]["count"] == 10:
            replySticker(msg, dejesedemamadas)
        elif timeout[sender]["count"] == 20:
            replySticker(msg, alapifu)
        elif timeout[sender]["count"] == 25:
            replySticker(msg, dejesedemamadas)
        elif timeout[sender]["count"] == 30:
            replySticker(msg, terco)
        elif timeout[sender]["count"] % 3 == 0:
            reply(msg, "esperate {} segundos".format(30-elapsed))

        return False

    return True


def checkSpecialWords(msg):
    if not "text" in msg:
        return

    customAnswer = ca.ValidateCustomAnswer(msg)
    if not customAnswer is None:
        manageResponse(msg, customAnswer["r"])
        return

    # if textMatch(inchebot, msg['text']):
    #     replySticker(msg, ora_bergha, False)
    if textMatch(iscoraline, msg['text']) and msg["from"]["id"] == 17760842:
        reply(msg, "si seras, si seras, que se llama Karelia, che terco!")
    

def textMatch(regex, test_str):
    matches = re.search(
        regex, test_str, re.MULTILINE | re.UNICODE | re.IGNORECASE)

    return matches is not None

def manageResponse(msg, response):
    answer = response["a"]
    answerype = response["at"]

    if answerype == com.AnswerType.STICKER:
        replySticker(msg, answer)
    elif answerype == com.AnswerType.GIF:
        replyDocument(msg, answer)
    elif answerype == com.AnswerType.PHOTO:
        responseImage(msg, answer)
    else:  # answerype == com.Answerype.TEXT:
        reply(msg, answer)
    
    if not "file_id" in response and not "file_t" in response:
        return

    if response["file_t"] == com.AnswerType.STICKER:
        replySticker(msg, response["file_id"])
    elif response["file_t"] == com.AnswerType.GIF:
        replyDocument(msg, response["file_id"])
    elif response["file_t"] == com.AnswerType.PHOTO:
        responseImage(msg, response["file_id"])



def waitForAnswer(chat_id, user_id, tipo):
    answerTransactions[chat_id] = {user_id: tipo}

def isWaitingAnswer(chat_id, user_id):
    return chat_id in answerTransactions and user_id in answerTransactions[chat_id]

def checkWaitingAnswer(msg):
    chat_id = msg['chat']['id']
    user_id = msg["from"]['id']

    if not isWaitingAnswer(chat_id, user_id):
        return False

    answer = ""
    answerType = -1
    if "text" in msg:
        answer = msg["text"]
        if answer.startswith("/cancel"):
            del answerTransactions[chat_id][user_id]
            response = dao.GetAnswer(dao.CANCEL)
            manageResponse(msg, response)
            return True
        elif answer.startswith("/"):
            reply(msg, "Eso es un comando, no una respuesta", False)
            return True
        answerType = com.AnswerType.TEXT
    elif "sticker" in msg:
        answer = msg["sticker"]['file_id']
        answerType = com.AnswerType.STICKER
    elif "animation" in msg:
        answer = msg["animation"]['file_id']
        answerType = com.AnswerType.GIF
    
    if answerType == -1:
        del answerTransactions[chat_id][user_id]
        replySticker(msg, dejesedemamadas)
        reply(msg, "Solo texto, sticker o gif. Ahora por vivo, tienes que volver a empezar", False)
        return True
    
    answerObj = {
        "t": answerTransactions[chat_id][user_id],
        "a": answer,
        "at": answerType
    }

    dao.InsertProposal(answerObj)

    reply(
        msg, "Listo, la propuesta para respuesta ha sido almacenada!. Usa el comando /voteanswer para votar por las respuestas mas originales.")

    del answerTransactions[chat_id][user_id]
    return True

def addAnswer(msg):
    chat_id = msg['chat']['id']
    user_id = msg["from"]['id']
    tokens = msg["text"].split(" ", 2)
    
    #solo comando
    if len(tokens) < 2:
        reply(
            msg, "Es necesario especificar el tipo. /addanswer {}".format(com.COMMANDS["/addanswer"][com.PARAMS]))
        return
        
    tipo = tokens[1]

    if not tipo.isdigit():
        reply(
            msg, "Tipo debe ser un numero y solo puede tomar valor de:\n1.- Hell\n2.- Heaven\n3.- Cancel")
        return

    tipo = int(tipo)
    if tipo < 1 and tipo > 3:
        reply(
            msg, "Tipo solo puede tomar valor de:\n1.- Hell\n2.- Heaven\n3.- Cancel")
        return

    #solo comando y tipo
    if len(tokens) == 2:
        waitForAnswer(chat_id, user_id, tipo)
        reply(
            msg, "Manda lo que quieras que responda. Puede ser un mensaje de texto, sticker o gif.")
        return
    
    answer = tokens[2].strip()
    answerObj = {
        "t": tipo,
        "a": answer,
        "at": com.AnswerType.TEXT
    }
    dao.InsertProposal(answerObj)

    reply(
        msg, "Listo, la propuesta para respuesta ha sido almacenada!. Usa el comando /voteanswer para votar por las respuestas mas originales.")


def IsVatotation(msg):
    user_id = msg["from"]["id"]
    return user_id in com.proposalVoting and "text" in msg and (msg["text"].startswith(emLike) or msg["text"].startswith(emDislike))


def AddVotation(msg):
    user_id = msg["from"]["id"]
    isUp = msg["text"].startswith(emLike)
    dao.UpdateScore(user_id, com.proposalVoting[user_id], isUp)
    del com.proposalVoting[user_id]
    replySticker(msg, gatolike, False)

def isEditing(msg):
    return "edit_date" in msg

def checkDocuments(msg):
    if not "document" in msg or not "file_id" in msg["document"]:
        return False
    docId = msg["document"]["file_id"]
    
    if docId == trabajaperro_gif:
        reply(msg, "trabaja, perro!", False)
        return True
    
    if docId == vanndame_dancing1 or docId == vanndame_dancing2:
        replyDocument(msg, vanndame_street)
        return True
    
    if docId == vanndame_street:
        r = random.randint(1, 2)
        if r == 1:
            replyDocument(msg, vanndame_dancing1, False)
        else:
            replyDocument(msg, vanndame_dancing2, False)

        return True

    return False

def on_chat_message(msg):
    if isBot(msg) or isEditing(msg):
        return 

    if checkWaitingAnswer(msg):
        return
    
    if ca.IsWaiting(msg):
        response = ca.ValidateMsg(msg)
        if "r" in response:
            manageResponse(msg, response["r"])
        return

    if IsVatotation(msg):
        AddVotation(msg)
        return 
    #u'file_id' (140128345684368):u'CgADBAADeRcAAsUdZAefc7VUnBenbwI'

    if checkDocuments(msg):
        return

    # if not has text or sticker
    if (not 'text' in msg and not 'sticker' in msg):
        return
    

    cmd = ''
    user = ''
    userSender = getUserSender(msg)

    ignoreTimeout = False
    chat_id = msg['chat']['id']
    user_id = msg["from"]['id']

    """  if chat_id in ticketWait and ticketWait[chat_id] != user_id:
        reply(msg, "a ti no te pregunte, metiche!")
        return """
    if com.IsWaiting(chat_id, user_id) and 'text' in msg:
        user = msg['text'].split(' ')[0].replace('@', '')

        if user.startswith("/cancel"):
            manageResponse(msg, com.cancel("", "", chat_id, msg)["r"])
            return 
        ignoreTimeout = True
        cmd = com.GetWaitingCmd(chat_id, user_id)
        com.cancel("", "", chat_id, msg)
    elif 'text' in msg:
        cmds = msg['text'].split(' ')
        response = ''
        if len(cmds) >= 2:
            cmd = cmds[0]
            user = cmds[1].replace('@', '')        
        elif len(cmds) == 1:
            cmd = cmds[0]
    elif 'sticker' in msg and msg['sticker']['file_id']==ticketHell:
        if not validTimeout(msg, userSender):
            return        
        wait(msg, dao.HELL)
        return
    elif 'sticker' in msg and msg['sticker']['file_id'] == ticketHeaven:
        if not validTimeout(msg, userSender):
            return
        wait(msg, dao.HEAVEN)
        return
    elif 'sticker' in msg and msg['sticker']['file_id'] == amivalevrg:
        replySticker(msg, amivalevrgtmb)
        return
    elif 'sticker' in msg and msg['sticker']['file_id'] == foca_gaaay:
        responseDocument(msg, hagaaay_gif)
        return

    if not cmd.startswith('/') and not 'sticker' in msg:
        checkSpecialWords(msg)
        return

    # acept commands type /command@HellOrHeavenBot
    cmd = cmd.replace('@HellOrHeavenBot', '')

    if not cmd in com.COMMANDS:
        cmd = com.VerifyAlias(cmd)
        if cmd == "":
            return

    # flag to validate or not the timeout
    if not ignoreTimeout:
        ignoreTimeout = not com.COMMANDS[cmd][com.WAIT]

    if not ignoreTimeout and not validTimeout(msg, userSender):
        return

    if user.upper() == 'HELLORHEAVENBOT':
        reply(msg, 'si tu, voy corriendo!')
        return 

    if user.upper() == userSender.upper():
        reply(msg, u'solo dios puede juzgarte... nah!, los demas lo haran \U0001f602')
        return
    
    if cmd.startswith("/addanswer"):
        addAnswer(msg)
    elif cmd.startswith("/voteanswer"):
        response = com.proposalStartVoting(msg)
        manageResponse(msg, response["r"])
    else:
        response = com.COMMANDS[cmd][com.FUNC](user, userSender, chat_id, msg)
        manageResponse(msg, response["r"])

        needWait = "needWait" in response and response["needWait"]
        if needWait and userSender != "Tecnologer" and user != "test":
            newRecord(userSender)
        

bot.message_loop({'chat': on_chat_message})

print('Listening ...')

while 1:
    time.sleep(10000)
