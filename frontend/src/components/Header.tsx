import React from "react";
import styled from "styled-components";

export const Header = () => {
  return <Wrapper></Wrapper>;
};

const Wrapper = styled.header`
  height: 56px;
  border-bottom: 1px solid ${props => props.theme.border};
`;
