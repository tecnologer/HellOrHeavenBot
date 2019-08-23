import dao
import com
# print dao.GetStats('tecnologer')
# for x in xrange(1, 11):
#     type = dao.HEAVEN
#     if x % 2 == 0:
#         type = dao.HELL

#     dao.Update('tecnologer', type)

""" proposals =[
    {"t": dao.HELL, "at": com.Answerype.TEXT, "a": 'eso es todo, sigue asi'}
    # {"t": dao.HEAVEN, "at": com.Answerype.GIF,"a": 'CgADAQADAQADLm_4TFkwvxivN4ncAg'},
]
for r in proposals:
    dao.InsertProposal(r) """
# print dao.Update('tecnologer', 'hell')

prop = dao.GetRandomProposal('10244644')

if prop != "":
    dao.UpdateScore('10244644', prop, True)

    prop = dao.GetRandomProposal('10244644')

if prop == "":
    print("No hay")
