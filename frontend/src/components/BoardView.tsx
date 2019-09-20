import React, { useContext } from "react";
import styled from "styled-components";
import { DataContext } from "../context/dataContext";
import { SectionPanel } from "./SectionPanel";

export const BoardView = () => {
  const { sections } = useContext(DataContext);
  return (
    <Wrapper>
      <Main>
        {sections.map(v => (
          <SectionPanel section={v} key={v.id} />
        ))}
      </Main>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  flex: 1;
`;

const Main = styled.div`
  display: flex;
`;
