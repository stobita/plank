import React, { useContext } from "react";
import styled from "styled-components";
import { DataContext } from "../context/dataContext";
import { SectionPanel } from "./SectionPanel";

export const BoardView = () => {
  const { sections } = useContext(DataContext);
  return (
    <Wrapper>
      {sections.map(v => (
        <SectionPanel section={v} key={v.id} />
      ))}
    </Wrapper>
  );
};

const Wrapper = styled.div`
  display: flex;
`;
