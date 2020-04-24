import repository from "./repository";
import { Card } from "../model/model";

const resource = "/sections";
const childResouce = {
  cards: "cards",
};

export type CreateCardPayload = {
  name: string;
  description: string;
  labels: string[];
  limitTime?: number;
};

export type MoveCardPayload = {
  prevCardId: number | null;
  targetSectionId: number | null;
};

export type UpdateCardPayload = CreateCardPayload;

export default {
  async createCard(
    sectionId: number,
    payload: CreateCardPayload
  ): Promise<Card> {
    const res = await repository.post<Card>(
      `${resource}/${sectionId}/${childResouce.cards}`,
      payload
    );
    return res.data;
  },
  async updateCard(
    sectionId: number,
    cardId: number,
    payload: UpdateCardPayload
  ): Promise<Card> {
    const res = await repository.put(
      `${resource}/${sectionId}/${childResouce.cards}/${cardId}`,
      payload
    );
    return res.data;
  },
  async deleteCard(sectionId: number, cardId: number): Promise<void> {
    await repository.delete(
      `${resource}/${sectionId}/${childResouce.cards}/${cardId}`
    );
    return;
  },
  async reorderCardPosition(
    sectionId: number,
    cardId: number,
    position: number
  ): Promise<void> {
    await repository.put(
      `${resource}/${sectionId}/${childResouce.cards}/${cardId}/reorder`,
      { position }
    );
  },
  async moveCard(
    sectionId: number,
    cardId: number,
    position: number,
    destinationSectionId: number
  ): Promise<void> {
    await repository.put(
      `${resource}/${sectionId}/${childResouce.cards}/${cardId}/move`,
      { position, destinationSectionId }
    );
  },
};
