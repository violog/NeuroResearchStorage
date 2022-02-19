# NeuroResearchStorage
Проект для прохождения курса Blockchain от Distributed Lab

# Анализ проблемы
Представим следующую ситуацию. Несколько передовых учёных из разных мест собралось для того, чтобы исправить современное положение в области нейронаук: при огромном количестве исследований общепринятые представления о мозге в среде учёных заканчиваются на уровне функциональности нервных клеток. Причиной тому является ряд проблем академической (организованной) науки. В ходе их обсуждения выяснилось, что некоторые можно решить путём разработки специализированного программного обеспечения. Рассмотрим эти проблемы в порядке убывания их значимости для развития нейронаук.
    1. Лженаучные сведения. Это непреднамеренные ошибки или целенаправленная фальсификация. Недостаточно компетентный учёный может пропустить в публикацию сведения, которые не соответствуют научной методологии, а значит не должны претендовать на научное описание действительности.
    2. Цензура. Это абсолютно необходимый компонент, ведь нельзя публиковать лженаучные сведения. Ввиду потенциальной некомпетенции или неадекватных развитию науки целей цензора ценные исследования могут не пройти фильтр.
    3. Редактирование и удаление данных. Корректные данные, показывающие несостоятельность дальнейшего финансирования в каком-либо направлении, могут быть изменены или удалены теми, кому невыгодно прекращение финансирования.

Все эти проблемы имеют общий корень – сильная зависимость от положения учёного в иерархии. В организованной науке авторитетный учёный может заслужить свою репутацию не только своим профессионализмом, но и удачным стечением обстоятельств, навыками манипуляции, умелым использованием ненаучных приёмов. Эту зависимость необходимо свести к минимуму.

# Техническое задание
Разработать прототип информационной системы (ИС) «Хранилище исследований в области нейронаук», опираясь на соответствующий анализ проблемы. Цель ИС – помочь учёным в осуществлении добавления, хранения и учёта эмпирических сведений так, чтобы обеспечить наилучшее соответствие последних изучаемой предметной области.

Задачи ИС:
    • хранение защищённых от изменения и удаления данных в открытом доступе: кто угодно может их просматривать без регистрации
    • отправка исследований на проверку от любых узлов или прошедших процедуру регистрации (этот вопрос разрешится в процессе реализации)
    • добавление исследований, которые строго соответствуют научной методологии, путём проверки отправляемых исследований признанными в данной системе валидаторами
    • разграничение прав доступа: читать могут все, отправлять – все или только «доверенные» узлы (защита от DoS-атак), записывать – только валидаторы
    • возможность коллективно назначить пользователя валидатором, если он демонстрирует высокий уровень компетенции
    • разрешение споров, которые валидаторы не могут разрешить самостоятельно в ходе дискуссий
    • рекомендация: поощрение валидаторов и отправителей качественных данных, наказание «спамщиков» и отправителей лженаучных работ

Основополагающее и критически важное требование: информационную систему необходимо разработать так, чтобы данные в ней как можно адекватнее реальности описывали предметную область. Для этого они должны быть неизменяемы, неудаляемы и проходить процедуру проверки на соответствие научной методологии. ИС не способна проверять данные: этим занимаются люди, в техническом контексте – валидаторы. Ими должны становиться исключительно компетентные и добросовестные учёные.

Общие требования к системе: надёжность, отказоустойчивость, безопасность.

Анализируя задачи и требования, приходим к выводу, что технология блокчейн – лучшее решение для данного проекта.
