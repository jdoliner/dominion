import dominion_pb2 as pb

playFs = {}
scoreFs = {}


def NewCard(name, type, cost, play, score):
    card = pb.Card()
    card.name = name
    card.type.extend(type)
    card.cost = cost
    playFs[name] = play
    scoreFs[name] = score
    return card


def PlusCoin(amount, game):
    game.coin += amount
    return game


def Coin(amount):
    return lambda game: PlusCoin(amount, game)


def VP(amount):
    return lambda game, owner: amount


# Victory
Estate = NewCard("Estate", [pb.Card.Victory], 2, None, VP(1))
Duchy = NewCard("Duchy", [pb.Card.Victory], 2, None, VP(1))
Province = NewCard("Province", [pb.Card.Victory], 2, None, VP(1))

# Treasure
Copper = NewCard("Copper", [pb.Card.Treasure], 0, Coin(1), None)
Silver = NewCard("Silver", [pb.Card.Treasure], 3, Coin(2), None)
Gold = NewCard("Gold", [pb.Card.Treasure], 6, Coin(3), None)
