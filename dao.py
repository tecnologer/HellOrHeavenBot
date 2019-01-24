from tinydb import TinyDB, Query, where
import re
import random

db = TinyDB('hellorheaven.json')
responsesT = db.table('responses')
proposalsT = db.table('proposals')
statsT = db.table('stats')
customAnswerT = db.table('customanswer')
chatLogT = db.table('chatlog')

q = Query()

#types
HELL = 1
HEAVEN = 2
CANCEL = 3
RESET = 4
STOP = 5

MAXVOTES = 3

def GetAllStats():
    return statsT.all()


def GetStats(user, user_id=None):
    if not user_id is None:
        return statsT.search((q.user_id == user_id) | (q.user == user))
    
    return statsT.search(q.user == user)
    # user = "^{}$".format(user)
    # return statsT.search(q.user.matches(user, flags=re.IGNORECASE))

def Insert(userdb):
    statsT.insert(userdb)

def Update(user, type, user_id=None):
    try:
        userdb = GetStats(user, user_id)
        isnew = False

        if userdb == []:
            userdb = {"user": user, 'hell': 0, 'heaven': 0, "user_id": user_id}
            isnew = True
        else:
            userdb = userdb[0]

        if not "user_id" in userdb:
            userdb["user_id"] = user_id

        if type == HELL:
            userdb['hell'] += 1
        elif type == HEAVEN:
            userdb['heaven'] += 1

        if isnew:
           Insert(userdb)
        else:
            statsT.update(userdb, doc_ids=[userdb.doc_id])

        return True
    except ValueError:
        print ValueError
        return False


def GetAnswer(type):
    rs = responsesT.search((q.t == type))
    if len(rs) == 0:
        return {
            "a": "ahi 'ta!",
            "at": 1
        }
    elif len(rs) == 1:
        return rs[0]
    i = random.randint(0, len(rs)-1)
    return rs[i]


def InsertAnswer(res):
    responsesT.insert(res)

# Region Proposal
def InsertProposal(prop):
    record = {
        "proposal": prop,
        "upvote": 0,
        "downvote": 0,
        "voters": []
    }
    proposalsT.insert(record)

def GetRandomProposal(user_id):    
    rs = proposalsT.search( (~ q.voters.all([user_id])) & (q.upvote < MAXVOTES) & (q.downvote < MAXVOTES) )
    if len(rs) == 0:
        return None
    elif len(rs) == 1:
        return rs[0]

    i = random.randint(0, len(rs)-1)
    return rs[i]

def UpdateScore(user_id, prop, isUp):
    if isUp:
        prop["upvote"] += 1
    else:
        prop["downvote"] += 1
    
    prop["voters"].append(user_id)

    if prop["upvote"] >= MAXVOTES:
        InsertAnswer(prop["proposal"])
        proposalsT.remove(doc_ids=[prop.doc_id])
    else:
        proposalsT.update(prop, doc_ids=[prop.doc_id])


def InsertCustomAnswer(author, regex, atype, answer, chat_id=None):
    newAnswer={
        "regex": regex,
        "type": atype,
        "author": author,
        "chat_id": chat_id
    }

    if not answer is None:
        newAnswer["answer"] = answer

    customAnswerT.insert(newAnswer)


def GetCustomAnswer(chat_id):
    return customAnswerT.search( (q.chat_id == chat_id) | (q.chat_id==None))


def StoreChatLog(chat_id):
    try:
        userdb = GetChatLog(chat_id)
        if userdb != []:
            return
        
        chatlog = {"id": chat_id}
        chatLogT.insert(chatlog)
    except ValueError:
        print ValueError
        return False


def GetChatLog(chat_id=None):
    if chat_id is None:
        return chatLogT.all()

    return chatLogT.search(q.id == chat_id)
