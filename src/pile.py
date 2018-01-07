import dominion_pb2 as pb


def NewPile(card, count):
    pile = pb.Pile()
    pile.card.CopyFrom(card)
    pile.count = count
    return pile
