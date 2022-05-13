import React from "react";
import useSWR from "swr";
import { TTemplate } from "../../requests";
import TemplateOption from "../TemplateOption";

const fetcher = (url: string) => fetch(url).then((res) => res.json());
const useTemplates = () => {
  const { data, error, mutate } = useSWR<[TTemplate]>(
    "http://127.0.0.1:5000/api/v1/templates/",
    fetcher
  );
  return { data, error, mutate };
};

export default () => {
  const { data, error, mutate } = useTemplates();
  return (
    <div
      style={{
        backgroundColor: "white",
      }}
    >
      {data?.map((item, index) => (
        <>
          <p>
            <small>{item.id} - </small>
            {item.template_message}{" "}
          </p>
            {item.options.map((item, index) => (
              <TemplateOption index={index} item={item} templates={data} mutate={mutate} />
            ))}
        </>
      ))}
    </div>
  );
};
