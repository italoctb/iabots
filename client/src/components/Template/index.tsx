import React from "react";
import useSWR from "swr";
import { BASE_URL, TTemplate } from "../../requests";
import TemplateOption from "../TemplateOption";

const fetcher = (url: string) => fetch(url).then((res) => res.json());
const useTemplates = () => {
  const { data, error, mutate } = useSWR<[TTemplate]>(
    `${BASE_URL}/api/v1/templates/`,
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
