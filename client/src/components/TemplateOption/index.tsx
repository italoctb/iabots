import { NativeSelect } from "@mantine/core";
import React from "react";
import useSWR from "swr";
import { TOption, TTemplate, updateOption } from "../../requests";

type Props = {
  index: number;
  item: TOption;
  templates: [TTemplate];
  mutate(): void;
};

export default ({ index, item, templates, mutate }: Props) => {
  const onChange = async (e: any) => {
    await updateOption({ ...item, goto: e.target.value });
    mutate()
  }
  return (
      <NativeSelect
        data={templates.map((i) => `${i.id}`)}
        placeholder="Pick one"
        onChange={onChange}
        value={item.goto}
        label={`${index} - ${item.label}`}
        required
      />
  );
};
