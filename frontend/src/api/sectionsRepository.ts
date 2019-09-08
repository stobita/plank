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

export default {
  async CreateCardForm(
    sectionId: number,
    payload: CreateCardPayload
  ): Promise<Card> {
    const res = await repository.post<Card>(
      `${resource}/${sectionId}/${childResouce.cards}`,
      payload
    );
    return res.data;
  }
};
