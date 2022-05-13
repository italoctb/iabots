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

export const sendMessage = async (message: string) => {
   await axios.post("http://localhost:5000/api/v1/messages/", {
    message,
  });
  await axios.get("http://localhost:5000/api/v1/messages/process");
};


export const updateOption = async (option:TOption) => {
  await axios.post("http://localhost:5000/api/v1/templates/option/"+option.id, {
    goto: option.goto
  });
};