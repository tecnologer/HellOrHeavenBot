import dao
import re
import com
import random

ALLAH = 10244644
ALLAH_CHAT = 10244644

# listado de las expresiones que aun no tienen respuesta
WaitingAnswer = {}


def AddForWaiting(chat_id, author, regex):
    if not chat_id in WaitingAnswer:
        WaitingAnswer[chat_id] = {}

    WaitingAnswer[chat_id][author] = {
        "regex": regex,
        "type": None,
        "author": author,
        "answer": None,
        "chat_id": chat_id
    }

    #print WaitingAnswer


def IsWaiting(msg):
    chat_id = msg["chat"]["id"]
    author = msg["from"]["id"]
    return chat_id in WaitingAnswer and author in WaitingAnswer[chat_id]


def ValidateMsg(msg):
    response = {
        "r": {
            "a": None,
            "at": com.AnswerType.TEXT
        },
        "needWait": False
    }
    chat_id = msg["chat"]["id"]
    author = msg["from"]["id"]
    answerType = -1

    if "text" in msg:
        answer = msg["text"]
        if answer.startswith("/cancel"):
            del WaitingAnswer[chat_id][author]
            return dao.GetAnswer(dao.CANCEL)
        elif answer.startswith("/"):
            response["r"]["a"] = "Eso es un comando, no una respuesta"
            return response
        answerType = com.AnswerType.TEXT
    elif "sticker" in msg:
        answer = msg["sticker"]['file_id']
        answerType = com.AnswerType.STICKER
    elif "animation" in msg:
        answer = msg["animation"]['file_id']
        answerType = com.AnswerType.GIF

    if answerType == -1:
        del WaitingAnswer[chat_id][author]
        response["r"] = {
            "a": u'CAADAQAD7wEAAs8UpQABdurS64LRGooC',
            "at": com.AnswerType.STICKER
        }
        return response

    return AddCustomAnswer(chat_id, author, answerType, answer)


def AddCustomAnswer(chat_id, author, aType, answer):
    response = {
        "r": {
            "a": None,
            "at": com.AnswerType.TEXT
        },
        "needWait": True
    }
    try:
        ca = WaitingAnswer[chat_id][author]
        chatid = None if author == ALLAH and chat_id == ALLAH_CHAT else chat_id
        dao.InsertCustomAnswer(author, ca["regex"], aType, answer, chatid)
        response["r"]["a"] = "Listoooo!"
    except:
        response["r"]["a"] = "Algo no salio como esperaba, intenta de nuevo",
        response["needWait"] = False

    del WaitingAnswer[chat_id][author]

    return response


def ValidateCustomAnswer(msg):
    chat_id = msg["chat"]["id"]
    results = dao.GetCustomAnswer(chat_id)
    text = msg["text"]
    matches = []
    for result in results:
        try:
            regex = re.compile(
                result["regex"], re.MULTILINE | re.UNICODE | re.IGNORECASE)
        except:
            print "error with `{}`".format(result["regex"])
            continue

        if textMatch(regex, text):
            matches.append({
                "r": {
                    "a": result["answer"],
                    "at": result["type"]
                },
                "needWait": False
            })

    if len(matches) == 0:
        return None
    elif len(matches) == 1:
        return matches[0]

    i = random.randint(0, len(matches)-1)

    return matches[i]


def textMatch(regex, test_str):
    matches = re.search(
        regex, test_str)

    return matches is not None
