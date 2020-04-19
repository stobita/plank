import React, { useState, useContext } from "react";
import { ReactComponent as MoreIconImage } from "../assets/more.svg";
import { ReactComponent as DeleteIconImage } from "../assets/trash.svg";
import { AddButton } from "./AddButton";
import styled from "styled-components";
import { CreateCardForm } from "./CreateCardForm";
import { Section } from "../model/model";
import { DeleteConfirmation } from "./DeleteConfirmation";
import boardsRepository from "../api/boardsRepository";
import { EditSectionForm } from "./EditSectionForm";
import { DataContext } from "../context/dataContext";

interface Props {
  section: Section;
}

export const SectionPanelHead = (props: Props) => {
  const { setSections } = useContext(DataContext);
  const [formActive, setFormActive] = useState(false);
  const [deleteActive, setDeleteActive] = useState(false);
  const [editActive, setEditActive] = useState(false);
  const [more, setMore] = useState(false);
  const handleOnClickAddButton = () => {
    if (deleteActive) {
      setDeleteActive(false);
    }
    setFormActive((prev) => !prev);
  };
  const handleAfterCreateCard = () => {
    setFormActive(false);
  };
  const handleOnClickMore = () => {
    setMore((prev) => !prev);
  };
  const handleOnClickDelete = () => {
    if (formActive) {
      setFormActive(false);
    }
    setDeleteActive((prev) => !prev);
  };
  const handleOnSubmitDelete = async () => {
    await boardsRepository.deleteSection(
      props.section.board.id,
      props.section.id
    );
    const current = await boardsRepository.getBoardSections(
      props.section.board.id
    );
    setSections(current);
  };
  const handleOnClickDeleteCancel = () => {
    setDeleteActive(false);
  };
  const handleOnClickSectionName = () => {
    setEditActive(true);
  };
  const afterEditSection = () => {
    setEditActive(false);
  };
  return (
    <Wrapper>
      <Top>
        {editActive ? (
          <EditSectionForm
            item={props.section}
            afterSubmit={afterEditSection}
          ></EditSectionForm>
        ) : (
          <Name onClick={handleOnClickSectionName}>{props.section.name}</Name>
        )}
        <HeadRight>
          <Menu active={more}>
            <DeleteIcon onClick={handleOnClickDelete}></DeleteIcon>
          </Menu>
          <MoreIcon onClick={handleOnClickMore}></MoreIcon>
          <AddButton
            onClick={handleOnClickAddButton}
            isClose={formActive}
          ></AddButton>
        </HeadRight>
      </Top>
      <FormArea active={formActive}>
        {formActive && (
          <FormAreaInner>
            <CreateCardForm
              section={props.section}
              afterSubmit={handleAfterCreateCard}
            />
          </FormAreaInner>
        )}
      </FormArea>
      <DeleteArea active={deleteActive}>
        <DeleteConfirmation
          onSubmit={handleOnSubmitDelete}
          onClickCancel={handleOnClickDeleteCancel}
        ></DeleteConfirmation>
      </DeleteArea>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
`;

const Top = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

const HeadRight = styled.div`
  display: flex;
  align-items: center;
`;

const Menu = styled.div<{ active: boolean }>`
  transition: 0.5s;
  max-width: ${(props) => (props.active ? 128 : 0)}px;
  overflow: hidden;
`;

const Name = styled.h3`
  margin: 0;
  font-size: 1.1rem;
  cursor: pointer;
`;

const MoreIcon = styled(MoreIconImage)`
  fill: ${(props) => props.theme.weak};
  height: 24px;
  margin-right: 8px;
  cursor: pointer;
`;

const FormArea = styled.div<{ active: boolean }>`
  transition: 0.5s;
  /** max-height: ${(props) => (props.active ? 180 : 0)}px; */
  overflow: hidden;
`;

const FormAreaInner = styled.div`
  margin-bottom: 4px;
`;

const DeleteArea = styled.div<{ active: boolean }>`
  transition: 0.5s;
  max-height: ${(props) => (props.active ? 160 : 0)}px;
  overflow: hidden;
`;

const DeleteIcon = styled(DeleteIconImage)`
  fill: ${(props) => props.theme.solid};
  height: 24px;
  margin-right: 8px;
  cursor: pointer;
`;
