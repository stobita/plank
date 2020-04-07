import React, { ReactNode, useState } from "react";
import { MoveContext } from "../context/moveContext";
import { Card, cardInit, Section } from "../model/model";

interface Props {
  children: ReactNode;
}

export interface DraggableSection extends Section {
  cards: DraggableCard[];
}

export interface DraggableCard extends Card {
  type: string;
  isDropTarget: boolean;
}

export type CardTargetPosition = {
  sectionId: number;
  index: number;
};

type DropTarget = {
  kind: "target";
};

export type DragItem = Card | DropTarget;

export const MoveContextProvider = (props: Props) => {
  // const [draggingCard, setDraggingCard] = useState<Card>(cardInit);
  // const [isDraggingCard, setIsDraggingCard] = useState(false);
  // const [cardTargetPosition, setCardTargetPosition] = useState<
  //   CardTargetPosition
  // >({ sectionId: 0, index: 0 });
  const [overedCard, setOveredCard] = useState<DragItem>(cardInit);

  // append and remove target
  // useEffect(() => {
  //   if (isDraggingCard && draggingCard) {
  //     const section = sections.find(v => v.id === draggingCard.section.id);

  //     if (!section) return;

  //     const toIndex = section.cards.findIndex(v => v.id === draggingCard.id);
  //     setCardTargetPosition({
  //       index: toIndex,
  //       sectionId: draggingCard.section.id
  //     });
  //   }

  //   // if (isDraggingCard && draggingCard) {
  //   //   const section = sections.find(v => v.id === draggingCard.section.id);
  //   //   if (!section) return;
  //   //   const toIndex = section.cards.findIndex(v => v.id === draggingCard.id);
  //   //   const cards = section.cards.filter(v => !v.isDropTarget);
  //   //   cards.splice(toIndex, 0, {
  //   //     ...cardInit,
  //   //     section: { ...sectionInit, id: draggingCard.section.id },
  //   //     type: ItemTypes.CARD,
  //   //     isDropTarget: true
  //   //   });
  //   //   setSections(prev =>
  //   //     prev
  //   //       .filter(v => v.id !== section.id)
  //   //       .concat({ ...section, cards })
  //   //       .sort((a, b) => a.id - b.id)
  //   //   );
  //   // } else {
  //   //   setSections(prev =>
  //   //     prev.map(s => ({ ...s, cards: s.cards.filter(c => !c.isDropTarget) }))
  //   //   );
  //   // }
  // }, [isDraggingCard, draggingCard]);

  // over card event
  const overCard = (card: DragItem) => {
    // const section = sections.find(v => v.id === card.section.id);

    // if (!section) return;

    // const toIndex = section.cards.findIndex(v => v.id === card.id);
    setOveredCard(card);
    //setCardTargetPosition({ index: toIndex, sectionId: card.section.id });

    // if (card.isDropTarget) return;
    // const section = sections.find(v => v.id === card.section.id);

    // if (!section) return;

    // const toIndex = section.cards.findIndex(v => v.id === card.id);
    // const cards = section.cards.filter(v => !v.isDropTarget);
    // cards.splice(toIndex, 0, {
    //   ...cardInit,
    //   section: { ...sectionInit, id: card.section.id },
    //   type: ItemTypes.CARD,
    //   isDropTarget: true
    // });
    // setSections(prev =>
    //   prev
    //     .map(s => ({ ...s, cards: s.cards.filter(c => !c.isDropTarget) }))
    //     .filter(v => v.id !== section.id)
    //     .concat({ ...section, cards })
    //     .sort((a, b) => a.id - b.id)
    // );
  };

  const dropCard = () => {
    console.log(overedCard);
    // if (!draggingCard) return;
    // const fromSection = sections.find(v => v.id === draggingCard.section.id);
    // const toSection = sections.find(v => v.cards.some(v => v.isDropTarget));
    // if (!fromSection || !toSection) return;
    // const fromIndex = fromSection.cards.findIndex(
    //   v => v.id === draggingCard.id
    // );
    // const toIndex = toSection.cards.findIndex(v => v.isDropTarget);
    // if (fromSection.id === toSection.id) {
    //   const cards = toSection.cards;
    //   const fromItem = cards[fromIndex];
    //   const nextCards = cards.slice();
    //   nextCards.splice(fromIndex, 1);
    //   nextCards.splice(toIndex, 0, fromItem);
    //   setSections(prev =>
    //     prev
    //       .map(s => ({ ...s, cards: s.cards.filter(c => !c.isDropTarget) }))
    //       .filter(v => v.id !== toSection.id)
    //       .concat({ ...toSection, cards: nextCards })
    //       .sort((a, b) => a.id - b.id)
    //   );
    // } else {
    //   const newFromSection = {
    //     ...fromSection,
    //     cards: fromSection.cards.filter(v => v.id !== draggingCard.id)
    //   };
    //   const newToSection = {
    //     ...toSection,
    //     cards: [...toSection.cards, draggingCard]
    //   };
    //   draggingCard.section = newToSection;
    //   setSections(prev =>
    //     prev
    //       .map(s => ({ ...s, cards: s.cards.filter(c => !c.isDropTarget) }))
    //       .filter(v => ![newFromSection.id, newToSection.id].includes(v.id))
    //       .concat([newToSection, newFromSection])
    //       .sort((a, b) => a.id - b.id)
    //   );
    // }
  };

  return (
    <MoveContext.Provider
      value={{
        dropCard,
        overCard,
        overedCard,
        setOveredCard
      }}
    >
      {props.children}
    </MoveContext.Provider>
  );
};
