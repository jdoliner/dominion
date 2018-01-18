from enum import Enum


class Type(Enum):
    Victory = 0
    Treasure = 1
    Action = 2
    Curse = 4


class Card(object):
    def __init__(self, name, type, cost, play, score):
        self.name = name
        self.type = type
        self.cost = cost
        self.play = play
        self.score = score

    def __repr__(self):
        return self.name


def PlusCoin(amount, game):
    game.coin += amount
    return game


def Coin(amount):
    return lambda game: PlusCoin(amount, game)


def VP(amount):
    return lambda game, owner: amount


# Victory
Estate = Card("Estate", [Type.Victory], 2, None, VP(1))
Duchy = Card("Duchy", [Type.Victory], 2, None, VP(1))
Province = Card("Province", [Type.Victory], 2, None, VP(1))

# Treasure
Copper = Card("Copper", [Type.Treasure], 0, Coin(1), None)
Silver = Card("Silver", [Type.Treasure], 3, Coin(2), None)
Gold = Card("Gold", [Type.Treasure], 6, Coin(3), None)
