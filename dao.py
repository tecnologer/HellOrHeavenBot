from tinydb import TinyDB, Query
import re
import random

db = TinyDB('hellorheaven.json')
responsesT = db.table('responses')
statsT = db.table('stats')

q = Query()

HELL = 1
HEAVEN = 2


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
    i = random.randint(0, len(rs)-1)
    return rs[i]["a"]


def InsertAnswer(res):
    responsesT.insert(res)
