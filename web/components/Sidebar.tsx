import React, { useState, } from "react";
import styled from "styled-components";
import { BoardList } from "./BoardList";
import SettingIconImage from "../assets/setting.svg";
import { Setting } from "./Setting";

export const Sidebar = () => {
  const [settingActive, setSettingActive] = useState(false);

  const handleOnClickSetting = () => {
    setSettingActive((prev) => !prev);
  };
  return (
    <Wrapper>
      <Top>
        <Title>plank</Title>
      </Top>
      <Middle>
        <BoardList />
      </Middle>
      <Bottom>
        {settingActive && <Setting />}
        <Icon>
          <SettingIcon onClick={handleOnClickSetting} />
        </Icon>
      </Bottom>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  background: ${(props) => props.theme.main};
  display: flex;
  flex-direction: column;
  flex: 1;
  border-right: 1px solid ${(props) => props.theme.border};
`;

const Top = styled.div`
  height: 64px;
  display: flex;
  align-items: center;
  padding: 0 16px;
`;

const Middle = styled.div`
  padding: 8px 16px;
`;

const Bottom = styled.div`
  display: flex;
  flex-direction: column;
  padding: 16px 16px;
  margin-top: auto;
`;

const Title = styled.h1`
  color: ${(props) => props.theme.solid};
  margin: 0;
`;

const Icon = styled.div`
  display: flex;
  justify-content: space-between;
`;

const SettingIcon = styled(SettingIconImage)`
  cursor: pointer;
  fill: ${(props) => props.theme.solid};
  height: 36px;
`;
