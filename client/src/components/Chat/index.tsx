import React from "react";
import ChatInput from "../ChatInput";
import useSWR from "swr";
import { Container, Loader } from "@mantine/core";
import Message from "../Message";
import {BASE_URL, TMessage} from '../../requests'

const fetcher = (url: string) => fetch(url).then((res) => res.json());
const useMessages = () => {
  const { data, error, mutate } = useSWR<[TMessage]>(
    `${BASE_URL}/api/v1/messages/`,
    fetcher
  );
  return { data, error, mutate };
};


export default () => {
  const { data, error, mutate } = useMessages()
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
        {data?.map((item, index) => {
          return <Message key={index} {...item} />;
        })}
      </div>
      <ChatInput onSent={mutate} />
    </div>
  );
};
