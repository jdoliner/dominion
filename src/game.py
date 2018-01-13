from enum import Enum
import kingdom
import player
import card


class Phase(Enum):
    Action = 0
    Buy = 1
    Cleanup = 2


class Game(object):
    def __init__(self, supply):
        self.kingdom = kingdom.Kingdom(supply)
        self.players = [player.Player([card.Estate]*3 + [card.Copper]*7)]*2
        self.play = []
        self.trash = []
        self.turn = 0
        self.active = 0
        self.phase = Phase.Action
        self.actions = 1
        self.coin = 0
        self.buys = 1


g = Game([])
print g
