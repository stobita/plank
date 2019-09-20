import React, { useContext } from "react";
import styled from "styled-components";
import { ViewContext } from "../context/viewContext";
import { SectionController } from "./SectionController";

export const Header = () => {
  const { currentBoard } = useContext(ViewContext);
  return (
    <Wrapper>
      <CreateSection>
        {currentBoard.id !== 0 && <SectionController />}
      </CreateSection>
    </Wrapper>
  );
};

const Wrapper = styled.header`
  background: ${props => props.theme.main};
  height: 56px;
  border-bottom: 1px solid ${props => props.theme.border};
  display: flex;
`;

const CreateSection = styled.div`
  padding-left: 16px;
  display: flex;
  align-items: center;
`;
