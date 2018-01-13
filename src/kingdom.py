import pile
import card


class Kingdom(object):
    def __init__(self, supply):
        self.supply = [
            pile.Pile(card.Estate, 8),
            pile.Pile(card.Duchy, 8),
            pile.Pile(card.Province, 8),
            pile.Pile(card.Copper, 46),
            pile.Pile(card.Silver, 40),
            pile.Pile(card.Gold, 30),
            ] + supply
