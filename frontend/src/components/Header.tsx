import React from "react";
import styled from "styled-components";
import { ReactComponent as SettingIconImage } from "../assets/setting.svg";

export const Header = () => {
  return (
    <Wrapper>
      <Title>plank</Title>
      <SettingIcon></SettingIcon>
    </Wrapper>
  );
};

const Wrapper = styled.header`
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 60px;
  border-bottom: 1px solid ${props => props.theme.border};
  text-align: center;
  padding: 0 16px;
`;

const Title = styled.h1`
  color: ${props => props.theme.text};
`;
const SettingIcon = styled(SettingIconImage)`
  cursor: pointer;
  fill: ${props => props.theme.text};
  height: 36px;
`;
