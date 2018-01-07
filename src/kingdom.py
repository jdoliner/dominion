import dominion_pb2 as pb
import pile
import card


def NewKingdom(supply):
    kingdom = pb.Kingdom()
    kingdom.supply.extend([
            pile.NewPile(card.Estate, 8),
            pile.NewPile(card.Duchy, 8),
            pile.NewPile(card.Province, 8),
            pile.NewPile(card.Copper, 46),
            pile.NewPile(card.Silver, 40),
            pile.NewPile(card.Gold, 30),
            ] + supply)
    return kingdom
