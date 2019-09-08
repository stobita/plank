import React, { useContext } from "react";
import styled from "styled-components";
import { DataContext } from "../context/dataContext";

export const BoardView = () => {
  const { sections } = useContext(DataContext);
  return (
    <Wrapper>
      {sections.map(v => (
        <div>{v.name}</div>
      ))}
    </Wrapper>
  );
};

const Wrapper = styled.div``;
