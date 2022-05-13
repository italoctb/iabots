import axios from "axios";

export type TMessage = {
  message: string;
  created: string;
};

export const sendMessage = async (message: string) => {
   await axios.post("http://localhost:5000/api/v1/messages/", {
    message,
  });
  await axios.get("http://localhost:5000/api/v1/messages/process");
};