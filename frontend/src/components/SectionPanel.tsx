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
            <AddButton onClick={handleOnClickAddButton}></AddButton>
          </HeadRow>
          {formActive && <CreateCardForm section={props.section} />}
        </Head>
        <Items>
          {props.section.cards !== null &&
            props.section.cards.map(card => <CardPanel card={card} />)}
        </Items>
      </Inner>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  padding: 0 8px;
  display: flex;
  flex-direction: column;
  min-width: 280px;
`;

const Inner = styled.div`
  background: ${props => props.theme.main};
  padding: 8px;
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
