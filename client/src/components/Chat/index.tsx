import React from "react";
import ChatInput from "../ChatInput";
import useSWR from "swr";
import { Container, Loader } from "@mantine/core";
import Message from "../Message";

const fetcher = (url: string) => fetch(url).then((res) => res.json());

type Message = {
  message: string;
  created: string;
};
export default () => {
  const { data, error, mutate } = useSWR<[Message]>(
    "http://127.0.0.1:5000/api/v1/messages/",
    fetcher
  );
  console.log(data);
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        paddingLeft: "16px",
        borderRadius: 16,
        marginTop: "16px",
        paddingBottom: "16px",
        backgroundColor: "white"
      }}
    >
      <div
        style={{
          display: "flex",
          flexDirection: "column-reverse",
          maxHeight: "500px",
          overflowY: "scroll",
        }}
      >
        {!data && !error && <Loader />}
        {data?.map((item) => {
          return <Message {...item} />;
        })}
      </div>
      <ChatInput onSent={mutate} />
    </div>
  );
};
