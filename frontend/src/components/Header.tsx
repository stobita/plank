import React, { useContext } from "react";
import styled from "styled-components";
import { AddButton } from "./AddButton";
import { ViewContext } from "../context/viewContext";
import { CreateSectionForm } from "./CreateSectionForm";

export const Header = () => {
  const {
    setCreateSectionActive,
    createSectionActive,
    currentBoardId
  } = useContext(ViewContext);
  const handleOnClickAddSectionButton = () => {
    setCreateSectionActive(true);
  };
  return (
    <Wrapper>
      <CreateSection>
        {currentBoardId !== 0 && (
          <>
            {!createSectionActive && (
              <>
                <AddButton onClick={handleOnClickAddSectionButton} />
                <span>Create Section</span>
              </>
            )}
            {createSectionActive && <CreateSectionForm />}
          </>
        )}
      </CreateSection>
    </Wrapper>
  );
};

const Wrapper = styled.header`
  height: 56px;
  border-bottom: 1px solid ${props => props.theme.border};
  display: flex;
`;

const CreateSection = styled.div`
  padding-left: 16px;
  display: flex;
  align-items: center;
`;
