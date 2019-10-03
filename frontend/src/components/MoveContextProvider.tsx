import React, { ReactNode, useState, useContext, useEffect } from "react";
import { MoveContext } from "../context/moveContext";
import { Card, cardInit, sectionInit, Section } from "../model/model";
import { DataContext } from "../context/dataContext";
import { ItemTypes } from "../constants/dnd";

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

export const MoveContextProvider = (props: Props) => {
  const { sections: dataSection } = useContext(DataContext);
  const [sections, setSections] = useState<DraggableSection[]>([]);
  const [draggingCard, setDraggingCard] = useState<DraggableCard>();
  const [isDraggingCard, setIsDraggingCard] = useState(false);

  // data update
  useEffect(() => {
    const data = dataSection.map(s => ({
      ...s,
      cards: s.cards.map<DraggableCard>(c => ({
        ...c,
        isDropTarget: false,
        type: ItemTypes.CARD
      }))
    }));
    setSections(data);
  }, [dataSection]);

  // append and remove target
  useEffect(() => {
    if (isDraggingCard && draggingCard) {
      const section = sections.find(v => v.id === draggingCard.section.id);
      if (!section) return;
      const toIndex = section.cards.findIndex(v => v.id === draggingCard.id);
      const cards = section.cards.filter(v => !v.isDropTarget);
      cards.splice(toIndex, 0, {
        ...cardInit,
        section: { ...sectionInit, id: draggingCard.section.id },
        type: ItemTypes.CARD,
        isDropTarget: true
      });
      setSections(prev =>
        prev
          .filter(v => v.id !== section.id)
          .concat({ ...section, cards })
          .sort((a, b) => a.id - b.id)
      );
    } else {
      setSections(prev =>
        prev.map(s => ({ ...s, cards: s.cards.filter(c => !c.isDropTarget) }))
      );
    }
  }, [isDraggingCard, draggingCard]);

  // over card event
  const overCard = (card: DraggableCard) => {
    if (card.isDropTarget) return;
    const section = sections.find(v => v.id === card.section.id);

    if (!section) return;

    const toIndex = section.cards.findIndex(v => v.id === card.id);
    const cards = section.cards.filter(v => !v.isDropTarget);
    cards.splice(toIndex, 0, {
      ...cardInit,
      section: { ...sectionInit, id: card.section.id },
      type: ItemTypes.CARD,
      isDropTarget: true
    });
    setSections(prev =>
      prev
        .map(s => ({ ...s, cards: s.cards.filter(c => !c.isDropTarget) }))
        .filter(v => v.id !== section.id)
        .concat({ ...section, cards })
        .sort((a, b) => a.id - b.id)
    );
  };

  const dropCard = () => {
    if (!draggingCard) return;
    const fromSection = sections.find(v => v.id === draggingCard.section.id);
    const toSection = sections.find(v => v.cards.some(v => v.isDropTarget));
    if (!fromSection || !toSection) return;
    const fromIndex = fromSection.cards.findIndex(
      v => v.id === draggingCard.id
    );
    const toIndex = toSection.cards.findIndex(v => v.isDropTarget);
    if (fromSection.id === toSection.id) {
      const cards = toSection.cards;
      const fromItem = cards[fromIndex];
      const nextCards = cards.slice();
      nextCards.splice(fromIndex, 1);
      nextCards.splice(toIndex, 0, fromItem);
      setSections(prev =>
        prev
          .map(s => ({ ...s, cards: s.cards.filter(c => !c.isDropTarget) }))
          .filter(v => v.id !== toSection.id)
          .concat({ ...toSection, cards: nextCards })
          .sort((a, b) => a.id - b.id)
      );
    } else {
      const newFromSection = {
        ...fromSection,
        cards: fromSection.cards.filter(v => v.id !== draggingCard.id)
      };
      const newToSection = {
        ...toSection,
        cards: [...toSection.cards, draggingCard]
      };
      draggingCard.section = newToSection;
      setSections(prev =>
        prev
          .map(s => ({ ...s, cards: s.cards.filter(c => !c.isDropTarget) }))
          .filter(v => ![newFromSection.id, newToSection.id].includes(v.id))
          .concat([newToSection, newFromSection])
          .sort((a, b) => a.id - b.id)
      );
    }
  };

  return (
    <MoveContext.Provider
      value={{
        draggingCard,
        setDraggingCard,
        dropCard,
        overCard,
        sections,
        isDraggingCard,
        setIsDraggingCard
      }}
    >
      {props.children}
    </MoveContext.Provider>
  );
};
