import React, { useState } from "react";
import { Button, TextInput } from "@mantine/core";
import { Send } from "tabler-icons-react";

import { sendMessage } from "../../requests";

type Props = {
  onSent(): void;
};

export default ({ onSent }: Props) => {
  const [text, setText] = useState<string>("");
  const handleKeyDown = (event: { key: string }) => {
    if (event.key === "Enter") {
      sendMsg();
    }
  };
  const sendMsg = async () => {
    try {
      await sendMessage(text)
      setText("");
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
        style={{ flex: 1 }}
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
    </div>
  );
};
