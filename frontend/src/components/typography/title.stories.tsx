import type { Meta, StoryObj } from "@storybook/react";

import { Title } from ".";

const meta: Meta<typeof Title> = {
  component: Title,
};

export default meta;
type Story = StoryObj<typeof Title>;

export const TitleComponent: Story = {
  render: () => <Title>Text</Title>,
};
