import React, { useState, useContext } from 'react';
import styled from 'styled-components';
import { Card } from '../model/model';
import { ReactComponent as DeleteIconImage } from '../assets/trash.svg';
import { ReactComponent as EditIconImage } from '../assets/edit.svg';
import { ReactComponent as TimeIconImage } from '../assets/time.svg';
import { DeleteConfirmation } from './DeleteConfirmation';
import sectionsRepository from '../api/sectionsRepository';
import { ViewContext } from '../context/viewContext';
import boardsRepository from '../api/boardsRepository';
import { EditCardForm } from './EditCardForm';
import { DataContext } from '../context/dataContext';
import dayjs from 'dayjs';

interface Props {
  item: Card;
}

export const CardPanelDetail = (props: Props) => {
  const { item } = props;
  const { currentBoard } = useContext(ViewContext);
  const { setSections } = useContext(DataContext);
  const [deleteConfirmation, setDeleteConfirmation] = useState(false);
  const [isEdit, setIsEdit] = useState(false);
  const handleOnClickDelete = (e: React.MouseEvent<HTMLOrSVGElement>) => {
    e.stopPropagation();
    setIsEdit(false);
    setDeleteConfirmation((prev) => !prev);
  };
  const handleOnClickEdit = (e: React.MouseEvent<HTMLOrSVGElement>) => {
    e.stopPropagation();
    setDeleteConfirmation(false);
    setIsEdit((prev) => !prev);
  };

  const handleOnClickCancel = () => {
    setDeleteConfirmation(false);
    setIsEdit(false);
  };

  const handleOnDeleteSubmit = async () => {
    await sectionsRepository.deleteCard(item.section.id, item.id);
    const current = await boardsRepository.getBoardSections(currentBoard.id);
    setSections(current);
  };

  const afterUpdateSubmit = () => {
    setIsEdit(false);
  };

  const displayTime = (): string => {
    if (!item.limitTime) return '';
    return dayjs.unix(item.limitTime).format('M/DD H:mm');
  };

  return (
    <Wrapper>
      <Top>
        {isEdit ? (
          <EditCardForm
            item={item}
            onClickCancel={handleOnClickCancel}
            afterSubmit={afterUpdateSubmit}
          ></EditCardForm>
        ) : (
          <TopInner>
            <TopInnerTop>
              <Description>
                {item.description !== '' ? item.description : 'no description'}
              </Description>
              <Operator>
                <EditIcon onClick={handleOnClickEdit}></EditIcon>
                <DeleteIcon onClick={handleOnClickDelete}></DeleteIcon>
              </Operator>
            </TopInnerTop>
            {item.labels.length > 0 || item.limitTime ? (
              <TopInnerBottom>
                <LabelList>
                  {item.labels.map((v) => (
                    <LabelItem key={v.id}>{v.name}</LabelItem>
                  ))}
                </LabelList>
                {item.limitTime ? (
                  <LimitTime>
                    <TimeIcon></TimeIcon>
                    {displayTime()}
                  </LimitTime>
                ) : null}
              </TopInnerBottom>
            ) : null}
            <Images>
              {item.images &&
                item.images?.length > 0 &&
                item.images?.map((v, i) => <img src={v} key={i} />)}
            </Images>
          </TopInner>
        )}
      </Top>
      <Expand expand={deleteConfirmation}>
        {deleteConfirmation && (
          <ExpandInner>
            <DeleteConfirmation
              onSubmit={handleOnDeleteSubmit}
              onClickCancel={handleOnClickCancel}
            ></DeleteConfirmation>
          </ExpandInner>
        )}
      </Expand>
    </Wrapper>
  );
};

const Wrapper = styled.div``;

const Top = styled.div`
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
`;

const TopInner = styled.div`
  display: flex;
  flex: 1;
  flex-direction: column;
`;

const TopInnerTop = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
`;

const TopInnerBottom = styled.div`
  display: flex;
  justify-content: space-between;
  margin-top: 8px;
  align-items: center;
`;

const LabelList = styled.div`
  display: flex;
`;

const LimitTime = styled.div`
  font-weight: bold;
  display: flex;
  align-items: center;
`;

const LabelItem = styled.div`
  padding: 4px 8px;
  border-radius: 2px;
  border: 1px solid ${(props) => props.theme.border};
  margin-right: 8px;
  background: ${(props) => props.theme.bg};
`;

const Expand = styled.div<{ expand: boolean }>`
  transition: 0.5s;
  max-height: ${(props) => (props.expand ? 128 : 0)}px;
  overflow: hidden;
  box-sizing: border-box;
`;

const ExpandInner = styled.div`
  padding-top: 8px;
`;

const Description = styled.p`
  padding-top: 8px;
  margin: 0;
  white-space: pre-wrap;
`;

const Operator = styled.div`
  display: flex;
`;

const TimeIcon = styled(TimeIconImage)`
  fill: ${(props) => props.theme.solid};
  height: 16px;
  width: 16px;
  margin-right: 4px;
`;

const DeleteIcon = styled(DeleteIconImage)`
  fill: ${(props) => props.theme.solid};
  height: 24px;
  margin-left: 8px;
`;

const EditIcon = styled(EditIconImage)`
  fill: ${(props) => props.theme.solid};
  height: 24px;
`;

const Images = styled.div`
  height: 100%;
  margin-top: 8px;
`;
