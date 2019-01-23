import dao
import re
import com
# listado de las expresiones que aun no tienen respuesta
WaitingAnswer = {}


def AddForWaiting(regex, author):
    WaitingAnswer[author] = {
        "regex": regex,
        "type": None,
        "author": author,
        "answer": None
    }

def IsWaiting(msg):
    author = msg["from"]["id"]
    return author in WaitingAnswer


def ValidateMsg(msg):
    response = {
        "r": {
            "a": None,
            "at": com.AnswerType.TEXT
        },
        "needWait": False
    }
    author = msg["from"]["id"]
    answerType = -1

    if "text" in msg:
        answer = msg["text"]
        if answer.startswith("/cancel"):
            del WaitingAnswer[author]
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
        del WaitingAnswer[author]
        response["r"] = {
            "a": u'CAADAQAD7wEAAs8UpQABdurS64LRGooC',
            "at": com.AnswerType.STICKER
        }
        return response
    
    return AddCustomAnswer(author, answerType, answer)

def AddCustomAnswer(author, aType, answer):
    response = {
        "r": {
            "a": None,
            "at": com.AnswerType.TEXT
        },
        "needWait": False
    }

    try:
        ca = WaitingAnswer[author]
        dao.InsertCustomAnswer(ca["regex"], aType, author, answer)
        response["r"]["a"] = "Listoooo!"
    except:
        response["r"]["a"] = "Algo no salio como esperaba, intenta de nuevo"
    
    del WaitingAnswer[author]

    return response

def ValidateCustomAnswer(msg):
    results = dao.GetCustomAnswer()
    text = msg["text"]
    for result in results:
        regex = re.compile(
            result["regex"], re.MULTILINE | re.UNICODE | re.IGNORECASE)

        if textMatch(regex, text):
            return {
                "r": {
                    "a": result["answer"],
                    "at": result["type"]
                },
                "needWait": False
            }

    return None

def textMatch(regex, test_str):
    matches = re.search(
        regex, test_str)

    return matches is not None
