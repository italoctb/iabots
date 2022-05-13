import React from "react";
import moment from "moment";

type Props = {
  message: string;
  created: string;
};

export default ({ message, created }: Props) => {
  return (
    <div style={{ display: "flex", maxWidth: '600px' }}>
      <p style={{whiteSpace: 'pre-line'}}>{message}</p>
      <small
        style={{ alignSelf: "flex-end", margin: "10px", fontSize: "10px" }}
      >
        {moment(created).format("HH:mm")}
      </small>
    </div>
  );
};
