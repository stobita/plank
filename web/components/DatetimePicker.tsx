import React, { useState } from 'react';
import DatePicker from 'react-datepicker';
import CloseIconImage from '../assets/close.svg';
import styled from 'styled-components';
import { Input } from './Input';
import { Button } from './Button';
import dayjs from 'dayjs';

type Props = {
  value: number | undefined;
  onChange: (unixtime: number) => void;
  placeholder?: string;
};

export const DatetimePicker = (props: Props) => {
  const [limitTimeActive, setLimitTimeActive] = useState(!!props.value);
  const onChangeDate = (date: Date) => {
    const got = dayjs(date);
    const current = dayjs.unix(props.value || dayjs().unix());
    const result = current
      .year(got.year())
      .month(got.month())
      .date(got.date())
      .hour(0)
      .minute(0);

    props.onChange(result.unix());
  };
  const handleOnClickAddLimit = () => {
    setLimitTimeActive(true);
  };
  const handleOnClickClose = () => {
    setLimitTimeActive(false);
  };
  const onChangeTime = (date: Date) => {
    const got = dayjs(date);
    const current = dayjs.unix(props.value || dayjs().unix());
    const result = current.hour(got.hour()).minute(got.minute());

    props.onChange(result.unix());
  };
  const DateInput = (props: { value?: string; onClick?: () => void }) => {
    return <Input onClick={props.onClick} value={props.value}></Input>;
  };

  const TimeButton = (props: {
    value?: string;
    onClick?: () => void;
    onChange?: () => void;
  }) => {
    return (
      <Input
        value={props.value}
        onFocus={props.onClick}
        onChange={props.onClick}
      ></Input>
    );
  };

  return (
    <>
      {limitTimeActive ? (
        <Wrapper>
          <DateWrapper>
            <DatePicker
              onChange={onChangeDate}
              selected={props.value ? dayjs.unix(props.value).toDate() : null}
              placeholderText={props.placeholder}
              dateFormat="yyyy/MM/dd"
              customInput={<DateInput />}
            ></DatePicker>
          </DateWrapper>
          <TimeWrapper>
            <DatePicker
              onChange={onChangeTime}
              selected={props.value ? dayjs.unix(props.value).toDate() : null}
              showTimeSelect
              showTimeSelectOnly
              timeIntervals={15}
              dateFormat="H:mm"
              timeFormat="H:mm"
              customInput={<TimeButton />}
            ></DatePicker>
          </TimeWrapper>
          <CloseIcon onClick={handleOnClickClose}></CloseIcon>
        </Wrapper>
      ) : (
        <Button type="button" onClick={handleOnClickAddLimit}>
          add limit
        </Button>
      )}
    </>
  );
};

const DateWrapper = styled.div`
  flex: 2;
`;
const TimeWrapper = styled.div`
  flex: 1;
  margin-left: 8px;
`;

const Wrapper = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-between;
  .react-datepicker {
    background: ${(props) => props.theme.main};
    color: ${(props) => props.theme.solid};
    border: 1px solid ${(props) => props.theme.border};
  }
  .react-datepicker-popper[data-placement^='bottom'] {
    margin-top: 8px;
  }
  .react-datepicker__header {
    border: none;
    border-bottom: 1px solid ${(props) => props.theme.border};
    background: transparent;
  }
  .react-datepicker-wrapper {
    width: 100%;
  }
  .react-datepicker__current-month,
  .react-datepicker-time__header,
  .react-datepicker__day,
  .react-datepicker__day-name {
    color: ${(props) => props.theme.solid};
  }
  .react-datepicker__time-container {
    border: 1px solid ${(props) => props.theme.border};
  }

  .react-datepicker__day--keyboard-selected {
    background: ${(props) => props.theme.primary};
  }

  .react-datepicker__time-container .react-datepicker__time {
    background: transparent;
  }

  .react-datepicker__triangle {
    border-bottom-color: ${(props) => props.theme.main} !important;
  }
  .react-datepicker__day:hover {
    background: ${(props) => props.theme.bg};
  }

  .react-datepicker__time-container
    .react-datepicker__time
    .react-datepicker__time-box
    ul.react-datepicker__time-list
    li.react-datepicker__time-list-item:hover {
    background: ${(props) => props.theme.bg};
  }
`;

const CloseIcon = styled(CloseIconImage)`
  fill: ${(props) => props.theme.solid};
  height: 24px;
  width: 24px;
  margin: 0 8px;
`;
