import React, { useContext, useState } from "react";
import styled from "styled-components";
import { EventContext } from "../context/eventContext";
import { ReactComponent as MoreIconImage } from "../assets/more.svg";
import { ReactComponent as SettingIconImage } from "../assets/setting.svg";
import { ViewContext } from "../context/viewContext";
import { Board } from "../model/model";
import { DataContext } from "../context/dataContext";

interface Props {
  item: Board;
}
export const BoardListItem = (props: Props) => {
  const { item } = props;
  const { boards } = useContext(DataContext);
  const { setCurrentBoard, currentBoard, setBoardSettingActive } = useContext(
    ViewContext
  );
  const { updatedBoardIds } = useContext(EventContext);
  const [menuActive, setMenuActive] = useState(false);
  const handleOnClickItem = (id: number) => {
    const board = boards.find(v => v.id === id);
    setCurrentBoard(board || { id: 0, name: "" });
  };
  const handleOnClickMore = () => {
    setMenuActive(prev => !prev);
  };
  const selected = currentBoard.id === item.id;
  const handleOnClickSetting = () => {
    setBoardSettingActive(prev => !prev);
  };
  return (
    <Wrapper key={item.id}>
      <Main
        onClick={() => handleOnClickItem(item.id)}
        selected={selected}
        updated={!!updatedBoardIds.find(id => id === item.id)}
      >
        {item.name}
      </Main>
      {selected && (
        <Right>
          <Menu active={menuActive}>
            {menuActive && (
              <SettingIcon onClick={handleOnClickSetting}></SettingIcon>
            )}
          </Menu>
          <MoreIcon onClick={handleOnClickMore}></MoreIcon>
        </Right>
      )}
    </Wrapper>
  );
};

const Wrapper = styled.li`
  display: flex;
  height: 32px;
  justify-content: space-between;
  align-items: center;
`;

const Main = styled.div<{ selected: boolean; updated: boolean }>`
  font-weight: bold;
  color: ${props =>
    props.selected
      ? props.theme.solid
      : props.updated
      ? props.theme.primary
      : props.theme.weak};
  font-size: 1.1rem;
  cursor: pointer;
`;

const Right = styled.div`
  display: flex;
`;

const Menu = styled.div<{ active: boolean }>`
  display: flex;
  transition: 0.5s;
  max-width: ${props => (props.active ? 128 : 0)}px;
  overflow: hidden;
`;

const MoreIcon = styled(MoreIconImage)`
  fill: ${props => props.theme.weak};
  height: 24px;
  cursor: pointer;
`;

const SettingIcon = styled(SettingIconImage)`
  cursor: pointer;
  fill: ${props => props.theme.solid};
  height: 24px;
  margin-right: 8px;
`;
