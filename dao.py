from tinydb import TinyDB, Query, where
import re
import random

db = TinyDB('hellorheaven.json')
responsesT = db.table('responses')
proposalsT = db.table('proposals')
statsT = db.table('stats')

q = Query()

#types
HELL = 1
HEAVEN = 2
CANCEL = 3
RESET = 4
STOP = 5

MAXVOTES = 10

def GetAllStats():
    return statsT.all()


def GetStats(user):
    user = "^{}$".format(user)
    return statsT.search(q.user.matches(user, flags=re.IGNORECASE))

def Insert(userdb):
    statsT.insert(userdb)

def Update(user, type):
    try:
        userdb = GetStats(user)
        isnew = False

        if userdb == []:
            userdb = {"user": user, 'hell': 0, 'heaven': 0}
            isnew = True
        else:
            userdb = userdb[0]

        if type == HELL:
            userdb['hell'] += 1
        elif type == HEAVEN:
            userdb['heaven'] += 1
        else:
            return False

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

