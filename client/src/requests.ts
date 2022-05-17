import axios from "axios";

export type TMessage = {
  message: string;
  created: string;
};

export type TTemplate = {
  id: string
  template_message: string
  options: [TOption]
  created: string;
};

export type TOption = {
  id: string
  label: string
  goto: string
}

export const BASE_URL = "http://127.0.0.1:5000";

export const sendMessage = async (message: string) => {
   await axios.post(`${BASE_URL}/api/v1/messages/`, {
    message,
  });
  await axios.get(`${BASE_URL}/api/v1/messages/process`);
};


export const updateOption = async (option:TOption) => {
  await axios.post(`${BASE_URL}/api/v1/templates/option/${option.id}`, {
    goto: option.goto
  });
};

export const getMessages = async () =>{
  const response = await axios.get(`${BASE_URL}/api/v1/messages/`);
  return response.data
}