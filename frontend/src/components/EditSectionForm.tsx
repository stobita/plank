import React, { useContext } from "react";
import { Section } from "../model/model";
import { Input } from "./Input";
import { useForm } from "../hooks/useForm";
import boardsRepository, {
  UpdateSectionPayload
} from "../api/boardsRepository";
import { DataContext } from "../context/dataContext";

interface Props {
  item: Section;
  afterSubmit: () => void;
}
export const EditSectionForm = (props: Props) => {
  const { item } = props;
  const { setSections } = useContext(DataContext);
  const updateSection = async () => {
    await boardsRepository.updateSection(item.board.id, item.id, formValue);
    const current = await boardsRepository.getBoardSections(item.board.id);
    setSections(current);
    props.afterSubmit();
  };
  const { formValue, handleOnChangeInput, handleOnSubmit } = useForm<
    UpdateSectionPayload
  >({ name: item.name }, updateSection);
  return (
    <form onSubmit={handleOnSubmit}>
      <Input
        name="name"
        value={formValue.name}
        type="text"
        placeholder="name"
        onChange={handleOnChangeInput}
        onBlur={props.afterSubmit}
      ></Input>
    </form>
  );
};
