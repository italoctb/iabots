import React, { useState } from "react";
import { Button, TextInput } from "@mantine/core";
import { Send, Asset } from "tabler-icons-react";

import axios from "axios";

type Props = {
  onSent(): void;
};

export default ({ onSent }: Props) => {
  const [text, setText] = useState<string>("");
  const handleKeyDown = (event: { key: string; }) => {
    if (event.key === "Enter") {
      sendMsg()
    }
  };
  const sendMsg = async () => {
    try {
      const response = await axios.post(
        "http://localhost:5000/api/v1/messages/",
        {
          message: text,
        }
      );
      setText("");
      onSent();
    } catch (e) {
      console.error(e);
    }
  };

const process = async () => {
  try {
    const response = await axios.get(
      "http://localhost:5000/api/v1/messages/process",
    );
    onSent();
  } catch (e) {
    console.error(e);
  }
};

  return (
    <div
      style={{
        flexDirection: "row",
        display: "flex",
      }}
    >
      <TextInput
        placeholder="Write a message"
        value={text}
        onChange={(event) => setText(event.target.value)}
        onSubmit={() => sendMsg()}
        onKeyDown={handleKeyDown}
      />
      <Button
        variant="outline"
        disabled={text === ""}
        onClick={() => sendMsg()}
        leftIcon={<Send size={14} />}
      >
        Send
      </Button>
      <Button
        variant="outline"
        onClick={() => process()}
        leftIcon={<Asset size={14} />}
      >
        Process
      </Button>
    </div>
  );
};
