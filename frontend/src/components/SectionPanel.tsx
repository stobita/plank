import React, { useState } from "react";
import styled from "styled-components";
import { Section } from "../model/model";
import { AddButton } from "./AddButton";
import { CreateCardForm } from "./CreateCardForm";
import { CardPanel } from "../api/CardPanel";

interface Props {
  section: Section;
}

export const SectionPanel = (props: Props) => {
  const [formActive, setFormActive] = useState(false);
  const handleOnClickAddButton = () => {
    setFormActive(prev => !prev);
  };
  return (
    <Wrapper>
      <Inner>
        <Head>
          <HeadRow>
            <Name>{props.section.name}</Name>
            <AddButton
              onClick={handleOnClickAddButton}
              isClose={formActive}
            ></AddButton>
          </HeadRow>
          <FormArea active={formActive}>
            <CreateCardForm section={props.section} />
          </FormArea>
        </Head>
        <Items>
          {props.section.cards !== null &&
            props.section.cards.map(card => (
              <CardPanel card={card} key={card.id} />
            ))}
        </Items>
      </Inner>
    </Wrapper>
  );
};

const FormArea = styled.div<{ active: boolean }>`
  transition: 0.5s;
  max-height: ${props => (props.active ? 200 : 0)}px;
  overflow: hidden;
`;

const Wrapper = styled.div`
  padding: 0 8px;
  display: flex;
  flex-direction: column;
  min-width: 360px;
  box-sizing: border-box;
`;

const Inner = styled.div`
  background: ${props => props.theme.main};
  padding: 8px;
  box-sizing: border-box;
  border-radius: 4px;
`;

const Head = styled.div`
  display: flex;
  flex-direction: column;
`;

const Items = styled.div`
  display: flex;
  flex-direction: column;
`;

const HeadRow = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-between;
`;

const Name = styled.h3`
  margin: 0;
  font-size: 1.1rem;
`;