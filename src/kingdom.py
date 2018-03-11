import card


class Kingdom(object):
    def __init__(self, supply):
        self.supply = [
            Pile(card.Estate, 8),
            Pile(card.Duchy, 8),
            Pile(card.Province, 8),
            Pile(card.Copper, 46),
            Pile(card.Silver, 40),
            Pile(card.Gold, 30),
            ] + supply

    def __str__(self):
        result = ""
        for pile in self.supply:
            if len(pile) == 0:
                result += "Empty\n"
            else:
                result += str(pile[0]) + " - " + "%d\n" % len(pile)
        return result


def Pile(cardF, count):
    result = []
    for i in range(count):
        result.append(cardF())
    return result
