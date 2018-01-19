import util
import random


class Player(object):
    def __init__(self, deck):
        self.deck = deck
        self.discard = []
        self.hand = []
        self.score = 0

    def __repr__(self):
        result = ""
        result += "Hand:\n"
        for card in self.hand:
            result += repr(card) + " "
        result += "\n"
        result += "Deck: %d cards\n" % len(self.hand)
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
