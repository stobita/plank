import React, { useContext } from "react";
import { AddButton } from "./AddButton";
import { ViewContext } from "../context/viewContext";
import { CreateSectionForm } from "./CreateSectionForm";
import styled from "styled-components";

export const SectionController = () => {
  const { setCreateSectionActive, createSectionActive } = useContext(
    ViewContext
  );
  const handleOnClickAddSectionButton = () => {
    setCreateSectionActive(prev => !prev);
  };
  return (
    <>
      <AddButton
        onClick={handleOnClickAddSectionButton}
        isClose={createSectionActive}
      />
      <LabelArea formActive={createSectionActive}>
        <ButtonLabel>Create Section</ButtonLabel>
      </LabelArea>
      <FormArea formActive={createSectionActive}>
        <CreateSectionForm />
      </FormArea>
    </>
  );
};

const LabelArea = styled.div<{ formActive: boolean }>`
  transition: 0.2s;
  max-width: ${props => (props.formActive ? "0" : "500px")};
  transition-delay: ${props => (props.formActive ? 0 : 0.3)}s;
  overflow: hidden;
`;

const ButtonLabel = styled.span`
  font-weight: bold;
  padding-left: 12px;
  white-space: nowrap;
`;

const FormArea = styled.div<{ formActive: boolean }>`
  transition: 0.2s;
  transition-delay: ${props => (props.formActive ? 0.3 : 0)}s;
  max-width: ${props => (props.formActive ? "500px" : "0")};
  overflow: hidden;
`;
