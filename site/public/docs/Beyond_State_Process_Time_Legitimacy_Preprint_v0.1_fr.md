Preprint — Position Paper — Work in Progress
Version: v0.1
Date: January 2026

Author
W.B
NOORCHAIN Research

Author Note / Disclaimer

This document is a preprint position paper exploring conceptual limits of state-based blockchain systems and proposing a process-oriented perspective on trust, time, and legitimacy.

NOORCHAIN is referenced as an experimental instantiation under active development, operating in controlled environments. This document does not constitute a technical specification, protocol implementation, roadmap, or investment proposal.

The ideas presented are intended to contribute to ongoing research and discussion at the intersection of distributed systems, cryptography, and institutional design.

Status

This document may be revised.

Feedback is welcome.

Citations should reference the version number.

1. Introduction — The State-Centric Assumption

Depuis plus d’une décennie, les blockchains ont démontré une capacité remarquable à sécuriser des états numériques partagés. Grâce à des mécanismes cryptographiques robustes et à des protocoles de consensus distribués, elles permettent à des acteurs ne se faisant pas confiance de maintenir un registre commun, vérifiable et résistant à la falsification. Cette réussite technique a ouvert la voie à de nombreux usages, allant des actifs numériques aux systèmes financiers décentralisés, en passant par des registres publics et des infrastructures applicatives globales.

Cependant, cette réussite repose sur une hypothèse implicite rarement interrogée : celle selon laquelle la vérité, la sécurité et la légitimité peuvent être entièrement exprimées et garanties à travers un état final. Dans le paradigme dominant, un système est considéré comme sûr dès lors que son état courant est valide, que ses transitions respectent des règles formelles, et que les signatures ou preuves associées sont cryptographiquement correctes. Cette approche, que l’on peut qualifier de state-based, est cohérente et efficace pour des objets purement informationnels ou financiers.

Or, lorsque les blockchains prétendent dépasser le simple enregistrement d’actifs pour s’appliquer à des systèmes humains, sociaux ou institutionnels, cette hypothèse montre rapidement ses limites. Dans ces contextes, la légitimité ne découle pas uniquement d’un résultat final, mais du processus par lequel ce résultat a été atteint. Le respect d’une procédure, la continuité dans le temps, l’absence de raccourcis opportunistes et la persistance des engagements jouent un rôle central dans la reconnaissance de ce qui est considéré comme valide ou juste.

Cette tension révèle un angle mort structurel des architectures blockchain actuelles. Elles excellent à prouver qu’un état est correct, mais peinent à démontrer qu’un chemin contraint a été suivi. Elles peuvent certifier qu’une décision a été enregistrée, mais pas nécessairement qu’elle résulte d’un processus légitime, non manipulé et conforme à des règles temporelles explicites. En d’autres termes, elles savent sécuriser des résultats, mais beaucoup moins bien des trajectoires.

Ce constat ne constitue pas une critique de la cryptographie elle-même. Les primitives cryptographiques modernes — signatures, hachages, preuves à divulgation nulle de connaissance — remplissent précisément les fonctions pour lesquelles elles ont été conçues. Elles garantissent l’intégrité, l’authenticité et la validité formelle. Le problème se situe ailleurs : dans l’écart entre ce que la cryptographie peut prouver efficacement et ce que les systèmes humains exigent pour établir la légitimité.

Dans de nombreux domaines institutionnels, la confiance ne repose pas sur un instant figé, mais sur la durée. Elle s’accumule, se maintient, et peut se perdre si une trajectoire est rompue. Elle est difficilement compressible, non accélérable, et intrinsèquement liée au temps. Pourtant, la plupart des architectures blockchain traitent le temps comme un simple paramètre technique — un coût à réduire, une latence à optimiser — plutôt que comme une contrainte fondamentale de sécurité et de légitimité.

Ce papier part de l’hypothèse que dépasser ces limites ne nécessite pas une nouvelle primitive cryptographique ni un consensus alternatif, mais un changement de cadre conceptuel. Il propose de déplacer l’attention, non plus seulement vers la validité des états, mais vers la vérifiabilité des processus et des trajectoires dans le temps. Cette approche, que nous désignons comme process-based trust, vise à réintroduire le temps, la continuité et la non-optimisabilité comme éléments centraux de la sécurité et de la légitimité dans les systèmes distribués.

À travers cette perspective, nous analysons les limites structurelles des blockchains fondées sur l’état, proposons un ensemble de propriétés attendues pour des mécanismes orientés processus, et discutons des implications en matière de gouvernance et de conception institutionnelle. Enfin, nous présentons NOORCHAIN comme une première instanciation expérimentale de ces principes, tout en reconnaissant explicitement les questions ouvertes et les défis qui restent à résoudre.

2. State-Based Truth vs Process-Based Reality

Les architectures blockchain contemporaines reposent sur une conception implicite mais structurante de la vérité : une vérité définie par un état. Un état est considéré comme valide dès lors qu’il satisfait un ensemble de règles formelles et qu’il est accepté par un mécanisme de consensus. Cette approche a fait ses preuves pour des systèmes où la valeur et la cohérence peuvent être entièrement capturées par des variables discrètes — soldes, propriétés, droits d’accès ou résultats de calcul.

Dans ce cadre, la sécurité d’un système se mesure essentiellement à la capacité de garantir que chaque transition d’état est correcte et que l’état final ne peut être altéré sans violer des hypothèses cryptographiques explicites. L’hypothèse sous-jacente est que si l’état est valide, alors le système l’est également. Cette équivalence fonctionne tant que l’on reste dans des domaines où le comment importe moins que le résultat.

Les systèmes humains et institutionnels obéissent à une logique différente. La légitimité n’y est pas uniquement attachée à un résultat final, mais à la manière dont ce résultat a été produit. Un même état peut être perçu comme légitime ou illégitime selon le chemin qui y conduit. Une décision correcte sur le plan formel peut perdre toute valeur si elle a été obtenue par contournement de procédure, par accélération artificielle du temps ou par exploitation d’asymétries d’information.

Cette distinction met en évidence une dissymétrie fondamentale : alors que les blockchains savent représenter et vérifier des états avec une grande précision, elles disposent de peu de mécanismes natifs pour représenter des trajectoires. Une trajectoire ne se réduit pas à une suite d’états successifs ; elle implique des contraintes sur l’ordre, la durée, la continuité et parfois l’irréversibilité de certaines étapes. Elle est intrinsèquement liée au temps, non comme simple index, mais comme condition de validité.

Dans la plupart des protocoles, le temps est traité comme un paramètre technique secondaire : un horodatage, un numéro de bloc, ou un intervalle à optimiser. Cette abstraction est suffisante pour garantir l’ordre des transactions, mais insuffisante pour exprimer des exigences telles que l’engagement prolongé, la persistance d’un comportement ou l’impossibilité de condenser un processus long en une action instantanée.

Cette limitation devient particulièrement visible dans la gouvernance décentralisée et la coordination collective. Les mécanismes existants tendent à réduire des processus complexes à des états agrégés : un vote final, un score, un quorum atteint. Ce faisant, ils perdent l’information essentielle sur la dynamique qui a précédé ces résultats : la stabilité des positions, la continuité de la participation, ou l’absence de manipulation opportuniste au cours du temps.

Cette observation invite à un déplacement du regard. Plutôt que de chercher à enrichir indéfiniment la représentation des états, il devient pertinent de se demander comment un système distribué pourrait attester qu’un processus contraint dans le temps a bien eu lieu, même de manière partielle ou indirecte. C’est cette transition — de la vérité fondée sur l’état vers une réalité fondée sur le processus — qui constitue le point de départ du cadre proposé dans ce travail.

Transition — On the Role and Limits of Smart Contracts

Il convient toutefois de préciser un point afin d’éviter toute ambiguïté. Les smart contracts permettent déjà d’introduire des contraintes temporelles dans les systèmes blockchain. Ils peuvent imposer des délais, ordonner des séquences d’actions, ou refuser certaines transitions d’état tant que des conditions prédéfinies ne sont pas satisfaites. À ce titre, ils jouent un rôle essentiel dans l’automatisation et l’exécution fiable de règles formelles dépendantes du temps.

Cependant, ces mécanismes restent fondamentalement limités à la vérification de conditions discrètes et instantanées, observables on-chain au moment de l’exécution. Un smart contract peut établir qu’une action est valide à un instant donné, mais il ne peut attester que cette action résulte d’un processus effectivement suivi dans la durée, sans interruption ni compression artificielle. Deux trajectoires radicalement différentes sur le plan humain peuvent produire un état final identique et parfaitement valide du point de vue du contrat.

Cette distinction est cruciale. Elle ne révèle pas une faiblesse de l’ingénierie des smart contracts, mais une limite structurelle du modèle d’évaluation qu’ils incarnent. Le temps y est traité comme un paramètre de contrôle, non comme une preuve de continuité. C’est précisément dans cet écart — entre contrainte temporelle formelle et trajectoire vécue — que se situe le problème abordé dans ce travail.

3. Why Cryptography Alone Is Not the Missing Piece

Face à ces limites, une réponse courante consiste à invoquer des mécanismes cryptographiques plus avancés. Il est tentant de considérer que les difficultés évoquées relèvent d’un déficit technique, appelant de nouvelles primitives ou des constructions plus complexes. Cette interprétation est toutefois réductrice.

La cryptographie excelle lorsqu’il s’agit de prouver des propriétés bien définies : la possession d’un secret, la validité d’un calcul, l’intégrité d’un message ou la conformité d’une transition à un ensemble de règles formelles. Ces capacités sont essentielles et constituent le socle de la sécurité des systèmes distribués modernes. Néanmoins, elles opèrent dans un cadre où les objets à prouver sont, par nature, instantanés et formalisables.

Le problème abordé ici ne relève pas d’une insuffisance de ces outils, mais d’un décalage entre ce qu’ils sont conçus pour démontrer et ce que certaines formes de légitimité exigent. Prouver qu’une signature est correcte ne permet pas d’établir qu’un engagement a été maintenu sur une période donnée. Vérifier qu’un calcul est valide n’implique pas que les étapes intermédiaires n’aient pas été contournées ou optimisées hors du regard du système.

Même les constructions cryptographiques les plus avancées peinent à capturer des propriétés telles que la non-accélérabilité, la persistance comportementale ou l’adhérence à une procédure sociale implicite. Ces propriétés ne se laissent pas aisément réduire à des énoncés statiques. Elles sont liées à la durée, à la répétition et à l’absence de discontinuité — des dimensions que la cryptographie, en tant que discipline, n’a pas vocation à représenter directement.

Il serait donc erroné de chercher une solution exclusivement dans l’ajout de nouvelles couches cryptographiques. Une telle approche risquerait de masquer le véritable enjeu : la nécessité de reconnaître que certaines formes de confiance ne peuvent être établies par la seule vérification d’états ouou de preuves ponctuelles. Elles requièrent une prise en compte explicite du processus, entendu comme une succession contrainte d’actions inscrites dans le temps.

Reconnaître cette limite ne diminue en rien la valeur de la cryptographie. Au contraire, cela permet de la repositionner à sa juste place : non comme une réponse universelle à toutes les dimensions de la confiance, mais comme un outil indispensable au service de cadres plus larges, capables d’intégrer des exigences temporelles et procédurales. C’est dans cette articulation — entre garanties cryptographiques et vérifiabilité des trajectoires — que se situe l’espace de recherche exploré dans la suite de ce travail.

4. Toward Process-Based Trust

Si les limites identifiées ne peuvent être levées ni par une validation d’état plus fine, ni par un empilement de primitives cryptographiques, alors le cadre de référence lui-même doit être déplacé. Plutôt que de chercher à exprimer toujours davantage de propriétés dans des états instantanés, il devient nécessaire de considérer la vérifiabilité des processus comme un objet de conception à part entière. C’est ce déplacement que nous désignons ici sous le terme de process-based trust.

Par process-based trust, nous entendons une forme de confiance qui ne se fonde pas uniquement sur la validité d’un résultat final, mais sur la possibilité d’établir qu’un chemin contraint a été suivi dans le temps. Ce chemin n’est pas réductible à une simple succession d’états ; il inclut des exigences sur l’ordre des actions, leur espacement temporel, leur continuité, et parfois leur irréversibilité de certaines étapes. La confiance ne découle alors pas d’un instant donné, mais de la cohérence d’une trajectoire.

Un tel cadre implique plusieurs propriétés distinctives. La première est la non-accélérabilité. Un processus digne de confiance ne doit pas pouvoir être condensé artificiellement sans en perdre la validité. Si une trajectoire peut être reproduite instantanément par préparation hors système, alors le temps n’y joue aucun rôle structurant, et la confiance associée est fragile. Dans une approche orientée processus, la durée n’est pas un paramètre optionnel, mais une contrainte constitutive.

La seconde propriété est la résistance au raccourci stratégique. De nombreux mécanismes échouent non parce qu’ils sont incorrects, mais parce qu’ils peuvent être optimisés d’une manière qui respecte la lettre des règles tout en en contournant l’esprit. Un cadre de confiance fondé sur le processus cherche précisément à limiter ce type d’optimisation, non par interdiction explicite, mais en rendant les raccourcis structurellement détectables ou non équivalents à une trajectoire conforme.

Une troisième propriété essentielle est l’auditabilité ex post. Il ne s’agit pas nécessairement de rendre chaque détail d’un processus public ou exhaustif, mais de permettre à un tiers de vérifier, après coup, que certaines contraintes ont bien été respectées dans la durée. Cette auditabilité peut être partielle, indirecte ou probabiliste, mais elle doit être suffisante pour distinguer une trajectoire conforme d’une trajectoire opportuniste produisant le même état final.

Il est important de souligner que le process-based trust ne se confond ni avec la réputation, ni avec le hasard, ni avec des mécanismes de consensus alternatifs. La réputation agrège des états passés sans garantir la continuité du comportement. Le hasard vise à imprévisibiliser, non à attester une trajectoire. Le consensus, quant à lui, cherche à synchroniser un état partagé, non à qualifier la manière dont cet état a été atteint sur le plan procédural.

Adopter une perspective orientée processus ne signifie pas que tout doit être formalisé ou mesuré. Les systèmes humains comportent des dimensions irréductibles à des règles strictes. En revanche, cela implique de reconnaître que certaines exigences — engagement dans le temps, persistance, absence de discontinuité — peuvent être partiellement capturées par des artefacts vérifiables, à condition d’être conçues explicitement comme telles.

Cette approche ouvre un espace intermédiaire entre la rigidité des preuves purement formelles et l’opacité des processus sociaux non vérifiables. Elle ne prétend pas éliminer la confiance humaine, mais en encadrer certaines dimensions critiques par des mécanismes qui rendent le temps, la continuité et la procédure observables et opposables. C’est dans cet espace que des systèmes distribués peuvent commencer à produire non seulement des résultats corrects, mais des formes de légitimité plus robustes.

La section suivante approfondit cette idée en examinant le rôle spécifique du temps, non plus comme simple paramètre technique, mais comme contrainte de sécurité à part entière, indispensable à toute tentative sérieuse de confiance fondée sur le processus.

5. Time as a First-Class Constraint

Dans la plupart des architectures blockchain, le temps est traité comme une variable auxiliaire. Il sert à ordonner des transactions, à rythmer la production de blocs ou à définir des délais minimaux avant certaines actions. Cette fonction est essentielle au bon fonctionnement du système, mais elle reste fondamentalement instrumentale. Le temps est enregistré, indexé, parfois optimisé, mais rarement considéré comme une contrainte de sécurité en soi.

Cette approche reflète une intuition profondément ancrée dans l’ingénierie des systèmes distribués : le temps est perçu comme un coût à réduire. Latence, finalité, débit et réactivité sont autant de métriques qui valorisent la rapidité et la compression temporelle. Dans ce cadre, un bon protocole est souvent celui qui parvient à produire des résultats corrects le plus vite possible, indépendamment de la durée réelle du processus ayant conduit à ces résultats.

Or, dans les systèmes humains et institutionnels, le temps joue un rôle radicalement différent. Il n’est pas seulement un support de coordination, mais une condition de légitimité. Certains engagements n’ont de valeur que s’ils sont maintenus. Certaines responsabilités ne peuvent être assumées qu’à travers la répétition et la persistance. Certaines formes de confiance disparaissent si elles peuvent être acquises instantanément, sans exposition à la durée ni au risque de rupture.

Cette dissymétrie conduit à une tension fondamentale. Les blockchains tendent à valoriser la capacité à produire un état valide rapidement, tandis que de nombreux processus sociaux exigent précisément l’inverse : l’impossibilité de raccourcir artificiellement le chemin. Lorsqu’un système permet de condenser en une transaction ce qui, sur le plan institutionnel, devrait s’étendre sur une période donnée, il affaiblit la signification même de l’engagement qu’il prétend représenter.

Traiter le temps comme une contrainte de premier ordre implique un changement de perspective. Il ne s’agit plus seulement de vérifier qu’une action se produit après un certain délai, mais de reconnaître que la durée elle-même est porteuse d’information. Un processus qui s’étend dans le temps expose un acteur à l’incertitude, à l’évolution du contexte et à la possibilité de renoncer. Cette exposition est précisément ce qui confère du poids à certains engagements.

Dans un cadre orienté processus, le temps devient ainsi non compressible par construction. Il ne peut pas être remplacé par une preuve instantanée, ni simulé intégralement hors du système. Même lorsque des éléments sont préparés à l’avance, la valeur du processus dépend du fait que certaines étapes aient été franchies à des moments distincts, sous des contraintes explicites, et sans possibilité de rétroaction opportuniste.

Il est important de noter que cette conception du temps ne suppose pas une mesure parfaite ou continue. Ce qui importe n’est pas la granularité absolue, mais l’existence de seuils temporels opposables et de transitions irréversibles. Le temps agit alors comme un mécanisme de friction structurelle, limitant la capacité d’un acteur à optimiser localement sans en subir les conséquences globales.

Réintroduire le temps comme contrainte centrale permet également de clarifier certaines ambiguïtés persistantes dans les systèmes décentralisés. Des notions telles que la loyauté, la crédibilité ou la légitimité cessent d’être de simples abstractions subjectives pour devenir liées à des trajectoires observables, même de manière imparfaite. Le temps ne garantit pas la vertu, mais il rend la simulation de la vertu plus coûteuse.

Cette approche ne remet pas en cause les objectifs d’efficacité ou de scalabilité là où ils sont pertinents. Elle suggère toutefois que tous les processus ne doivent pas être optimisés selon les mêmes critères. Pour certaines fonctions critiques — gouvernance, validation sociale, reconnaissance d’engagements — la lenteur relative n’est pas un défaut, mais une propriété de sécurité.

En traitant le temps comme une contrainte de premier ordre plutôt que comme un simple paramètre technique, les systèmes distribués peuvent commencer à exprimer des formes de confiance qui échappent aux modèles purement instantanés. Cette reconnaissance prépare le terrain pour une réflexion plus large sur la gouvernance et la réduction de l’arbitraire, abordées dans la section suivante.

6. Governance, Legitimacy, and Anti-Arbitrariness

Lorsque les systèmes blockchain abordent la gouvernance, ils déplacent implicitement la question de la sécurité vers celle de la légitimité. Il ne s’agit plus seulement de vérifier que des règles sont respectées, mais de déterminer qui définit ces règles, comment elles évoluent, et sur quelles bases elles peuvent être considérées comme justes ou acceptables. Dans ce contexte, les limites des approches purement techniques deviennent particulièrement visibles.

Une observation récurrente dans les systèmes décentralisés est que le pouvoir réel ne réside pas uniquement dans l’exécution des règles, mais dans leur paramétrage initial et leur interprétation. Choisir un seuil, un calendrier, un ordre de priorité ou un mécanisme de résolution en cas d’égalité constitue déjà un acte de gouvernance. Même lorsque ces choix sont rendus publics et audités, leur caractère discrétionnaire demeure, et avec lui le soupçon potentiel de biais ou de capture.

Les mécanismes de gouvernance on-chain tentent souvent de répondre à ce problème par le vote ou l’agrégation de préférences. Si ces outils sont utiles pour exprimer des décisions collectives, ils ne suppriment pas l’arbitraire inhérent au cadre dans lequel ces décisions sont prises. Les règles du vote, la définition des électeurs, la pondération des voix ou le moment du scrutin restent, en dernière analyse, le produit de choix humains préalables.

Dans un cadre orienté processus, la réduction de l’arbitraire ne passe pas nécessairement par davantage de participation ou de complexité procédurale, mais par la désubjectivisation de certains choix structurants. Il s’agit d’identifier les points où la décision humaine peut être remplacée par des règles impersonnelles, déterministes et vérifiables, non pour éliminer toute gouvernance, mais pour en limiter la portée discrétionnaire là où elle est la plus sensible.

Cette approche conduit à distinguer clairement deux niveaux. Le premier concerne les décisions substantives, qui relèvent inévitablement de jugements humains et de délibérations collectives. Le second concerne les mécanismes procéduraux — ordonnancement, rotation, priorisation, résolution de conflits équivalents — pour lesquels l’intervention humaine ajoute peu de valeur et introduit au contraire un risque de contestation. C’est principalement à ce second niveau que des mécanismes anti-arbitraires peuvent jouer un rôle structurant.

La légitimité procédurale ainsi recherchée ne repose pas sur l’idée que les règles seraient parfaites ou incontestables, mais sur le fait qu’elles sont appliquées sans choix opportuniste. Lorsqu’un résultat défavorable peut être attribué à une règle connue et impersonnelle plutôt qu’à une décision discrétionnaire, le système gagne en acceptabilité, même en l’absence de consensus total sur le fond.

Dans ce contexte, les références publiques déterministes — qu’elles soient mathématiques, protocolaires ou historiques — peuvent servir de supports à des décisions non négociables. Leur intérêt ne réside pas dans leur caractère aléatoire ou optimal, mais dans leur capacité à rendre certains choix non appropriables. En retirant aux acteurs la possibilité de sélectionner ex ante le cadre qui leur serait favorable, ces mécanismes réduisent l’espace de la manipulation stratégique.

Il est important de souligner que cette logique n’élimine ni le conflit ni la nécessité de la gouvernance. Elle vise plutôt à déplacer la frontière entre ce qui relève de la décision collective et ce qui peut être traité comme une contrainte institutionnelle stable. En ce sens, l’anti-arbitraire ne constitue pas une fin en soi, mais une condition permettant aux processus de gouvernance de se déployer sur un terrain plus lisible et moins contestable.

En articulant gouvernance, légitimité et contraintes procédurales non discrétionnaires, les systèmes distribués peuvent commencer à produire des formes de coordination plus robustes que celles fondées uniquement sur l’agrégation d’états ou de préférences instantanées. Cette articulation prépare l’analyse du cas de NOORCHAIN, présenté dans la section suivante comme une instanciation précoce — et volontairement limitée — de ces principes.

7. NOORCHAIN as an Early Instantiation

Au moment de la rédaction, NOORCHAIN demeure un système expérimental en cours de développement, opéré dans des environnements contrôlés et à portée limitée. Il ne s’agit ni d’un réseau largement déployé, ni d’une infrastructure stabilisée en production.

Les considérations développées jusqu’ici restent largement conceptuelles. Elles visent à clarifier un espace de problématiques plutôt qu’à prescrire une architecture définitive. Dans ce contexte, NOORCHAIN est présenté non comme une solution achevée, mais comme une instanciation précoce et partielle de certains principes associés à une approche orientée processus.

NOORCHAIN repose sur une séparation explicite entre deux dimensions souvent confondues dans les architectures blockchain : la sécurité du registre et la légitimité des interactions sociales qu’il supporte. Le consensus y est traité comme un mécanisme de sécurité technique, chargé de garantir l’intégrité et la continuité de l’état du réseau. Les questions de reconnaissance, de contribution et de gouvernance sociale sont abordées dans une couche distincte, volontairement découplée du consensus.

Cette couche, désignée comme Proof of Signal Social (PoSS), s’appuie sur des signaux émis dans le temps, agrégés sous forme de snapshots périodiques. Ces snapshots ne visent pas à produire une vérité instantanée, mais à refléter une dynamique temporelle : participation récurrente, continuité des engagements, et validation par des acteurs identifiés comme curators. La valeur accordée à un signal dépend ainsi moins de son existence ponctuelle que de sa persistance et de sa cohérence au fil des périodes.

Le rôle des curators illustre cette orientation. Leur fonction ne se limite pas à valider des données, mais à attester que certains processus ont été suivis conformément à des règles explicites. Cette attestation est inscrite dans le temps, renouvelée et exposée au risque de retrait ou de contestation. La légitimité associée à ce rôle ne peut être acquise instantanément ; elle repose sur une trajectoire observable, même si celle-ci reste imparfaite et sujette à interprétation.

Il est important de souligner que NOORCHAIN ne prétend pas résoudre l’ensemble des défis liés à la vérifiabilité des processus humains. Les mécanismes mis en œuvre restent nécessairement approximatifs et dépendants de choix de conception spécifiques. De nombreux aspects — notamment la mesure fine de la non-accélérabilité ou la formalisation complète des trajectoires — demeurent ouverts et font l’objet d’hypothèses plutôt que de garanties strictes.

Cette limitation est assumée. NOORCHAIN se positionne comme un terrain d’expérimentation réel, permettant d’explorer comment des contraintes temporelles, des rôles persistants et des procédures non discrétionnaires peuvent être combinés dans un système distribué opérationnel. L’objectif n’est pas de fournir un modèle universel, mais de tester la plausibilité d’une approche dans des conditions concrètes.

En ce sens, NOORCHAIN illustre une thèse centrale de ce travail : la transition vers une confiance fondée sur le processus ne se décrète pas par une nouvelle primitive ou un changement de consensus, mais par une série de choix architecturaux modestes, orientés vers la durée, la continuité et la limitation des raccourcis stratégiques. Ces choix ne suppriment pas la confiance humaine ; ils cherchent à en encadrer certaines dimensions critiques de manière vérifiable.

8. What This Paper Is Not

Afin d’éviter toute interprétation excessive, il est nécessaire de préciser explicitement ce que ce travail ne cherche pas à être. Cette clarification fait partie intégrante de la démarche proposée, dans la mesure où les sujets abordés sont souvent associés à des attentes techniques, économiques ou normatives qui dépassent leur portée réelle.

Ce papier ne présente pas une nouvelle primitive cryptographique. Il ne propose ni algorithme inédit, ni construction formelle destinée à remplacer les mécanismes existants de signature, de hachage ou de consensus. Les outils cryptographiques actuels sont considérés comme adéquats pour les fonctions qu’ils remplissent, et aucune hypothèse de rupture à ce niveau n’est avancée.

Il ne s’agit pas non plus d’un mécanisme de consensus alternatif. Le cadre développé ici ne vise pas à déterminer comment un réseau parvient à un accord sur un état partagé, ni à améliorer les performances, la finalité ou la résistance aux attaques de tels mécanismes. Les questions de consensus sont volontairement laissées hors du champ principal de l’analyse.

Ce travail ne constitue pas une proposition économique ou financière. Il ne formule aucune promesse de rendement, d’incitation monétaire ou de modèle de distribution de valeur. Les considérations économiques ne sont abordées qu’indirectement, dans la mesure où elles interagissent avec des processus de gouvernance ou de reconnaissance sociale.

Il ne s’agit pas davantage d’une solution complète aux problèmes de gouvernance décentralisée. Les mécanismes évoqués ne prétendent pas éliminer le conflit, l’asymétrie d’information ou la nécessité de décisions humaines. Ils cherchent uniquement à circonscrire certains espaces où l’arbitraire procédural peut être réduit par des règles impersonnelles et vérifiables.

Enfin, ce papier ne doit pas être lu comme une feuille de route technologique ou un manifeste prescriptif. Il propose un déplacement de cadre conceptuel, destiné à éclairer des limites structurelles des systèmes actuels et à ouvrir un espace de recherche. Les choix d’implémentation, les compromis pratiques et les formes institutionnelles restent largement ouverts et dépendront des contextes spécifiques dans lesquels ces idées pourraient être explorées.

9. Open Questions and Research Directions

Le cadre proposé dans ce travail vise avant tout à identifier une limite structurelle des architectures blockchain actuelles et à esquisser une direction possible pour la dépasser. Il ne constitue ni une théorie achevée ni un ensemble de prescriptions techniques. À ce titre, de nombreuses questions fondamentales demeurent ouvertes et appellent une exploration approfondie.

Une première série de questions concerne la formalisation des trajectoires. Si la notion de processus contraint dans le temps est centrale, elle reste difficile à définir de manière rigoureuse sans la réduire à une simple succession d’états. Quels types de contraintes temporelles peuvent être exprimés de façon vérifiable sans exiger une observation exhaustive ? Comment distinguer une trajectoire conforme d’une trajectoire opportuniste lorsque les deux produisent des résultats identiques ? Ces questions touchent aux limites de ce qui peut être attesté dans un système distribué sans recourir à des hypothèses irréalistes.

Une seconde ligne de recherche porte sur la mesure de la non-optimisabilité. L’idée qu’un processus ne doit pas pouvoir être accéléré ou condensé artificiellement est conceptuellement claire, mais opérationnellement délicate. Quelles formes de raccourcis sont réellement problématiques, et lesquelles relèvent d’une optimisation légitime ? À partir de quel point un gain d’efficacité compromet-il la signification institutionnelle d’un engagement ? Répondre à ces questions suppose un dialogue entre ingénierie des systèmes, économie et sciences sociales.

La question de l’auditabilité partielle constitue un autre axe critique. Les processus humains ne peuvent ni ne doivent être entièrement rendus transparents. Il s’agit donc de déterminer quels signaux, quels artefacts ou quelles traces minimales suffisent à rendre une trajectoire vérifiable a posteriori, sans exposer inutilement les acteurs ni rigidifier excessivement les interactions. Cette problématique soulève des enjeux à la fois techniques, éthiques et institutionnels.

Un champ de recherche connexe concerne l’articulation entre couches on-chain et off-chain. Les processus orientés temps et continuité se déploient inévitablement en partie hors de la chaîne. Comment concevoir des mécanismes qui permettent de relier ces dynamiques off-chain à des garanties on-chain, sans introduire de dépendance excessive à des autorités centrales ou à des oracles opaques ? Cette articulation reste largement inexplorée dans les architectures actuelles.

La gouvernance des systèmes orientés processus soulève également des questions spécifiques. Si certaines décisions procédurales peuvent être désubjectivisées, d’autres relèvent inévitablement de choix collectifs et de compromis contextuels. Comment ces deux dimensions peuvent-elles coexister sans que l’une ne capture l’autre ? Quels mécanismes permettent d’adapter les règles procédurales sans invalider les trajectoires construites dans le temps ?

Enfin, une question plus large concerne la portée réelle de l’approche proposée. Tous les systèmes n’ont pas vocation à intégrer des contraintes temporelles fortes ou des preuves de trajectoire. Identifier les domaines où ces mécanismes apportent une valeur institutionnelle nette — par opposition à ceux où ils introduisent une complexité inutile — constitue un enjeu central pour toute exploration future.

Ces questions ne peuvent être résolues par un seul protocole ni par une discipline isolée. Elles appellent une recherche interdisciplinaire, à l’intersection des systèmes distribués, de la cryptographie appliquée, de l’économie institutionnelle et des sciences sociales. Le cadre esquissé ici se veut une contribution à cette exploration, en proposant un vocabulaire et une grille de lecture pour aborder des problématiques encore largement sous-théorisées dans l’écosystème blockchain.

La conclusion qui suit revient sur la thèse centrale de ce travail et en reformule les implications essentielles, sans prétendre en clore les débats.

10. Conclusion — Beyond State

Ce travail est parti d’un constat simple : les blockchains excellent à sécuriser des états, mais rencontrent des limites structurelles dès lors qu’elles cherchent à produire ou à soutenir des formes de légitimité qui dépendent du temps, de la continuité et du respect de processus. Cette limite n’est ni accidentelle ni le résultat d’un manque d’ingénierie. Elle découle directement d’un cadre conceptuel dans lequel la vérité est principalement conçue comme un état final vérifiable.

En distinguant explicitement la validité d’un état de la légitimité d’une trajectoire, ce papier propose un déplacement du regard. Il ne s’agit pas de remettre en cause les fondements techniques des systèmes distribués actuels, mais de reconnaître que certaines dimensions essentielles des systèmes humains — engagement, persistance, non-accélérabilité — ne peuvent être adéquatement capturées par des mécanismes exclusivement orientés vers l’instantané.

La notion de process-based trust introduite ici vise à fournir un cadre pour penser cette lacune. Elle suggère que la confiance, dans certains contextes critiques, ne peut émerger que si le temps est traité comme une contrainte de premier ordre, et non comme un simple paramètre technique. Dans cette perspective, la durée devient porteuse d’information, et la trajectoire suivie acquiert une valeur propre, distincte du résultat qu’elle produit.

Cette approche n’implique pas l’abandon des outils existants. La cryptographie, les smart contracts et les mécanismes de consensus conservent un rôle central dans la sécurisation des systèmes distribués. Toutefois, leur efficacité dépend du cadre dans lequel ils sont mobilisés. Lorsqu’ils sont utilisés pour attester des processus plutôt que des seuls états, leur fonction se transforme : ils ne garantissent plus uniquement la correction formelle, mais contribuent à rendre certaines formes de comportement opposables et auditables dans le temps.

L’examen de NOORCHAIN comme instanciation précoce illustre à la fois le potentiel et les limites de cette orientation. Il montre que la formalisation de processus humains ne vise pas l’exhaustivité, mais la mise sous contrainte de dimensions critiques — notamment le temps, la persistance et l’absence de raccourci stratégique — qui échappent aux modèles fondés exclusivement sur l’état.

Au-delà de ce cas particulier, l’ambition de ce travail est plus modeste et plus large à la fois. Il ne propose pas une solution universelle, mais un cadre conceptuel destiné à clarifier un espace de recherche encore largement inexploré. En ce sens, il invite à repenser ce que les systèmes blockchain cherchent à prouver, ce qu’ils peuvent raisonnablement garantir, et ce qui doit rester du ressort de la délibération humaine.

Aller au-delà de l’état ne signifie pas rejeter les acquis des architectures existantes, mais reconnaître leurs frontières. C’est en prenant ces frontières au sérieux — plutôt qu’en tentant de les masquer par une complexité croissante — que les systèmes distribués pourront évoluer vers des formes de coordination plus adaptées aux réalités institutionnelles et sociales qu’ils aspirent à soutenir.
