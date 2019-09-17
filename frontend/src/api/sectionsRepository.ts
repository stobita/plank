import repository from "./repository";
import { Card } from "../model/model";

const resource = "/sections";
const childResouce = {
  cards: "cards"
};

export type CreateCardPayload = {
  name: string;
  description: string;
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
  async deleteCard(sectionId: number, cardId: number): Promise<void> {
    await repository.delete(
      `${resource}/${sectionId}/${childResouce.cards}/${cardId}`
    );
    return;
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
  }
};
