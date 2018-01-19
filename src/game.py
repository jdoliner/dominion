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
        self.players = [
                player.Player([card.Estate]*3 + [card.Copper]*7),
                player.Player([card.Estate]*3 + [card.Copper]*7)]
        for p in self.players:
            p.draw(5)
        self.play = []
        self.trash = []
        self.setup_turn(0)

    def __repr__(self):
        result = ""
        result += "Kingdom:\n" + str(self.kingdom) + "\n"
        result += "Play:\n" + str(self.play) + "\n"
        result += "Trash:\n" + str(self.trash) + "\n"
        for i, p in enumerate(self.players):
            result += "Player %d:\n" % i
            result += str(p) + "\n"
        return result

    def setup_turn(self, player):
        self.turn_player = player
        self.active_player = player
        self.phase = Phase.Action
        self.actions = 1
        self.coin = 0
        self.buys = 1

    def cleanup_turn(self):
        self.players[player].cleanup(self.play)


g = Game([])
print g
