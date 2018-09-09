import telepot
import sys
import time
import random
import os
from pprint import pprint
reload(sys)
sys.setdefaultencoding('utf-8')
bot = telepot.Bot('684372282:AAEHSrUFkvdOoCci8c0FzZ0H39BGfJ5Zbxc')  # token

stats = {
    'PAKOY3K': {
        'hell': 90,
        'heaven': 1
    },
    'XHAUL': {
        'hell': 2,
        'heaven': 0
    },
    'VSANCHEZ': {
        'hell': 3,
        'heaven': 0
    },
    'TECNOLOGER': {
        'hell': 3,
        'heaven': 2
    }
}

awnsers = {
    'hell': [
        'eso es todo, sigue asi',
        'una raya mas al tigre',
        'ha ha!',
        'anda la osa!',
        'fierro pariente... eso te espera xD',
        u'\U0001f608',
        u'\U0001f449\U0001f44c',
        'peleishon de bananeishon!'
    ],
    'heaven': [
        'una de cal por las que van de arena',
        'no importa,tu bien sabes a donde perteneces',
        'amen!',
        'ni tu te la crees',
        'el pastel era mentira',
        'santa no existe',
        u'\u2b50\ufe0f'
    ]
}

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

def getAwnser(type):
    i = random.randint(0, len(awnsers[type])-1)
    return awnsers[type][i]

def on_chat_message(msg):

    if isBot(msg):
        return
    # pprint(msg)
    # pprint(msg['text'])
    cmds = msg['text'].split(' ')
    cmd = ''
    user = ''
    response = ''
    if len(cmds) >= 2:
        cmd = cmds[0]
        user = cmds[1].upper().replace('@', '')        
    elif len(cmds) == 1:
        cmd = cmds[0]

    if not cmd.startswith('/'):
        return

    if user == 'HELLORHEAVENBOT':
        reply(msg, 'si tu, voy corriendo!')
        return 

    userSender = getUserSender(msg)
    if user == userSender.upper():
        reply(msg, u'solo dios puede juzgarte... nah!, los demas lo haran \U0001f602')
        return
    cmd = cmd.replace('@HellOrHeavenBot','')
    if cmd == "/hell":
        if user == '':
            reply(msg, 'que raro que tu... lee el manual!')
            return    

        if user in stats:
            stats[user]['hell'] += 1
        else:
            stats[user] = {'hell': 1, 'heaven': 0}
        
        response = getAwnser('hell')
    elif cmd == "/heaven":
        if user == '':
            reply(msg, 'que raro que tu... lee el manual!')
            return

        if user in stats:
            stats[user]['heaven'] += 1
        else:
            stats[user] = {'hell': 0, 'heaven': 1}
            
        response = getAwnser('heaven')
    elif cmd == "/stats":
        userKey = userSender.upper()

        if userKey in stats:
            hell = stats[userKey]['hell']
            heaven = stats[userKey]['heaven']
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
        
        if len(stats) == 0:
            return

        for key, val in stats.items():
            heaven = val['heaven']
            hell = val["hell"]
            
            emoji = u'\U0001f608'

            if heaven > hell:
                emoji = u'\u271d\ufe0f'

            response += '- {} -> Heaven: {}, Hell: {} ... {}\n'.format(key,  heaven, hell, emoji)
    else:
        response = 'por no saber leer manuales te iras al infierno'
    
    reply(msg, response)

bot.message_loop({'chat': on_chat_message})

print('Listening ...')

while 1:
    time.sleep(10)
