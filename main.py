import telepot
import sys
import time
import random
import os
import key
from pprint import pprint
import dao

reload(sys)
sys.setdefaultencoding('utf-8')
bot = telepot.Bot(key.BOT_KEY)  # token

timeout = {}
ticketToHell={}

tranquiloviejo = u"CAADAQADNwADzxSlAAEpVbCJbOTMsAI"
awanta = u'CAADAQADqwADJaHuBMhw3ty2zbpjAg'
dejesedemamadas = u'CAADAQAD7wEAAs8UpQABdurS64LRGooC'
alapifu = u"CAADAQADkQADJaHuBAABSnzPxbzjJQI"
ticket = u'CAADAQADnQADJaHuBGvY1E43XYjJAg'
terco = u'CAADAQADqgADJaHuBEK37px2YeW-Ag'

def handle(msg):
    pprint(msg)

def getUserSender(msg):
    if 'username' in msg['from']:
        return msg['from']['username']
    else:
        return msg['from'][u'first_name']


def isBot(msg):
    return 'from' in msg and 'is_bot' in msg['from'] and msg['from']['is_bot']

def reply(msg, response):
    chat_id = msg['chat']['id']
    msgId = msg['message_id']
    bot.sendMessage(chat_id=chat_id, text=response, reply_to_message_id=msgId)


def replySticker(msg, sticker):
    chat_id = msg['chat']['id']
    msgId = msg['message_id']
    bot.sendSticker(chat_id=chat_id, sticker=sticker, reply_to_message_id=msgId)

def getAwnser(type):
    return dao.GetAnswer(type)


def newRecord(sender):
    timeout[sender] = {"time": time.time(), "count": 0}


def waitToHell(msg):
    sender = getUserSender(msg)
    chat_id = msg['chat']['id']
    user_id = msg["from"]['id']
    
    reply(msg, "a quien le damos ese?")    
    newRecord(sender)

    ticketToHell[chat_id]=user_id


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
            replySticker(msg, dejesedemamadas)
        elif timeout[sender]["count"] == 30:
            replySticker(msg, terco)

        return False

    return True

def on_chat_message(msg):
    # if not has text or sticker
    if isBot(msg) or (not 'text' in msg and not 'sticker' in msg):
        return

    cmd = ''
    user = ''
    userSender = getUserSender(msg)

    ignoreTimeout = False
    chat_id = msg['chat']['id']
    user_id = msg["from"]['id']

    if chat_id in ticketToHell and ticketToHell[chat_id] != user_id:
        reply(msg, "a ti no te pregunte, metiche!")
        return
    elif chat_id in ticketToHell and ticketToHell[chat_id] == user_id and 'text' in msg:
        ignoreTimeout = True
        cmd = "/hell"
        user = msg['text'].split(' ')[0].replace('@', '')
        del ticketToHell[chat_id]
    elif 'text' in msg:
        cmds = msg['text'].split(' ')
        response = ''
        if len(cmds) >= 2:
            cmd = cmds[0]
            user = cmds[1].replace('@', '')        
        elif len(cmds) == 1:
            cmd = cmds[0]
    elif 'sticker' in msg and msg['sticker']['file_id']==ticket:
        if not validTimeout(msg, userSender):
            return        
        waitToHell(msg)
        return

    if not cmd.startswith('/') and not 'sticker' in msg:
        return

    if not ignoreTimeout and not validTimeout(msg, userSender):
        return

    if user.upper() == 'HELLORHEAVENBOT':
        reply(msg, 'si tu, voy corriendo!')
        return 

    if user.upper() == userSender.upper():
        reply(msg, u'solo dios puede juzgarte... nah!, los demas lo haran \U0001f602')
        return
    cmd = cmd.replace('@HellOrHeavenBot','')
    
    if cmd == '':
        return
    elif cmd == "/hell":
        if user == '':
            reply(msg, 'que raro que tu... lee el manual!')
            return 

        dao.Update(user,dao.HELL)
        response = getAwnser(dao.HELL)
        newRecord(userSender)
    elif cmd == "/heaven":
        if user == '':
            reply(msg, 'que raro que tu... lee el manual!')
            return

        dao.Update(user, dao.HEAVEN)
        response = getAwnser(dao.HEAVEN)
        
        newRecord(userSender)
    elif cmd == "/stats":
        stats = dao.GetStats(userSender)

        if stats != []:
            hell = stats[0]['hell']
            heaven = stats[0]['heaven']
            emoji = u'\U0001f608'
            
            if heaven > hell:
                emoji = u'\u271d\ufe0f'

            response = 'Heaven: {}, Hell: {} ... {}'.format(heaven, hell, emoji)
        else:
            response = '{} la estadisticas no importan, vas al infierno de cualquier manera.'.format(userSender)
    elif cmd == "/man" or cmd == '/help' or cmd == '/?':
        response = '- /hell <username>: el usuario gana un boleto directo al infierno muajaja\n- /heaven <username>: le dices que esa persona ha obrado bien\n- /stats : ves tus estadisticas'
    elif cmd == '/all':
        user = msg['from']['username']
        if user != 'Tecnologer':
            reply(msg, "solo dios tiene ese poder")
            return
        
        stats = dao.GetAllStats()

        if len(stats) == 0:
            return

        for val in stats:
            heaven = val['heaven']
            hell = val["hell"]
            
            emoji = u'\U0001f608'

            if heaven > hell:
                emoji = u'\u271d\ufe0f'

            response += '- {} -> Heaven: {}, Hell: {} ... {}\n'.format(val["user"],  heaven, hell, emoji)
    else:
        return
        # response = 'por no saber leer manuales te iras al infierno'
        
    reply(msg, response)

bot.message_loop({'chat': on_chat_message})

print('Listening ...')

while 1:
    time.sleep(10000)
