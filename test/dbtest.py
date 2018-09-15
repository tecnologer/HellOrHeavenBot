import dao

# print dao.GetStats('tecnologer')
# for x in xrange(1, 11):
#     type = dao.HEAVEN
#     if x % 2 == 0:
#         type = dao.HELL
    
#     dao.Update('tecnologer', type)

responses =[
    {"t": dao.HELL, "a": 'eso es todo, sigue asi'},
    {"t": dao.HELL, "a": 'una raya mas al tigre'},
    {"t": dao.HELL, "a": 'ha ha!'},
    {"t": dao.HELL, "a": 'anda la osa!'},
    {"t": dao.HELL, "a": 'fierro pariente... eso te espera xD'},
    {"t": dao.HELL, "a": u'\U0001f608'},
    {"t": dao.HELL, "a": u'\U0001f449\U0001f44c'},
    {"t": dao.HELL, "a": 'peleishon de bananeishon!'},
    {"t": dao.HELL, "a": 'en el infierno estaras mejor, aqui nadie te quiere'},
    {"t": dao.HELL, "a": 'los chilangos meteran tu comida favorita en bolillo'},

    {"t": dao.HEAVEN, "a": 'una de cal por las que van de arena'},
    {"t": dao.HEAVEN, "a": 'no importa,tu bien sabes a donde perteneces'},
    {"t": dao.HEAVEN, "a": 'amen!'},
    {"t": dao.HEAVEN, "a": 'ni tu te la crees'},
    {"t": dao.HEAVEN, "a": 'el pastel era mentira'},
    {"t": dao.HEAVEN, "a": 'santa no existe'},
    {"t": dao.HEAVEN, "a": u'\u2b50\ufe0f'},
    {"t": dao.HEAVEN, "a": 'te agregare uno mas porque el cura hablo bien de ti'}
]
for r in responses:
    dao.InsertAnswer(r)
# print dao.Update('tecnologer', 'hell')

print dao.GetAllStats()
