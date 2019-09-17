import React from "react";
import styled from "styled-components";
import { ButtonPair } from "./ButtonPair";
import { Button } from "./Button";

interface Props {
  onSubmit: () => void;
  onClickCancel: () => void;
}

export const DeleteConfirmation = (props: Props) => {
  const handleOnSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    props.onSubmit();
  };
  const onClickWrapper = (e: React.MouseEvent<HTMLElement>) => {
    e.stopPropagation();
  };
  return (
    <Wrapper onClick={onClickWrapper}>
      <Message>Are you sure you want to delete it?</Message>
      <form onSubmit={handleOnSubmit}>
        <ButtonPair
          left={<Button danger>Delete</Button>}
          right={
            <Button type="button" onClick={props.onClickCancel}>
              Cancel
            </Button>
          }
        />
      </form>
    </Wrapper>
  );
};

const Wrapper = styled.div``;

const Message = styled.p`
  color: ${props => props.theme.danger};
  margin-bottom: 8px;
  font-weight: bold;
`;
