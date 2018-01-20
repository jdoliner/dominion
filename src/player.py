import util
import random
import card


class Player(object):
    def __init__(self):
        self.deck = []
        self.discard = [card.Estate, card.Estate, card.Estate,
                        card.Copper, card.Copper, card.Copper, card.Copper,
                        card.Copper, card.Copper, card.Copper]
        self.hand = []
        self.score = 0

    def __str__(self):
        result = ""
        result += "Hand:\n"
        for c in self.hand:
            result += str(c) + " "
        result += "\n"
        result += "Deck: %d cards\n" % len(self.deck)
        result += "Discard: %d cards\n" % len(self.discard)
        return result

    def cleanup(self, play):
        util.move_list(play, self.discard)
        util.move_list(self.hand, self.discard)
        self.draw(5)

    def shuffle(self):
        util.move_list(self.discard, self.deck)
        random.shuffle(self.deck)

    def draw(self, count):
        if count == 0:
            return
        if not self.deck:
            if not self.discard:
                return
            self.shuffle()
        self.hand.append(self.deck.pop())
        self.draw(count - 1)
