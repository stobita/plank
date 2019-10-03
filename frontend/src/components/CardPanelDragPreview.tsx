import React from "react";
import { CSSProperties } from "styled-components";
import { useDragLayer } from "react-dnd";
import { CardPanel } from "./CardPanel";

export const CardPanelDragPreview = () => {
  const { isDragging, currentOffset, item } = useDragLayer(monitor => ({
    item: monitor.getItem(),
    isDragging: !!monitor.isDragging(),
    currentOffset: monitor.getSourceClientOffset()
  }));

  if (!currentOffset) {
    return null;
  }

  // for performance
  const style: CSSProperties = {
    display: "flex",
    height: "100%",
    width: "320px",
    position: "fixed",
    pointerEvents: "none",
    left: "0",
    top: "0",
    zIndex: 1,
    transform: `translate(${currentOffset.x}px, ${currentOffset.y}px)`
  };

  return (
    <>
      {isDragging && (
        <div style={style}>
          <CardPanel card={item}></CardPanel>
        </div>
      )}
    </>
  );
};
